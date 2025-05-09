import { Module } from '@nestjs/common';
import { RabbitmqInfrastructure } from './rabbitmq/rabbitmq';
import { RpcInfrastructure } from './rpc/rpc.infrastructure';
import { ConfigModule, ConfigService } from '@nestjs/config';
import { ClientsModule, Transport } from '@nestjs/microservices';
import { Exchange } from '../enum/exchange';
import { Queue } from '../enum/queue';
import { ClientModuleAsyncConfig } from './rabbitmq/publisher.config';
import { RabbitMQModule } from '@golevelup/nestjs-rabbitmq';

@Module({
  imports: [
    ConfigModule.forRoot(),
    // ClientsModule.registerAsync([
    //   {
    //     name: Queue.EVENT_CREATED.toString(),
    //     inject: [ConfigService],
    //     useFactory: (configService: ConfigService) => ({
    //       transport: Transport.RMQ,
    //       options: {
    //         urls: [`amqp://${configService.get<string>('RABBITMQ_USERNAME')}:${configService.get<string>('RABBITMQ_PASSWORD')}@${configService.get<string>('RABBITMQ_HOST')}:${configService.get<number>('RABBITMQ_PORT')}`],
    //         queue: Queue.EVENT_CREATED.toString(),
    //         queueOptions: {
    //           durable: true,
    //         },
    //         routingKey: Queue.EVENT_CREATED.toString(),
    //         exchange: Exchange.EventExchange.toString(),
    //         exchangeType: 'direct',
    //         noDelay: true,
    //         gracefulShutdown: true,
    //         noAssert: true,
    //       },
    //     }),
    //   }
    // ]),
    RabbitMQModule.forRootAsync({
      imports: [ConfigModule],
      inject: [ConfigService],
      useFactory: (configService: ConfigService) => ({
        uri: `amqp://${configService.get<string>('RABBITMQ_USERNAME')}:${configService.get<string>('RABBITMQ_PASSWORD')}@${configService.get<string>('RABBITMQ_HOST')}:${configService.get<number>('RABBITMQ_PORT')}`,
        exchanges: [
          {
            name: Exchange.EventExchange.toString(),
            type: 'direct',
          },
        ]
      })
    })
  ],
  providers: [RabbitmqInfrastructure, RpcInfrastructure],
  exports: [RabbitmqInfrastructure, RpcInfrastructure],
})
export class InfrastructureModule {
}
