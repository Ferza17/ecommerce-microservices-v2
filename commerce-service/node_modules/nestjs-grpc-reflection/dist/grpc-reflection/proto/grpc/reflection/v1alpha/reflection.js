"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.SERVER_REFLECTION_SERVICE_NAME = exports.ServerReflectionControllerMethods = exports.GRPC_REFLECTION_V1ALPHA_PACKAGE_NAME = exports.protobufPackage = void 0;
const microservices_1 = require("@nestjs/microservices");
exports.protobufPackage = 'grpc.reflection.v1alpha';
exports.GRPC_REFLECTION_V1ALPHA_PACKAGE_NAME = 'grpc.reflection.v1alpha';
function ServerReflectionControllerMethods() {
    return function (constructor) {
        const grpcMethods = [];
        for (const method of grpcMethods) {
            const descriptor = Reflect.getOwnPropertyDescriptor(constructor.prototype, method);
            (0, microservices_1.GrpcMethod)('ServerReflection', method)(constructor.prototype[method], method, descriptor);
        }
        const grpcStreamMethods = ['serverReflectionInfo'];
        for (const method of grpcStreamMethods) {
            const descriptor = Reflect.getOwnPropertyDescriptor(constructor.prototype, method);
            (0, microservices_1.GrpcStreamMethod)('ServerReflection', method)(constructor.prototype[method], method, descriptor);
        }
    };
}
exports.ServerReflectionControllerMethods = ServerReflectionControllerMethods;
exports.SERVER_REFLECTION_SERVICE_NAME = 'ServerReflection';
//# sourceMappingURL=reflection.js.map