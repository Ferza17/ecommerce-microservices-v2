import * as joi from 'joi';
import { ConfigModule, ConfigService } from '@nestjs/config';
import Environment from '../enum/environment';
import { MongooseModule } from '@nestjs/mongoose';

export const configRoot = ConfigModule.forRoot({
  envFilePath: '.env',
  isGlobal: true,
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
    PRODUCT_SERVICE_RPC_HOST: joi.string().required(),
    PRODUCT_SERVICE_RPC_PORT: joi.number().required(),
    USER_SERVICE_RPC_HOST: joi.string().required(),
    USER_SERVICE_RPC_PORT: joi.number().required(),
  }),
});

export const MongooseRootAsync = MongooseModule.forRootAsync({
  inject: [ConfigService],
  useFactory: (configService: ConfigService) => ({
    uri: `mongodb://${configService.get('MONGO_USERNAME')}:${configService.get('MONGO_PASSWORD')}@${configService.get('MONGO_HOST')}:${configService.get('MONGO_PORT')}/${configService.get('MONGO_DATABASE_NAME')}?authSource=admin`,
  }),
});

