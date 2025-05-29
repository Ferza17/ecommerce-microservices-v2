import { Inject, Injectable, Logger, OnModuleInit } from '@nestjs/common';
import { Service } from '../../enum/service';
import { ClientGrpc } from '@nestjs/microservices';
import { FindUserByIdRequest, User } from '../../model/rpc/userMessage';
import { Metadata } from '@grpc/grpc-js';
import { Header } from '../../enum/header';
import { JaegerTelemetryService } from '../telemetry/jaeger.telemetry.service';
import { Context, propagation } from '@opentelemetry/api';
import { UserServiceService } from '../../model/rpc/userService';
import CircuitBreaker from 'opossum';
import { lastValueFrom } from 'rxjs';
import { UserServiceCircuitOptions } from '../../config/circuitOptions';

@Injectable()
export class UserRpcService implements OnModuleInit {
  private readonly logger = new Logger(UserRpcService.name);
  private userService: any;


  constructor(
    @Inject(Service.UserService.toString()) private client: ClientGrpc,
    private readonly otel: JaegerTelemetryService,
  ) {
  }

  onModuleInit() {
    this.userService = this.client.getService<typeof UserServiceService>('UserService');
  }

  async findUserById(requestId: string, req: FindUserByIdRequest, context?: Context): Promise<User> {
    const span = this.otel.tracer('RpcService.findUserById', context);
    try {
      const breakerFn: (requestId: string, req: FindUserByIdRequest, context?: Context) => Promise<User> =
        async (requestId, req, context) => {
          const metadata = new Metadata();
          metadata.set(Header.X_REQUEST_ID, requestId);
          if (context) {
            propagation.inject(context, metadata, {
              set: (metadata, key, value) => metadata.set(key, value as string),
            });
          }
          const observableResult = this.userService.findUserById(req, metadata);
          return lastValueFrom(observableResult);
        };

      const cb = new CircuitBreaker<[string, FindUserByIdRequest, Context], User>(breakerFn, UserServiceCircuitOptions);
      return await cb.fire(requestId, req, <Context>context);
    } catch (e) {
      this.logger.error(`requestId: ${requestId} , error: ${e.message}`);
      span.recordException(e);
      throw e;
    } finally {
      span.end();
    }
  }
}

