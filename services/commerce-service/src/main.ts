import { GrpcServer } from './server/rpc/rpc';
import { ConfigService } from '@nestjs/config';
import { GrpcClientOptions } from './server/rpc/options';
import { RabbitmqConsumer } from './server/rabbitmq/rabbitmq';
import { ConsulService } from './config/consul.service';


function bootstrap() {
  const configService: ConfigService = new ConfigService();
  const consulConfig: ConsulService = new ConsulService(configService);

  const rmqConsumer = new RabbitmqConsumer(consulConfig);
  rmqConsumer.Serve();

  const grpcClientOptions: GrpcClientOptions = new GrpcClientOptions(consulConfig);
  const grpcServer = new GrpcServer(grpcClientOptions);
  grpcServer.Serve();

}

bootstrap();
