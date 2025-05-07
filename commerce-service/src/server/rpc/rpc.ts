import { NestFactory } from '@nestjs/core';
import { RpcServerModule } from './rpc.module';
import { MicroserviceOptions } from '@nestjs/microservices';
import { GrpcClientOptions } from './options';
import { RequestIdInterceptor } from './interceptor/requestIdInterceptor.service';
import { Logger } from '@nestjs/common';


export class GrpcServer {
  private readonly logger = new Logger(GrpcServer.name);

  constructor(
    private readonly grpcClientOptions: GrpcClientOptions,
  ) {
  }

  async Serve() {
    const app = await NestFactory.createMicroservice<MicroserviceOptions>(RpcServerModule, this.grpcClientOptions.getGRPCConfig);
    app.useGlobalInterceptors(new RequestIdInterceptor());
    await app.listen();
    this.logger.log('GRPC Server is running...');
  }
}
