import { Logger } from '@nestjs/common';
import { RabbitmqOptions } from './options';
import { NestFactory } from '@nestjs/core';
import { RabbitmqModule } from './rabbitmq.module';
import { MicroserviceOptions, Transport } from '@nestjs/microservices';

export class RabbitmqConsumer {
  private readonly logger = new Logger(RabbitmqConsumer.name);

  constructor(
    private readonly rmqOptions: RabbitmqOptions,
  ) {
  }

  async Serve() {
    const options = this.rmqOptions.getRabbitmqOptions();
    const app = await NestFactory.createMicroservice<MicroserviceOptions>(RabbitmqModule, options);
    await app.listen();
  }

}
