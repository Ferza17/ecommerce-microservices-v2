import { Injectable } from '@nestjs/common';
import { ConfigService } from '@nestjs/config';
import { RmqOptions, Serializer, Transport } from '@nestjs/microservices';
import { Queue } from '../../enum/queue';
import { Exchange } from '../../enum/exchange';


@Injectable()
export class RabbitmqOptions {
  constructor(
    private readonly configService: ConfigService,
    private queue: Queue,
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
        exchange: Exchange.CommerceExchange,
        exchangeType: 'direct',
        routingKey: this.queue,
      },
    });
  }


}