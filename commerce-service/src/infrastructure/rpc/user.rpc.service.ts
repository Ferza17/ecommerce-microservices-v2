import { Inject, Injectable, Logger, OnModuleInit } from '@nestjs/common';
import { Service } from '../../enum/service';
import { ClientGrpc } from '@nestjs/microservices';
import { FindUserByIdRequest, User } from '../../model/rpc/userMessage';
import { Metadata } from '@grpc/grpc-js';
import { Header } from '../../enum/header';
import { JaegerTelemetryService } from '../telemetry/jaeger.telemetry.service';
import { Context } from '@opentelemetry/api';
import { UserServiceService } from '../../model/rpc/userService';
import CircuitBreaker, { Options as CircuitBreakerOptions } from 'opossum';
import { lastValueFrom } from 'rxjs';
import {UserServiceCircuitOptions} from '../../config/circuitOptions'

@Injectable()
export class UserRpcService implements OnModuleInit {
  private readonly logger = new Logger(UserRpcService.name);
  private userService: any;
  private findUserByIdBreaker: CircuitBreaker<[FindUserByIdRequest, Metadata], User>;


  constructor(
    @Inject(Service.UserService.toString()) private client: ClientGrpc,
    private readonly otel: JaegerTelemetryService,
  ) {
  }

  onModuleInit() {
    this.userService = this.client.getService<typeof UserServiceService>('UserService');

    const breakerFn: (req: FindUserByIdRequest, metadata: Metadata) => Promise<User> =
      async (req, metadata) => {
        const observableResult = this.userService.findUserById(req, metadata);
        return lastValueFrom(observableResult);
      };

    this.findUserByIdBreaker = new CircuitBreaker<[FindUserByIdRequest, Metadata], User>(breakerFn, UserServiceCircuitOptions);
  }

  async findUserById(requestId: string, req: FindUserByIdRequest, context?: Context): Promise<User> {
    const span = this.otel.tracer('RpcService.findUserById', context);
    try {
      const metadata = new Metadata();
      metadata.set(Header.X_REQUEST_ID, requestId);
      return await this.findUserByIdBreaker.fire(req, metadata)
    } catch (e) {
      this.logger.error(`requestId: ${requestId} , error: ${e.message}`);
      span.recordException(e);
      throw e;
    } finally {
      span.end();
    }
  }
}