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
var CartConsumer_1;
Object.defineProperty(exports, "__esModule", { value: true });
exports.CartConsumer = void 0;
const common_1 = require("@nestjs/common");
const microservices_1 = require("@nestjs/microservices");
const header_1 = require("../../enum/header");
const cartMessage_1 = require("../../model/rpc/gen/commerce/v1/cartMessage");
const cart_service_1 = require("./cart.service");
const queue_1 = require("../../enum/queue");
const jaeger_telemetry_service_1 = require("../../infrastructure/telemetry/jaeger.telemetry.service");
const api_1 = require("@opentelemetry/api");
let CartConsumer = CartConsumer_1 = class CartConsumer {
    cartService;
    otel;
    logger = new common_1.Logger(CartConsumer_1.name);
    constructor(cartService, otel) {
        this.cartService = cartService;
        this.otel = otel;
    }
    async consumeCreateCartItem(data, ctx) {
        const { properties: { headers } } = ctx.getMessage();
        const requestId = headers[header_1.Header.X_REQUEST_ID];
        const spanCtx = api_1.propagation.extract(api_1.context.active(), headers);
        const span = this.otel.tracer('Consumer.ConsumeCreateCartItem', spanCtx);
        try {
            return await this.cartService.CreateCartItem(requestId, data, spanCtx);
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
    async updateCartItemByIdRequest(data, ctx) {
        const { properties: { headers } } = ctx.getMessage();
        const requestId = headers[header_1.Header.X_REQUEST_ID];
        const spanCtx = api_1.propagation.extract(api_1.context.active(), headers);
        const span = this.otel.tracer('Consumer.UpdateCartItemByIdRequest', spanCtx);
        try {
            return await this.cartService.UpdateCartItemByIdRequest(requestId, data);
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
    async deleteCartItemByIdRequest(data, ctx) {
        const { properties: { headers } } = ctx.getMessage();
        const requestId = headers[header_1.Header.X_REQUEST_ID];
        const spanCtx = api_1.propagation.extract(api_1.context.active(), headers);
        const span = this.otel.tracer('Consumer.DeleteCartItemByIdRequest', spanCtx);
        try {
            return await this.cartService.DeleteCartItemById(requestId, data);
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
exports.CartConsumer = CartConsumer;
__decorate([
    (0, microservices_1.MessagePattern)(`${queue_1.Queue.CART_CREATED}`),
    __param(0, (0, microservices_1.Payload)()),
    __param(1, (0, microservices_1.Ctx)()),
    __metadata("design:type", Function),
    __metadata("design:paramtypes", [Object, microservices_1.RmqContext]),
    __metadata("design:returntype", Promise)
], CartConsumer.prototype, "consumeCreateCartItem", null);
__decorate([
    (0, microservices_1.MessagePattern)(`${queue_1.Queue.CART_UPDATED}`),
    __param(0, (0, microservices_1.Payload)()),
    __param(1, (0, microservices_1.Ctx)()),
    __metadata("design:type", Function),
    __metadata("design:paramtypes", [Object, microservices_1.RmqContext]),
    __metadata("design:returntype", Promise)
], CartConsumer.prototype, "updateCartItemByIdRequest", null);
__decorate([
    (0, microservices_1.MessagePattern)(`${queue_1.Queue.CART_DELETED}`),
    __param(0, (0, microservices_1.Payload)()),
    __param(1, (0, microservices_1.Ctx)()),
    __metadata("design:type", Function),
    __metadata("design:paramtypes", [Object, microservices_1.RmqContext]),
    __metadata("design:returntype", Promise)
], CartConsumer.prototype, "deleteCartItemByIdRequest", null);
exports.CartConsumer = CartConsumer = CartConsumer_1 = __decorate([
    (0, common_1.Controller)(),
    __metadata("design:paramtypes", [cart_service_1.CartService,
        jaeger_telemetry_service_1.JaegerTelemetryService])
], CartConsumer);
//# sourceMappingURL=cart.consumer.js.map