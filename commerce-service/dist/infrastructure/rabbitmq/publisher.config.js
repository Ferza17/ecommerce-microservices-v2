"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.ClientModuleAsyncConfig = void 0;
const microservices_1 = require("@nestjs/microservices");
const consul_service_1 = require("../../config/consul.service");
const consul_module_1 = require("../../config/consul.module");
const ClientModuleAsyncConfig = (exchange, queue) => ({
    name: queue,
    imports: [consul_module_1.ConsulModule],
    inject: [consul_service_1.ConsulService],
    useFactory: async (configService, queue) => ({
        transport: microservices_1.Transport.RMQ,
        options: {
            urls: [`amqp://${await configService.get("/broker/rabbitmq/RABBITMQ_USERNAME")}:${await configService.get("/broker/rabbitmq/RABBITMQ_PASSWORD")}@${await configService.get("/broker/rabbitmq/RABBITMQ_HOST")}:${await configService.get("/broker/rabbitmq/RABBITMQ_PORT")}`],
            queue: queue,
            queueOptions: {
                durable: true,
            },
            routingKey: queue,
            exchange: exchange,
            exchangeType: 'direct',
            noDelay: true,
            gracefulShutdown: true,
        },
    }),
});
exports.ClientModuleAsyncConfig = ClientModuleAsyncConfig;
//# sourceMappingURL=publisher.config.js.map