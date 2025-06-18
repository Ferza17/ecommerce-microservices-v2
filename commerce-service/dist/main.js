"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const rpc_1 = require("./server/rpc/rpc");
const config_1 = require("@nestjs/config");
const options_1 = require("./server/rpc/options");
const rabbitmq_1 = require("./server/rabbitmq/rabbitmq");
const consul_service_1 = require("./config/consul.service");
function bootstrap() {
    const configService = new config_1.ConfigService();
    const consulConfig = new consul_service_1.ConsulService(configService);
    const rmqConsumer = new rabbitmq_1.RabbitmqConsumer(consulConfig);
    rmqConsumer.Serve();
    const grpcClientOptions = new options_1.GrpcClientOptions(consulConfig);
    const grpcServer = new rpc_1.GrpcServer(grpcClientOptions);
    grpcServer.Serve();
}
bootstrap();
//# sourceMappingURL=main.js.map