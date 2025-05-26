import { Module } from '@nestjs/common';
import { RabbitmqInfrastructure } from './rabbitmq/rabbitmq';
import { ConfigModule } from '@nestjs/config';
import { Exchange } from '../enum/exchange';
import { RabbitMQModule } from '@golevelup/nestjs-rabbitmq';
import { ClientsModule, Transport } from '@nestjs/microservices';
import { glob } from 'fast-glob';
import { join } from 'path';
import { Service } from '../enum/service';
import { ProductRpcService } from './rpc/product.rpc.service';
import { UserRpcService } from './rpc/user.rpc.service';
import { JaegerTelemetryService } from './telemetry/jaeger.telemetry.service';
import { ConsulModule } from '../config/consul.module';
import { ConsulService } from '../config/consul.service';
import { RabbitMQRootAsync } from '../config/configRoot';


@Module({
  imports: [
    ConfigModule.forRoot(),
    ClientsModule.registerAsync([
      {
        name: Service.ProductService.toString(),
        imports: [ConsulModule],
        inject: [ConsulService],
        useFactory: async (config: ConsulService) => ({
          transport: Transport.GRPC,
          options: {
            url: `${await config.get('/services/product/RPC_HOST')}:${await config.get('/services/product/RPC_PORT')}`,
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
        imports: [ConsulModule],
        inject: [ConsulService],
        useFactory: async (configService: ConsulService) => ({
          transport: Transport.GRPC,
          options: {
            url: `${ await configService.get('/services/user/RPC_HOST')}:${await configService.get('/services/user/RPC_PORT')}`,
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
    RabbitMQRootAsync,
    ConsulModule,
  ],
  providers: [RabbitmqInfrastructure, ProductRpcService, UserRpcService, JaegerTelemetryService],
  exports: [RabbitmqInfrastructure, ProductRpcService, UserRpcService, JaegerTelemetryService],
})
export class InfrastructureModule {
}
