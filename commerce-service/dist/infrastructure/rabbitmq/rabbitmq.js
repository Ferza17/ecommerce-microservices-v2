"use strict";
var __decorate = (this && this.__decorate) || function (decorators, target, key, desc) {
    var c = arguments.length, r = c < 3 ? target : desc === null ? desc = Object.getOwnPropertyDescriptor(target, key) : desc, d;
    if (typeof Reflect === "object" && typeof Reflect.decorate === "function") r = Reflect.decorate(decorators, target, key, desc);
    else for (var i = decorators.length - 1; i >= 0; i--) if (d = decorators[i]) r = (c < 3 ? d(r) : c > 3 ? d(target, key, r) : d(target, key)) || r;
    return c > 3 && r && Object.defineProperty(target, key, r), r;
};
var __metadata = (this && this.__metadata) || function (k, v) {
    if (typeof Reflect === "object" && typeof Reflect.metadata === "function") return Reflect.metadata(k, v);
};
var RabbitmqInfrastructure_1;
Object.defineProperty(exports, "__esModule", { value: true });
exports.RabbitmqInfrastructure = void 0;
const common_1 = require("@nestjs/common");
const queue_1 = require("../../enum/queue");
const header_1 = require("../../enum/header");
const nestjs_rabbitmq_1 = require("@golevelup/nestjs-rabbitmq");
const exchange_1 = require("../../enum/exchange");
const jaeger_telemetry_service_1 = require("../telemetry/jaeger.telemetry.service");
let RabbitmqInfrastructure = RabbitmqInfrastructure_1 = class RabbitmqInfrastructure {
    amqpConnection;
    otel;
    logger = new common_1.Logger(RabbitmqInfrastructure_1.name);
    constructor(amqpConnection, otel) {
        this.amqpConnection = amqpConnection;
        this.otel = otel;
    }
    async publishEventCreated(requestId, event, context) {
        const span = this.otel.tracer('Infrastructure.publishEventCreated', context);
        try {
            await this.amqpConnection.publish(exchange_1.Exchange.EventExchange.toString(), queue_1.Queue.EVENT_CREATED.toString(), event, {
                headers: {
                    [header_1.Header.X_REQUEST_ID]: requestId,
                },
                contentType: 'application/json',
                deliveryMode: 1,
                timestamp: new Date().getTime(),
                persistent: true,
            });
        }
        catch (e) {
            span.recordException(e);
            throw e;
        }
        finally {
            span.end();
        }
    }
};
exports.RabbitmqInfrastructure = RabbitmqInfrastructure;
exports.RabbitmqInfrastructure = RabbitmqInfrastructure = RabbitmqInfrastructure_1 = __decorate([
    (0, common_1.Injectable)(),
    __metadata("design:paramtypes", [nestjs_rabbitmq_1.AmqpConnection,
        jaeger_telemetry_service_1.JaegerTelemetryService])
], RabbitmqInfrastructure);
//# sourceMappingURL=rabbitmq.js.map