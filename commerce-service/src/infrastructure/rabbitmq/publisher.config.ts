import { Transport } from '@nestjs/microservices';
import { ConfigService } from '@nestjs/config';
import { ClientsProviderAsyncOptions } from '@nestjs/microservices/module/interfaces/clients-module.interface';
import { Queue } from '../../enum/queue';

export const ClientModuleAsyncConfig = (exchange: string, queue: string): ClientsProviderAsyncOptions => ({
  name: queue,
  inject: [ConfigService],
  useFactory: (configService: ConfigService, queue: Queue) => ({
    transport: Transport.RMQ,
    options: {
      urls: [`amqp://${configService.get<string>('RABBITMQ_USERNAME')}:${configService.get<string>('RABBITMQ_PASSWORD')}@${configService.get<string>('RABBITMQ_HOST')}:${configService.get<number>('RABBITMQ_PORT')}`],
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