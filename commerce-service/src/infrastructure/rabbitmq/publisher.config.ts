import { Transport } from '@nestjs/microservices';
import { ClientsProviderAsyncOptions } from '@nestjs/microservices/module/interfaces/clients-module.interface';
import { Queue } from '../../enum/queue';
import { ConsulService } from '../../config/consul.service';
import { ConfigService } from '@nestjs/config';
import { ConsulModule } from '../../config/consul.module';

export const ClientModuleAsyncConfig = (exchange: string, queue: string): ClientsProviderAsyncOptions => ({
  name: queue,
  imports: [ConsulModule],
  inject: [ConsulService],
  useFactory: async (configService: ConfigService, queue: Queue) => ({
    transport: Transport.RMQ,
    options: {
      urls: [`amqp://${await configService.get("/broker/rabbitmq/RABBITMQ_USERNAME")}:${await configService.get("/broker/rabbitmq/RABBITMQ_PASSWORD")}@${await configService.get("/broker/rabbitmq/RABBITMQ_HOST")}:${await configService.get("/broker/rabbitmq/RABBITMQ_PORT")}`],
      queue: queue,
      queueOptions: {
        durable: true,
      },
      routingKey: queue,
      exchange: exchange,
      exchangeType: 'direct',
      noDelay: true,
      gracefulShutdown: true,
    },
  }),
});