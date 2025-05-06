import { NestFactory } from '@nestjs/core';
import { RpcServerModule } from './rpc.module';
import { MicroserviceOptions } from '@nestjs/microservices';
import { GrpcClientOptions } from './options';
import { RequestIdInterceptor } from './interceptor/requestIdInterceptor.service';


export class GrpcServer {
  constructor(
    private readonly grpcClientOptions: GrpcClientOptions,
  ) {
  }

  async Serve() {
    const app = await NestFactory.createMicroservice<MicroserviceOptions>(RpcServerModule, this.grpcClientOptions.getGRPCConfig);
    app.useGlobalInterceptors(new RequestIdInterceptor());
    await app.listen();
  }
}
