import { Logger } from '@nestjs/common';
import { RabbitmqOptions } from './consumer.config';
import { NestFactory } from '@nestjs/core';
import { RabbitmqModule } from './rabbitmq.module';
import { MicroserviceOptions, Transport } from '@nestjs/microservices';
import { Queue } from '../../enum/queue';
import { ConsulService } from '../../config/consul.service';

export class RabbitmqConsumer {
  private readonly logger = new Logger(RabbitmqConsumer.name);

  constructor(
    private readonly consulConfig: ConsulService,
  ) {
  }

  async Serve() {
    const options = [
      {
        queue: Queue.CART_CREATED,
        option: await new RabbitmqOptions(this.consulConfig, Queue.CART_CREATED).getRabbitmqOptions(),
      },
      {
        queue: Queue.CART_UPDATED,
        option: await new RabbitmqOptions(this.consulConfig, Queue.CART_UPDATED).getRabbitmqOptions(),
      },
      {
        queue: Queue.CART_DELETED,
        option: await new RabbitmqOptions(this.consulConfig, Queue.CART_DELETED).getRabbitmqOptions(),
      },
    ];

    for (const option of options) {
      this.logger.log(`Rabbitmq Consumer ${option.queue} is running...`);
      const app = await NestFactory.createMicroservice<MicroserviceOptions>(RabbitmqModule, option.option);
      app.listen();
    }

  }

}
