import { Inject, Injectable, Logger, OnModuleInit } from '@nestjs/common';
import { Service } from '../../enum/service';
import { ClientGrpc } from '@nestjs/microservices';
import { FindUserByIdRequest, User } from '../../model/rpc/userMessage';
import { Metadata } from '@grpc/grpc-js';
import { Header } from '../../enum/header';
import { JaegerTelemetryService } from '../telemetry/jaeger.telemetry.service';
import { Context } from '@opentelemetry/api';
import { UserServiceService } from '../../model/rpc/userService';
import CircuitBreaker from 'opossum';

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
    this.userService = this.client.getService<typeof UserServiceService>("UserService");
  }

  async findUserById(requestId: string, req: FindUserByIdRequest, context?: Context): Promise<User> {
    const span = this.otel.tracer('RpcService.findUserById', context);
    try {
      const metadata = new Metadata();
      metadata.set(Header.X_REQUEST_ID, requestId);

      const breaker = new CircuitBreaker(this.userService.findUserById(req, metadata), {
        timeout: 1000,
        errorThresholdPercentage: 50,
        resetTimeout: 10000,
      });

      breaker.on('success', (result, latencyMs) => {
        this.logger.log(`findUserById requestId: ${requestId} , latencyMs: ${latencyMs}`);
      })
      breaker.on('failure', (result, latencyMs) => {
        this.logger.log(`findUserById requestId: ${requestId} , latencyMs: ${latencyMs}`);
      })
      breaker.on('open', () => {
        this.logger.log(`findUserById requestId: ${requestId} , circuit breaker is open`);
      })
      breaker.on('close', () => {
        this.logger.log(`findUserById requestId: ${requestId} , circuit breaker is closed`);
      })
      breaker.on('halfOpen', () => {
        this.logger.log(`findUserById requestId: ${requestId} , circuit breaker is half open`);
      })
      return await breaker.fire() as User;
    } catch (e) {
      this.logger.error(`requestId: ${requestId} , error: ${e.message}`);
      span.recordException(e);
      throw e;
    } finally {
      span.end();
    }
  }
}