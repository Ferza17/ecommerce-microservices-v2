import { GrpcServer } from './server/rpc/rpc';
import { ConfigService } from '@nestjs/config';
import { GrpcClientOptions } from './server/rpc/options';

async function bootstrap() {

  const configService: ConfigService = new ConfigService();
  const grpcClientOptions: GrpcClientOptions = new GrpcClientOptions(configService);
  const grpcServer = new GrpcServer( grpcClientOptions);
  await grpcServer.Serve();

  // TODO: RabbitMQ Server

}

bootstrap();
