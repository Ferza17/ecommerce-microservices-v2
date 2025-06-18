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
exports.GrpcClientOptions = void 0;
const nestjs_grpc_reflection_1 = require("nestjs-grpc-reflection");
const microservices_1 = require("@nestjs/microservices");
const fast_glob_1 = require("fast-glob");
const path_1 = require("path");
const common_1 = require("@nestjs/common");
const consul_service_1 = require("../../config/consul.service");
let GrpcClientOptions = class GrpcClientOptions {
    consulConfig;
    constructor(consulConfig) {
        this.consulConfig = consulConfig;
    }
    async getGRPCConfig() {
        const rpcHost = await this.consulConfig.get('/services/commerce/RPC_HOST');
        const rpcPort = await this.consulConfig.get('/services/commerce/RPC_PORT') || '5000';
        const projectRoot = (0, path_1.join)(__dirname, '../../../');
        return (0, nestjs_grpc_reflection_1.addReflectionToGrpcConfig)({
            transport: microservices_1.Transport.GRPC,
            options: {
                url: `${rpcHost}:${rpcPort}`,
                package: `commerce_v1`,
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
        });
    }
};
exports.GrpcClientOptions = GrpcClientOptions;
exports.GrpcClientOptions = GrpcClientOptions = __decorate([
    (0, common_1.Injectable)(),
    __metadata("design:paramtypes", [consul_service_1.ConsulService])
], GrpcClientOptions);
//# sourceMappingURL=options.js.map