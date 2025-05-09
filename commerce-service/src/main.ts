import { GrpcServer } from './server/rpc/rpc';
import { ConfigService } from '@nestjs/config';
import { GrpcClientOptions } from './server/rpc/options';
import { RabbitmqConsumer } from './server/rabbitmq/rabbitmq';

function bootstrap() {
  const configService: ConfigService = new ConfigService();

  const rmqConsumer = new RabbitmqConsumer(configService);
  rmqConsumer.Serve();

  const grpcClientOptions: GrpcClientOptions = new GrpcClientOptions(configService);
  const grpcServer = new GrpcServer(grpcClientOptions);
  grpcServer.Serve();

}

bootstrap();
