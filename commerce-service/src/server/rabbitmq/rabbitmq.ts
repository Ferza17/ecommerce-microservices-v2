import { Controller, Logger, OnModuleInit } from '@nestjs/common';
import { RabbitmqOptions } from './consumer.config';
import { NestFactory } from '@nestjs/core';
import { RabbitmqModule } from './rabbitmq.module';
import { MicroserviceOptions, Transport } from '@nestjs/microservices';
import { ConfigService } from '@nestjs/config';
import { Queue } from '../../enum/queue';

export class RabbitmqConsumer {
  private readonly logger = new Logger(RabbitmqConsumer.name);

  constructor(
    private readonly configService: ConfigService,
  ) {
  }

  async Serve() {
    const options = [
      {
        queue: Queue.CART_CREATED,
        option: new RabbitmqOptions(this.configService, Queue.CART_CREATED).getRabbitmqOptions(),
      },
      {
        queue: Queue.CART_UPDATED,
        option: new RabbitmqOptions(this.configService, Queue.CART_UPDATED).getRabbitmqOptions(),
      },
    ];

    for (const option of options) {
      this.logger.log(`Rabbitmq Consumer ${option.queue} is running...`);
      const app = await NestFactory.createMicroservice<MicroserviceOptions>(RabbitmqModule, option.option);
      app.listen();
    }

  }

}
