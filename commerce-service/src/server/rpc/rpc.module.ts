import { Module } from '@nestjs/common';
import { ModuleModule } from '../../module/module.module';
import { ConfigModule, ConfigService } from '@nestjs/config';
import { GrpcReflectionModule } from 'nestjs-grpc-reflection';
import { GrpcClientOptions } from './options';
import * as joi from 'joi';
import Environment from '../../enum/environment';


@Module({
  imports: [
    ConfigModule.forRoot({
      envFilePath: '.env',
      validationSchema: joi.object({
        ENV: joi.string()
          .valid(
            Environment.DEVELOPMENT,
            Environment.LOCAL,
            Environment.PRODUCTION,
          )
          .default(Environment.LOCAL),
        SERVICE_NAME: joi.string().required(),
        MONGO_USERNAME: joi.string().required(),
        MONGO_PASSWORD: joi.string().required(),
        MONGO_HOST: joi.string().required(),
        MONGO_PORT: joi.number().required(),
        MONGO_DATABASE_NAME: joi.string().required(),
        RABBITMQ_USERNAME: joi.string().required(),
        RABBITMQ_PASSWORD: joi.string().required(),
        RABBITMQ_HOST: joi.string().required(),
        RABBITMQ_PORT: joi.number().required(),
        RPC_HOST: joi.string().required(),
        RPC_PORT: joi.number().required(),
      }),
      isGlobal: true,
    }),
    ModuleModule,
    GrpcReflectionModule.registerAsync({
      useFactory: async (configService: ConfigService) => {
        const grpcClientOptions: GrpcClientOptions = new GrpcClientOptions(configService);
        return grpcClientOptions.getGRPCConfig;
      },
      inject: [ConfigService],
    }),


  ],
  providers: [],
})
export class RpcServerModule {
}
