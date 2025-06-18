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
Object.defineProperty(exports, "__esModule", { value: true });
exports.RabbitmqOptions = void 0;
const common_1 = require("@nestjs/common");
const microservices_1 = require("@nestjs/microservices");
const queue_1 = require("../../enum/queue");
const exchange_1 = require("../../enum/exchange");
const consul_service_1 = require("../../config/consul.service");
let RabbitmqOptions = class RabbitmqOptions {
    consulConfig;
    queue;
    constructor(consulConfig, queue) {
        this.consulConfig = consulConfig;
        this.queue = queue;
    }
    async getRabbitmqOptions() {
        const rmqUsername = await this.consulConfig.get('/broker/rabbitmq/RABBITMQ_USERNAME');
        const rmqPassword = await this.consulConfig.get('/broker/rabbitmq/RABBITMQ_PASSWORD');
        const rmqHost = await this.consulConfig.get('/broker/rabbitmq/RABBITMQ_HOST');
        const rmqPort = await this.consulConfig.get('/broker/rabbitmq/RABBITMQ_PORT');
        const url = `amqp://${rmqUsername}:${rmqPassword}@${rmqHost}:${rmqPort}`;
        return ({
            transport: microservices_1.Transport.RMQ,
            options: {
                urls: [url],
                queue: this.queue,
                queueOptions: {
                    durable: true,
                },
                exchange: exchange_1.Exchange.CommerceExchange,
                exchangeType: 'direct',
                routingKey: this.queue,
            },
        });
    }
};
exports.RabbitmqOptions = RabbitmqOptions;
exports.RabbitmqOptions = RabbitmqOptions = __decorate([
    (0, common_1.Injectable)(),
    __metadata("design:paramtypes", [consul_service_1.ConsulService, String])
], RabbitmqOptions);
//# sourceMappingURL=consumer.config.js.map