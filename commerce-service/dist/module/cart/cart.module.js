"use strict";
var __decorate = (this && this.__decorate) || function (decorators, target, key, desc) {
    var c = arguments.length, r = c < 3 ? target : desc === null ? desc = Object.getOwnPropertyDescriptor(target, key) : desc, d;
    if (typeof Reflect === "object" && typeof Reflect.decorate === "function") r = Reflect.decorate(decorators, target, key, desc);
    else for (var i = decorators.length - 1; i >= 0; i--) if (d = decorators[i]) r = (c < 3 ? d(r) : c > 3 ? d(target, key, r) : d(target, key)) || r;
    return c > 3 && r && Object.defineProperty(target, key, r), r;
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.CartModule = void 0;
const common_1 = require("@nestjs/common");
const cart_controller_1 = require("./cart.controller");
const cart_service_1 = require("./cart.service");
const cart_mongodb_repository_1 = require("./cart.mongodb.repository");
const cart_consumer_1 = require("./cart.consumer");
const mongoose_1 = require("@nestjs/mongoose");
const cart_1 = require("../../model/mongo/cart");
const mongodbCollection_1 = require("../../enum/mongodbCollection");
const infrastructure_module_1 = require("../../infrastructure/infrastructure.module");
let CartModule = class CartModule {
};
exports.CartModule = CartModule;
exports.CartModule = CartModule = __decorate([
    (0, common_1.Module)({
        imports: [
            infrastructure_module_1.InfrastructureModule,
            mongoose_1.MongooseModule.forFeature([
                { name: cart_1.CartItem.name, schema: cart_1.CartSchema, collection: mongodbCollection_1.MongoDBCollection.CartItemCollection },
            ]),
        ],
        controllers: [cart_controller_1.CartController, cart_consumer_1.CartConsumer],
        providers: [cart_service_1.CartService, cart_mongodb_repository_1.CartMongodbRepository],
    })
], CartModule);
//# sourceMappingURL=cart.module.js.map