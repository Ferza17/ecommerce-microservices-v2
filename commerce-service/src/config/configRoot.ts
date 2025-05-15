import * as joi from 'joi';
import { ConfigModule } from '@nestjs/config';
import Environment from '../enum/environment';
import { MongooseModule } from '@nestjs/mongoose';
import { ConsulModule } from './consul.module';
import { ConsulService } from './consul.service';
import { RabbitMQModule } from '@golevelup/nestjs-rabbitmq';
import { Exchange } from '../enum/exchange';

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
    CONSUL_HOST: joi.string(),
    CONSUL_PORT: joi.number(),
  }),
});

export const MongooseRootAsync = MongooseModule.forRootAsync({
  imports: [ConsulModule],
  inject: [ConsulService],
  useFactory: async (configService: ConsulService) => ({
    uri: `mongodb://${await configService.get('/database/mongodb/MONGO_USERNAME')}:${await configService.get('/database/mongodb/MONGO_PASSWORD')}@${await configService.get('/database/mongodb/MONGO_HOST')}:${await configService.get('/database/mongodb/MONGO_PORT')}/${await configService.get('/database/mongodb/MONGO_DATABASE_NAME/COMMERCE')}?authSource=admin`,
  }),
});

export const RabbitMQRootAsync = RabbitMQModule.forRootAsync({
  imports: [ConsulModule],
  inject: [ConsulService],
  useFactory: async (configService: ConsulService) => ({
    uri: `amqp://${await configService.get("/broker/rabbitmq/RABBITMQ_USERNAME")}:${await configService.get("/broker/rabbitmq/RABBITMQ_PASSWORD")}@${await configService.get("/broker/rabbitmq/RABBITMQ_HOST")}:${await configService.get("/broker/rabbitmq/RABBITMQ_PORT")}`,
    exchanges: [
      {
        name: Exchange.EventExchange.toString(),
        type: 'direct',
      },
    ],
  }),
})