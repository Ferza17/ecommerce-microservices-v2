"use strict";
var __decorate = (this && this.__decorate) || function (decorators, target, key, desc) {
    var c = arguments.length, r = c < 3 ? target : desc === null ? desc = Object.getOwnPropertyDescriptor(target, key) : desc, d;
    if (typeof Reflect === "object" && typeof Reflect.decorate === "function") r = Reflect.decorate(decorators, target, key, desc);
    else for (var i = decorators.length - 1; i >= 0; i--) if (d = decorators[i]) r = (c < 3 ? d(r) : c > 3 ? d(target, key, r) : d(target, key)) || r;
    return c > 3 && r && Object.defineProperty(target, key, r), r;
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.InfrastructureModule = void 0;
const common_1 = require("@nestjs/common");
const rabbitmq_1 = require("./rabbitmq/rabbitmq");
const config_1 = require("@nestjs/config");
const microservices_1 = require("@nestjs/microservices");
const fast_glob_1 = require("fast-glob");
const path_1 = require("path");
const service_1 = require("../enum/service");
const product_rpc_service_1 = require("./rpc/product.rpc.service");
const user_rpc_service_1 = require("./rpc/user.rpc.service");
const jaeger_telemetry_service_1 = require("./telemetry/jaeger.telemetry.service");
const consul_module_1 = require("../config/consul.module");
const consul_service_1 = require("../config/consul.service");
const configRoot_1 = require("../config/configRoot");
const projectRoot = (0, path_1.join)(__dirname, '../../');
let InfrastructureModule = class InfrastructureModule {
};
exports.InfrastructureModule = InfrastructureModule;
exports.InfrastructureModule = InfrastructureModule = __decorate([
    (0, common_1.Module)({
        imports: [
            config_1.ConfigModule.forRoot(),
            microservices_1.ClientsModule.registerAsync([
                {
                    name: service_1.Service.ProductService.toString(),
                    imports: [consul_module_1.ConsulModule],
                    inject: [consul_service_1.ConsulService],
                    useFactory: async (config) => ({
                        transport: microservices_1.Transport.GRPC,
                        options: {
                            url: `${await config.get('/services/product/RPC_HOST')}:${await config.get('/services/product/RPC_PORT')}`,
                            package: `product_v1`,
                            protoPath: fast_glob_1.glob.sync?.(['proto/**/*.proto'], {
                                cwd: projectRoot,
                                absolute: true,
                            }),
                            loader: {
                                includeDirs: [
                                    (0, path_1.join)(projectRoot, 'proto'),
                                ],
                                oneofs: true,
                            },
                        },
                    }),
                },
            ]),
            microservices_1.ClientsModule.registerAsync([
                {
                    name: service_1.Service.UserService.toString(),
                    imports: [consul_module_1.ConsulModule],
                    inject: [consul_service_1.ConsulService],
                    useFactory: async (configService) => ({
                        transport: microservices_1.Transport.GRPC,
                        options: {
                            url: `${await configService.get('/services/user/RPC_HOST')}:${await configService.get('/services/user/RPC_PORT')}`,
                            package: 'user_v1',
                            protoPath: fast_glob_1.glob.sync?.(['proto/**/*.proto'], {
                                cwd: projectRoot,
                                absolute: true,
                            }),
                            loader: {
                                includeDirs: [
                                    (0, path_1.join)(projectRoot, 'proto'),
                                ],
                                oneofs: true,
                            },
                        },
                    }),
                },
            ]),
            configRoot_1.RabbitMQRootAsync,
            consul_module_1.ConsulModule,
        ],
        providers: [rabbitmq_1.RabbitmqInfrastructure, product_rpc_service_1.ProductRpcService, user_rpc_service_1.UserRpcService, jaeger_telemetry_service_1.JaegerTelemetryService],
        exports: [rabbitmq_1.RabbitmqInfrastructure, product_rpc_service_1.ProductRpcService, user_rpc_service_1.UserRpcService, jaeger_telemetry_service_1.JaegerTelemetryService],
    })
], InfrastructureModule);
//# sourceMappingURL=infrastructure.module.js.map