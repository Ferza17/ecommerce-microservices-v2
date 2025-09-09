import { Injectable } from '@nestjs/common';
import { ConfigService } from '@nestjs/config';
import { RmqOptions, Serializer, Transport } from '@nestjs/microservices';
import { Queue } from '../../enum/queue';
import { Exchange } from '../../enum/exchange';
import { ConsulService } from '../../config/consul.service';


@Injectable()
export class RabbitmqOptions {
  constructor(
    private readonly consulConfig: ConsulService,
    private queue: Queue,
  ) {
  }

  async getRabbitmqOptions(): Promise<RmqOptions> {
    const rmqUsername = await this.consulConfig.get('/broker/rabbitmq/RABBITMQ_USERNAME');
    const rmqPassword = await this.consulConfig.get('/broker/rabbitmq/RABBITMQ_PASSWORD');
    const rmqHost = await this.consulConfig.get('/broker/rabbitmq/RABBITMQ_HOST');
    const rmqPort = await this.consulConfig.get('/broker/rabbitmq/RABBITMQ_PORT');
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