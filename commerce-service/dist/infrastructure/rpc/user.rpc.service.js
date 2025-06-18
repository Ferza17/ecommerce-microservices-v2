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
var __param = (this && this.__param) || function (paramIndex, decorator) {
    return function (target, key) { decorator(target, key, paramIndex); }
};
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
var UserRpcService_1;
Object.defineProperty(exports, "__esModule", { value: true });
exports.UserRpcService = void 0;
const common_1 = require("@nestjs/common");
const service_1 = require("../../enum/service");
const grpc_js_1 = require("@grpc/grpc-js");
const header_1 = require("../../enum/header");
const jaeger_telemetry_service_1 = require("../telemetry/jaeger.telemetry.service");
const api_1 = require("@opentelemetry/api");
const opossum_1 = __importDefault(require("opossum"));
const rxjs_1 = require("rxjs");
const circuitOptions_1 = require("../../config/circuitOptions");
let UserRpcService = UserRpcService_1 = class UserRpcService {
    client;
    otel;
    logger = new common_1.Logger(UserRpcService_1.name);
    userService;
    constructor(client, otel) {
        this.client = client;
        this.otel = otel;
    }
    onModuleInit() {
        this.userService = this.client.getService('UserService');
    }
    async findUserById(requestId, req, context) {
        const span = this.otel.tracer('RpcService.findUserById', context);
        try {
            const breakerFn = async (requestId, req, context) => {
                const metadata = new grpc_js_1.Metadata();
                metadata.set(header_1.Header.X_REQUEST_ID, requestId);
                if (context) {
                    api_1.propagation.inject(context, metadata, {
                        set: (metadata, key, value) => metadata.set(key, value),
                    });
                }
                const observableResult = this.userService.findUserById(req, metadata);
                return (0, rxjs_1.lastValueFrom)(observableResult);
            };
            const cb = new opossum_1.default(breakerFn, circuitOptions_1.UserServiceCircuitOptions);
            return await cb.fire(requestId, req, context);
        }
        catch (e) {
            this.logger.error(`requestId: ${requestId} , error: ${e.message}`);
            span.recordException(e);
            throw e;
        }
        finally {
            span.end();
        }
    }
};
exports.UserRpcService = UserRpcService;
exports.UserRpcService = UserRpcService = UserRpcService_1 = __decorate([
    (0, common_1.Injectable)(),
    __param(0, (0, common_1.Inject)(service_1.Service.UserService.toString())),
    __metadata("design:paramtypes", [Object, jaeger_telemetry_service_1.JaegerTelemetryService])
], UserRpcService);
//# sourceMappingURL=user.rpc.service.js.map