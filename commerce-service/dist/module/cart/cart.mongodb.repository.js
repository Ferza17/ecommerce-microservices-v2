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
var CartMongodbRepository_1;
Object.defineProperty(exports, "__esModule", { value: true });
exports.CartMongodbRepository = void 0;
const common_1 = require("@nestjs/common");
const cart_1 = require("../../model/mongo/cart");
const mongoose_1 = require("@nestjs/mongoose");
const mongoose_2 = require("mongoose");
const cartMessage_1 = require("../../model/rpc/gen/commerce/v1/cartMessage");
const jaeger_telemetry_service_1 = require("../../infrastructure/telemetry/jaeger.telemetry.service");
let CartMongodbRepository = CartMongodbRepository_1 = class CartMongodbRepository {
    cartModel;
    otel;
    logger = new common_1.Logger(CartMongodbRepository_1.name);
    constructor(cartModel, otel) {
        this.cartModel = cartModel;
        this.otel = otel;
    }
    async CreateCartItem(requestId, request, context) {
        const span = this.otel.tracer('Repository.CreateCartItem', context);
        try {
            const resp = new this.cartModel({
                ...request,
                created_at: new Date(),
                updated_at: new Date(),
            });
            await resp.save();
            return resp._id.toString();
        }
        catch (e) {
            this.logger.error(`requestId: ${requestId} , error: ${e.message}`);
            span.recordException(e);
            return null;
        }
        finally {
            span.end();
        }
    }
    async FindCartItemById(requestId, id, context) {
        const span = this.otel.tracer('Repository.FindCartItemById', context);
        try {
            return await this.cartModel.findOne({ _id: id });
        }
        catch (e) {
            this.logger.error(`requestId: ${requestId} , error: ${e.message}`);
            span.recordException(e);
            return null;
        }
        finally {
            span.end();
        }
    }
    async FindCartItemByProductId(requestId, productId, context) {
        const span = this.otel.tracer('Repository.FindCartItemByProductId', context);
        try {
            return await this.cartModel.findOne({ productId: productId });
        }
        catch (e) {
            this.logger.error(`requestId: ${requestId} , error: ${e.message}`);
            span.recordException(e);
            return null;
        }
        finally {
            span.end();
        }
    }
    async FindCartItemsWithPagination(requestId, request, context) {
        const span = this.otel.tracer('Repository.FindCartItemsWithPagination', context);
        try {
            const skip = (request.page - 1) * request.limit;
            const query = {};
            if (request.userId && request.userId.trim() !== '') {
                query.userId = request.userId;
            }
            if (request.productIds && request.productIds.length > 0) {
                query.productId = {
                    '$in': request.productIds,
                };
            }
            const cartItems = await this.cartModel
                .find(query)
                .skip(skip)
                .limit(request.limit)
                .exec();
            const total = await this.cartModel.countDocuments().exec();
            return cartMessage_1.FindCartItemsWithPaginationResponse.create({
                items: cartItems,
                total: total,
                page: request.page,
                limit: request.limit,
            });
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
    async UpdateCartItemById(requestId, id, request, context) {
        const span = this.otel.tracer('Repository.UpdateCartItemById', context);
        try {
            return await this.cartModel.findByIdAndUpdate(id, request);
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
    async DeleteCartItemById(requestId, id, context) {
        const span = this.otel.tracer('Repository.DeleteCartItemById', context);
        try {
            await this.cartModel.findByIdAndDelete(id);
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
exports.CartMongodbRepository = CartMongodbRepository;
exports.CartMongodbRepository = CartMongodbRepository = CartMongodbRepository_1 = __decorate([
    (0, common_1.Injectable)(),
    __param(0, (0, mongoose_1.InjectModel)(cart_1.CartItem.name)),
    __metadata("design:paramtypes", [mongoose_2.Model,
        jaeger_telemetry_service_1.JaegerTelemetryService])
], CartMongodbRepository);
//# sourceMappingURL=cart.mongodb.repository.js.map