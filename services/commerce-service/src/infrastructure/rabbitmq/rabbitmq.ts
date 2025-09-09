import { Injectable, Logger } from '@nestjs/common';
import { Queue } from '../../enum/queue';
import { EventStore } from '../../model/rpc/gen/v1/event/model';
import { Header } from '../../enum/header';
import { AmqpConnection } from '@golevelup/nestjs-rabbitmq';
import { Exchange } from '../../enum/exchange';
import { JaegerTelemetryService } from '../telemetry/jaeger.telemetry.service';
import { Context } from '@opentelemetry/api';

@Injectable()
export class RabbitmqInfrastructure {
  private readonly logger = new Logger(RabbitmqInfrastructure.name);

  constructor(
    private readonly amqpConnection: AmqpConnection,
    private readonly otel: JaegerTelemetryService,
  ) {
  }

  async publishEventCreated(requestId: string, event: EventStore, context?: Context) {
    const span = this.otel.tracer('Infrastructure.publishEventCreated', context);
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
      span.recordException(e);
      throw e;
    } finally {
      span.end();
    }
  }
}