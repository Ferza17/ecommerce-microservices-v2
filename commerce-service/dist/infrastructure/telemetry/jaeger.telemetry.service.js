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
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
var JaegerTelemetryService_1;
Object.defineProperty(exports, "__esModule", { value: true });
exports.JaegerTelemetryService = void 0;
const common_1 = require("@nestjs/common");
const service_1 = require("../../enum/service");
const exporter_jaeger_1 = require("@opentelemetry/exporter-jaeger");
const api_1 = __importDefault(require("@opentelemetry/api"));
const sdk_trace_base_1 = require("@opentelemetry/sdk-trace-base");
const semantic_conventions_1 = require("@opentelemetry/semantic-conventions");
const context_async_hooks_1 = require("@opentelemetry/context-async-hooks");
const core_1 = require("@opentelemetry/core");
const resources_1 = require("@opentelemetry/resources");
const consul_service_1 = require("../../config/consul.service");
let JaegerTelemetryService = JaegerTelemetryService_1 = class JaegerTelemetryService {
    configService;
    logger = new common_1.Logger(JaegerTelemetryService_1.name);
    tt;
    constructor(configService) {
        this.configService = configService;
    }
    async onModuleInit() {
        const exporter = new exporter_jaeger_1.JaegerExporter({
            endpoint: `${await this.configService.get(`/telemetry/jaeger/JAEGER_TELEMETRY_HOST`)}:${await this.configService.get(`/telemetry/jaeger/JAEGER_TELEMETRY_HOST`)}/api/traces`,
        });
        const provider = new sdk_trace_base_1.BasicTracerProvider({
            resource: new resources_1.Resource({
                [semantic_conventions_1.ATTR_SERVICE_NAME]: service_1.Service.CommerceService.toString(),
            }),
        });
        provider.addSpanProcessor(new sdk_trace_base_1.SimpleSpanProcessor(exporter));
        api_1.default.context.setGlobalContextManager(new context_async_hooks_1.AsyncLocalStorageContextManager());
        api_1.default.propagation.setGlobalPropagator(new core_1.CompositePropagator({
            propagators: [
                new core_1.W3CTraceContextPropagator(),
                new core_1.W3CBaggagePropagator()
            ],
        }));
        api_1.default.trace.setGlobalTracerProvider(provider);
        this.tt = api_1.default.trace.getTracer(service_1.Service.CommerceService.toString());
    }
    tracer(operationName, ctx) {
        const spanOptions = {};
        let span;
        if (ctx) {
            span = this.tt.startSpan(operationName, spanOptions, ctx);
        }
        else {
            span = this.tt.startSpan(operationName, spanOptions, api_1.default.context.active());
        }
        return span;
    }
};
exports.JaegerTelemetryService = JaegerTelemetryService;
exports.JaegerTelemetryService = JaegerTelemetryService = JaegerTelemetryService_1 = __decorate([
    (0, common_1.Injectable)(),
    __metadata("design:paramtypes", [consul_service_1.ConsulService])
], JaegerTelemetryService);
//# sourceMappingURL=jaeger.telemetry.service.js.map