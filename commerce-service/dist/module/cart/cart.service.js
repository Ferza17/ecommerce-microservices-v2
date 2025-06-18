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
var CartService_1;
Object.defineProperty(exports, "__esModule", { value: true });
exports.CartService = void 0;
const common_1 = require("@nestjs/common");
const cart_mongodb_repository_1 = require("./cart.mongodb.repository");
const rabbitmq_1 = require("../../infrastructure/rabbitmq/rabbitmq");
const sagaStatus_1 = require("../../enum/sagaStatus");
const service_1 = require("../../enum/service");
const queue_1 = require("../../enum/queue");
const product_rpc_service_1 = require("../../infrastructure/rpc/product.rpc.service");
const productMessage_1 = require("../../model/rpc/gen/product/v1/productMessage");
const user_rpc_service_1 = require("../../infrastructure/rpc/user.rpc.service");
const jaeger_telemetry_service_1 = require("../../infrastructure/telemetry/jaeger.telemetry.service");
const userMessage_1 = require("../../model/rpc/gen/user/v1/userMessage");
let CartService = CartService_1 = class CartService {
    cartItemRepository;
    rabbitMQInfrastructure;
    productRpcService;
    userRpcService;
    otel;
    logger = new common_1.Logger(CartService_1.name);
    constructor(cartItemRepository, rabbitMQInfrastructure, productRpcService, userRpcService, otel) {
        this.cartItemRepository = cartItemRepository;
        this.rabbitMQInfrastructure = rabbitMQInfrastructure;
        this.productRpcService = productRpcService;
        this.userRpcService = userRpcService;
        this.otel = otel;
    }
    async CreateCartItem(requestId, req, context) {
        const span = this.otel.tracer('Service.CreateCartItem', context);
        let id = '';
        let event = {
            service: service_1.Service.CommerceService.toString(),
            eventType: queue_1.Queue.CART_CREATED.toString(),
            requestId: requestId,
            payload: req,
            updatedAt: new Date(),
            createdAt: new Date(),
            id: '',
            status: '',
        };
        try {
            const user = await this.userRpcService.findUserById(requestId, userMessage_1.FindUserByIdRequest.create({ id: req.userId }), context);
            if (user === null) {
                throw new Error('User not found');
            }
            const product = await this.productRpcService.findProductById(requestId, productMessage_1.FindProductByIdRequest.create({ id: req.productId }), context);
            if (product === null) {
                throw new Error('Product not found');
            }
            if (product.stock < req.qty) {
                throw new Error('Product stock is not enough');
            }
            const cartItem = await this.cartItemRepository.FindCartItemByProductId(requestId, req.productId);
            if (cartItem !== null) {
                await this.cartItemRepository.UpdateCartItemById(requestId, cartItem._id.toString(), {
                    price: product.price * (req.qty + cartItem.qty),
                    qty: req.qty + cartItem.qty,
                    userId: user.id,
                    productId: req.productId,
                }, context);
                id = cartItem._id.toString() || '';
            }
            else {
                req.price = product.price * req.qty;
                const result = await this.cartItemRepository.CreateCartItem(requestId, req, context);
                if (result !== null) {
                }
            }
            event.status = sagaStatus_1.SagaStatus.SUCCESS.toString();
            event.previousState = cartItem || undefined;
            await this.rabbitMQInfrastructure.publishEventCreated(requestId, event, context);
            return { id: id };
        }
        catch (e) {
            event.status = sagaStatus_1.SagaStatus.FAILED.toString();
            await this.rabbitMQInfrastructure.publishEventCreated(requestId, event, context);
            span.recordException(e);
            throw e;
        }
        finally {
            span.end();
        }
    }
    async FindCartItemById(requestId, req, context) {
        const span = this.otel.tracer('Service.FindCartItemById', context);
        try {
            const cartItem = await this.cartItemRepository.FindCartItemById(requestId, req.id);
            if (!cartItem) {
                throw new Error('Cart item not found');
            }
            return {
                id: cartItem._id.toString(),
                productId: cartItem.productId,
                userId: cartItem.userId,
                qty: cartItem.qty,
                price: cartItem.price,
                cratedAt: cartItem.created_at,
                updatedAt: cartItem.updated_at,
            };
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
    async FindCartItemsWithPagination(requestId, req, context) {
        const span = this.otel.tracer('Service.FindCartItemsWithPagination', context);
        try {
            return await this.cartItemRepository.FindCartItemsWithPagination(requestId, req);
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
    async UpdateCartItemByIdRequest(requestId, req, context) {
        const span = this.otel.tracer('Service.UpdateCartItemByIdRequest', context);
        let event = {
            service: service_1.Service.CommerceService.toString(),
            eventType: queue_1.Queue.CART_UPDATED.toString(),
            requestId: requestId,
            payload: req,
            updatedAt: new Date(),
            createdAt: new Date(),
            id: '',
            status: '',
        };
        try {
            const user = await this.userRpcService.findUserById(requestId, productMessage_1.FindProductByIdRequest.create({ id: req.userId }));
            if (user === null || user === undefined) {
                throw new Error('User not found');
            }
            const product = await this.productRpcService.findProductById(requestId, productMessage_1.FindProductByIdRequest.create({ id: req.productId }));
            if (product === null) {
                throw new Error('Product not found');
            }
            if (product.stock < req.qty) {
                throw new Error('Product stock is not enough');
            }
            req.price = product.price * req.qty;
            const cartItem = await this.cartItemRepository.FindCartItemById(requestId, req.id);
            if (cartItem === null) {
                throw new Error('Cart item not found');
            }
            await this.cartItemRepository.UpdateCartItemById(requestId, req.id, req);
            event.status = sagaStatus_1.SagaStatus.SUCCESS.toString();
            await this.rabbitMQInfrastructure.publishEventCreated(requestId, event);
            return { id: req.id };
        }
        catch (e) {
            this.logger.error(`requestId: ${requestId} , error: ${e.message}`);
            event.status = sagaStatus_1.SagaStatus.FAILED.toString();
            await this.rabbitMQInfrastructure.publishEventCreated(requestId, event);
            span.recordException(e);
            throw e;
        }
        finally {
            span.end();
        }
    }
    async DeleteCartItemById(requestId, req, context) {
        const span = this.otel.tracer('Service.DeleteCartItemById', context);
        try {
            await this.cartItemRepository.DeleteCartItemById(requestId, req.id);
            return {
                message: 'Cart item deleted successfully',
            };
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
exports.CartService = CartService;
exports.CartService = CartService = CartService_1 = __decorate([
    (0, common_1.Injectable)(),
    __metadata("design:paramtypes", [cart_mongodb_repository_1.CartMongodbRepository,
        rabbitmq_1.RabbitmqInfrastructure,
        product_rpc_service_1.ProductRpcService,
        user_rpc_service_1.UserRpcService,
        jaeger_telemetry_service_1.JaegerTelemetryService])
], CartService);
//# sourceMappingURL=cart.service.js.map