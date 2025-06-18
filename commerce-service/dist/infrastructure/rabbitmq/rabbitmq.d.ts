import { EventStore } from '../../model/rpc/gen/event/v1/eventStoreMessage';
import { AmqpConnection } from '@golevelup/nestjs-rabbitmq';
import { JaegerTelemetryService } from '../telemetry/jaeger.telemetry.service';
import { Context } from '@opentelemetry/api';
export declare class RabbitmqInfrastructure {
    private readonly amqpConnection;
    private readonly otel;
    private readonly logger;
    constructor(amqpConnection: AmqpConnection, otel: JaegerTelemetryService);
    publishEventCreated(requestId: string, event: EventStore, context?: Context): Promise<void>;
}
