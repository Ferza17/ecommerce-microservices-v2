import { Module } from '@nestjs/common';
import { RabbitmqInfrastructure } from './rabbitmq/rabbitmq';
import { ConfigModule, ConfigService } from '@nestjs/config';
import { Exchange } from '../enum/exchange';
import { RabbitMQModule } from '@golevelup/nestjs-rabbitmq';
import { ClientsModule, Transport } from '@nestjs/microservices';
import { glob } from 'fast-glob';
import { join } from 'path';
import { Service } from '../enum/service';
import { ProductRpcService } from './rpc/product.rpc.service';
import { UserRpcService } from './rpc/user.rpc.service';


@Module({
  imports: [
    ConfigModule.forRoot(),
    ClientsModule.registerAsync([
      {
        name: Service.ProductService.toString(),
        imports: [ConfigModule],
        inject: [ConfigService],
        useFactory: (configService: ConfigService) => ({
          transport: Transport.GRPC,
          options: {
            url: `${configService.get<number>('PRODUCT_SERVICE_RPC_HOST')}:${configService.get<number>('PRODUCT_SERVICE_RPC_PORT')}`,
            package: 'proto',
            protoPath: glob.sync(['proto/*.proto'], {
              cwd: join(__dirname, '../../'),
              absolute: true,
            }),
            loader: {
              oneofs: true,
            },
          },
        }),
      },
    ]),
    ClientsModule.registerAsync([
      {
        name: Service.UserService.toString(),
        imports: [ConfigModule],
        inject: [ConfigService],
        useFactory: (configService: ConfigService) => ({
          transport: Transport.GRPC,
          options: {
            url: `${configService.get<number>('USER_SERVICE_RPC_HOST')}:${configService.get<number>('USER_SERVICE_RPC_PORT')}`,
            package: 'proto',
            protoPath: glob.sync(['proto/*.proto'], {
              cwd: join(__dirname, '../../'),
              absolute: true,
            }),
            loader: {
              oneofs: true,
            },
          },
        }),
      },
    ]),
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
        ],
      }),
    }),
  ],
  providers: [RabbitmqInfrastructure, ProductRpcService, UserRpcService],
  exports: [RabbitmqInfrastructure, ProductRpcService, UserRpcService],
})
export class InfrastructureModule {
}
