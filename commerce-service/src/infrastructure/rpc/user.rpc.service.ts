import { Inject, Injectable, Logger, OnModuleInit } from '@nestjs/common';
import { Service } from '../../enum/service';
import { ClientGrpc } from '@nestjs/microservices';
import { UserServiceService } from '../../model/rpc/userService';
import { FindUserByIdRequest, User } from '../../model/rpc/userMessage';
import { lastValueFrom } from 'rxjs';
import { Metadata } from '@grpc/grpc-js';
import { Header } from '../../enum/header';
import { JaegerTelemetryService } from '../telemetry/jaeger.telemetry.service';
import { Context } from '@opentelemetry/api';


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
      const metadata = new Metadata();
      metadata.set(Header.X_REQUEST_ID, requestId);
      return lastValueFrom(this.userService.findUserById(req, metadata));
    } catch (e) {
      this.logger.error(`requestId: ${requestId} , error: ${e.message}`);
      span.recordException(e);
      throw e;
    } finally {
      span.end();
    }
  }
}