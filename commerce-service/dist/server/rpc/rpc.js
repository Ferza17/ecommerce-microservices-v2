"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.GrpcServer = void 0;
const core_1 = require("@nestjs/core");
const rpc_module_1 = require("./rpc.module");
const requestIdInterceptor_service_1 = require("./interceptor/requestIdInterceptor.service");
const common_1 = require("@nestjs/common");
class GrpcServer {
    grpcClientOptions;
    logger = new common_1.Logger(GrpcServer.name);
    constructor(grpcClientOptions) {
        this.grpcClientOptions = grpcClientOptions;
    }
    async Serve() {
        const app = await core_1.NestFactory.createMicroservice(rpc_module_1.RpcServerModule, await this.grpcClientOptions.getGRPCConfig());
        app.useGlobalInterceptors(new requestIdInterceptor_service_1.RequestIdInterceptor());
        await app.listen();
        this.logger.log('GRPC Server is running...');
    }
}
exports.GrpcServer = GrpcServer;
//# sourceMappingURL=rpc.js.map