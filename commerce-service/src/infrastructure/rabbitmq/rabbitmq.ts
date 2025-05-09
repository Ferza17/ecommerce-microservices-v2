import { Injectable, Logger } from '@nestjs/common';
import { Queue } from '../../enum/queue';
import { EventStore } from '../../model/rpc/eventStoreMessage';
import { Header } from '../../enum/header';
import { AmqpConnection } from '@golevelup/nestjs-rabbitmq';
import { Exchange } from '../../enum/exchange';

@Injectable()
export class RabbitmqInfrastructure {
  private readonly logger = new Logger(RabbitmqInfrastructure.name);

  constructor(private readonly amqpConnection: AmqpConnection) {
  }

  async publishEventCreated(requestId: string, event: EventStore) {
    try {
      await this.amqpConnection.publish(Exchange.EventExchange.toString(), Queue.EVENT_CREATED.toString(), event, {
        headers: {
          [Header.X_REQUEST_ID]: requestId,
        },
        contentType: 'application/json',
        deliveryMode: 1,
        timestamp: new Date().getTime(),
        persistent: true,
      });
    } catch (e) {
      throw e;
    }
  }
}