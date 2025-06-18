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
var WishlistService_1;
Object.defineProperty(exports, "__esModule", { value: true });
exports.WishlistService = void 0;
const common_1 = require("@nestjs/common");
const product_rpc_service_1 = require("../../infrastructure/rpc/product.rpc.service");
const user_rpc_service_1 = require("../../infrastructure/rpc/user.rpc.service");
const rabbitmq_1 = require("../../infrastructure/rabbitmq/rabbitmq");
let WishlistService = WishlistService_1 = class WishlistService {
    rabbitMQInfrastructure;
    productRpcService;
    userRpcService;
    logger = new common_1.Logger(WishlistService_1.name);
    constructor(rabbitMQInfrastructure, productRpcService, userRpcService) {
        this.rabbitMQInfrastructure = rabbitMQInfrastructure;
        this.productRpcService = productRpcService;
        this.userRpcService = userRpcService;
    }
};
exports.WishlistService = WishlistService;
exports.WishlistService = WishlistService = WishlistService_1 = __decorate([
    (0, common_1.Injectable)(),
    __metadata("design:paramtypes", [rabbitmq_1.RabbitmqInfrastructure,
        product_rpc_service_1.ProductRpcService,
        user_rpc_service_1.UserRpcService])
], WishlistService);
//# sourceMappingURL=wishlist.service.js.map