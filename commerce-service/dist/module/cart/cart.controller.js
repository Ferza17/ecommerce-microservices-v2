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
var CartController_1;
Object.defineProperty(exports, "__esModule", { value: true });
exports.CartController = void 0;
const common_1 = require("@nestjs/common");
const microservices_1 = require("@nestjs/microservices");
const cart_service_1 = require("./cart.service");
const grpc_js_1 = require("@grpc/grpc-js");
const header_1 = require("../../enum/header");
const cartMessage_1 = require("../../model/rpc/gen/commerce/v1/cartMessage");
const jaeger_telemetry_service_1 = require("../../infrastructure/telemetry/jaeger.telemetry.service");
const api_1 = require("@opentelemetry/api");
let CartController = CartController_1 = class CartController {
    cartService;
    otel;
    logger = new common_1.Logger(CartController_1.name);
    constructor(cartService, otel) {
        this.cartService = cartService;
        this.otel = otel;
    }
    async createCartItem(req, metadata) {
        const requestId = metadata.get(header_1.Header.X_REQUEST_ID)[0].toString();
        const spanCtx = api_1.propagation.extract(api_1.context.active(), metadata.getMap());
        const span = this.otel.tracer('Controller.ConsumeCreateCartItem', spanCtx);
        try {
            return await this.cartService.CreateCartItem(requestId, req);
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
    async findCartItemById(req, metadata) {
        const requestId = metadata.get(header_1.Header.X_REQUEST_ID)[0].toString();
        const spanCtx = api_1.propagation.extract(api_1.context.active(), metadata.getMap());
        const span = this.otel.tracer('Controller.FindCartItemById', spanCtx);
        try {
            return await this.cartService.FindCartItemById(requestId, req);
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
    async FindCartItemsWithPagination(req, metadata) {
        const requestId = metadata.get(header_1.Header.X_REQUEST_ID)[0].toString();
        const spanCtx = api_1.propagation.extract(api_1.context.active(), metadata.getMap());
        const span = this.otel.tracer('Controller.FindCartItemsWithPagination', spanCtx);
        try {
            return await this.cartService.FindCartItemsWithPagination(requestId, req);
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
    async DeleteCartItemById(req, metadata) {
        const requestId = metadata.get(header_1.Header.X_REQUEST_ID)[0].toString();
        const spanCtx = api_1.propagation.extract(api_1.context.active(), metadata.getMap());
        const span = this.otel.tracer('Controller.ConsumeCreateCartItem', spanCtx);
        try {
            return await this.cartService.DeleteCartItemById(requestId, req);
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
exports.CartController = CartController;
__decorate([
    (0, microservices_1.GrpcMethod)('CartService', 'CreateCartItem'),
    __metadata("design:type", Function),
    __metadata("design:paramtypes", [Object, grpc_js_1.Metadata]),
    __metadata("design:returntype", Promise)
], CartController.prototype, "createCartItem", null);
__decorate([
    (0, microservices_1.GrpcMethod)('CartService', 'FindCartItemById'),
    __metadata("design:type", Function),
    __metadata("design:paramtypes", [Object, grpc_js_1.Metadata]),
    __metadata("design:returntype", Promise)
], CartController.prototype, "findCartItemById", null);
__decorate([
    (0, microservices_1.GrpcMethod)('CartService', 'FindCartItemsWithPagination'),
    __metadata("design:type", Function),
    __metadata("design:paramtypes", [Object, grpc_js_1.Metadata]),
    __metadata("design:returntype", Promise)
], CartController.prototype, "FindCartItemsWithPagination", null);
__decorate([
    (0, microservices_1.GrpcMethod)('CartService', 'DeleteCartItemById'),
    __metadata("design:type", Function),
    __metadata("design:paramtypes", [Object, grpc_js_1.Metadata]),
    __metadata("design:returntype", Promise)
], CartController.prototype, "DeleteCartItemById", null);
exports.CartController = CartController = CartController_1 = __decorate([
    (0, common_1.Controller)('cart-controller'),
    __metadata("design:paramtypes", [cart_service_1.CartService,
        jaeger_telemetry_service_1.JaegerTelemetryService])
], CartController);
//# sourceMappingURL=cart.controller.js.map