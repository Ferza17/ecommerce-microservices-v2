import { GrpcServer } from './server/rpc/rpc';
import { ConfigService } from '@nestjs/config';
import { GrpcClientOptions } from './server/rpc/options';
import { RabbitmqConsumer } from './server/rabbitmq/rabbitmq';
import { RabbitmqOptions } from './server/rabbitmq/options';
import { Queue } from './enum/queue';
import { Exchange } from './enum/exchange';

function bootstrap() {
  const configService: ConfigService = new ConfigService();

  const rmqCartEventConfig = new RabbitmqOptions(configService, Queue.CartQueue, Exchange.CartExchange);
  const rmqConsumer = new RabbitmqConsumer(rmqCartEventConfig);
  rmqConsumer.Serve();

  const grpcClientOptions: GrpcClientOptions = new GrpcClientOptions(configService);
  const grpcServer = new GrpcServer(grpcClientOptions);
  grpcServer.Serve();

}

bootstrap();
