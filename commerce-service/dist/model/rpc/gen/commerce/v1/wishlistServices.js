"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.WishlistServiceClient = exports.WishlistServiceService = exports.protobufPackage = void 0;
const grpc_js_1 = require("@grpc/grpc-js");
const wishlistMessage_1 = require("./wishlistMessage");
exports.protobufPackage = "commerce_v1";
exports.WishlistServiceService = {
    createWishlistItem: {
        path: "/commerce_v1.WishlistService/CreateWishlistItem",
        requestStream: false,
        responseStream: false,
        requestSerialize: (value) => Buffer.from(wishlistMessage_1.CreateWishlistItemRequest.encode(value).finish()),
        requestDeserialize: (value) => wishlistMessage_1.CreateWishlistItemRequest.decode(value),
        responseSerialize: (value) => Buffer.from(wishlistMessage_1.CreateWishlistItemResponse.encode(value).finish()),
        responseDeserialize: (value) => wishlistMessage_1.CreateWishlistItemResponse.decode(value),
    },
    findWishlistItemWithPagination: {
        path: "/commerce_v1.WishlistService/FindWishlistItemWithPagination",
        requestStream: false,
        responseStream: false,
        requestSerialize: (value) => Buffer.from(wishlistMessage_1.FindWishlistItemWithPaginationRequest.encode(value).finish()),
        requestDeserialize: (value) => wishlistMessage_1.FindWishlistItemWithPaginationRequest.decode(value),
        responseSerialize: (value) => Buffer.from(wishlistMessage_1.FindWishlistItemWithPaginationResponse.encode(value).finish()),
        responseDeserialize: (value) => wishlistMessage_1.FindWishlistItemWithPaginationResponse.decode(value),
    },
    deleteWishlistItemById: {
        path: "/commerce_v1.WishlistService/DeleteWishlistItemById",
        requestStream: false,
        responseStream: false,
        requestSerialize: (value) => Buffer.from(wishlistMessage_1.DeleteWishlistItemByIdRequest.encode(value).finish()),
        requestDeserialize: (value) => wishlistMessage_1.DeleteWishlistItemByIdRequest.decode(value),
        responseSerialize: (value) => Buffer.from(wishlistMessage_1.DeleteWishlistItemByIdResponse.encode(value).finish()),
        responseDeserialize: (value) => wishlistMessage_1.DeleteWishlistItemByIdResponse.decode(value),
    },
};
exports.WishlistServiceClient = (0, grpc_js_1.makeGenericClientConstructor)(exports.WishlistServiceService, "commerce_v1.WishlistService");
//# sourceMappingURL=wishlistServices.js.map