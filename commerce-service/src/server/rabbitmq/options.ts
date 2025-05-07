import { Injectable } from '@nestjs/common';
import { ConfigService } from '@nestjs/config';
import { RmqOptions, Transport } from '@nestjs/microservices';
import { RoutingKey } from '../../enum/routingKey';


@Injectable()
export class RabbitmqOptions {
  constructor(
    private readonly configService: ConfigService,
    private readonly queue: string,
    private readonly exchange: string,
  ) {
  }

  getRabbitmqOptions(): RmqOptions {
    const rmqUsername = this.configService.get<string>('RABBITMQ_USERNAME');
    const rmqPassword = this.configService.get<string>('RABBITMQ_PASSWORD');
    const rmqHost = this.configService.get<string>('RABBITMQ_HOST');
    const rmqPort = this.configService.get<number>('RABBITMQ_PORT');
    const url = `amqp://${rmqUsername}:${rmqPassword}@${rmqHost}:${rmqPort}`;


    return ({
      transport: Transport.RMQ,
      options: {
        urls: [url],
        queue: this.queue,
        queueOptions: {
          durable: true,
        },
        exchange: this.exchange,
        exchangeType: 'topic',
      },
    });
  }


}