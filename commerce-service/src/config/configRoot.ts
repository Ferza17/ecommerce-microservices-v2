import * as joi from 'joi';
import { ConfigModule } from '@nestjs/config';
import Environment from '../enum/environment';

export const configRoot = ConfigModule.forRoot({
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
});
