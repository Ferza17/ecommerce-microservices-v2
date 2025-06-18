"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.CartServiceClient = exports.CartServiceService = exports.protobufPackage = void 0;
const grpc_js_1 = require("@grpc/grpc-js");
const cartMessage_1 = require("./cartMessage");
exports.protobufPackage = "commerce_v1";
exports.CartServiceService = {
    createCartItem: {
        path: "/commerce_v1.CartService/CreateCartItem",
        requestStream: false,
        responseStream: false,
        requestSerialize: (value) => Buffer.from(cartMessage_1.CreateCartItemRequest.encode(value).finish()),
        requestDeserialize: (value) => cartMessage_1.CreateCartItemRequest.decode(value),
        responseSerialize: (value) => Buffer.from(cartMessage_1.CreateCartItemResponse.encode(value).finish()),
        responseDeserialize: (value) => cartMessage_1.CreateCartItemResponse.decode(value),
    },
    findCartItemById: {
        path: "/commerce_v1.CartService/FindCartItemById",
        requestStream: false,
        responseStream: false,
        requestSerialize: (value) => Buffer.from(cartMessage_1.FindCartItemByIdRequest.encode(value).finish()),
        requestDeserialize: (value) => cartMessage_1.FindCartItemByIdRequest.decode(value),
        responseSerialize: (value) => Buffer.from(cartMessage_1.CartItem.encode(value).finish()),
        responseDeserialize: (value) => cartMessage_1.CartItem.decode(value),
    },
    findCartItemsWithPagination: {
        path: "/commerce_v1.CartService/FindCartItemsWithPagination",
        requestStream: false,
        responseStream: false,
        requestSerialize: (value) => Buffer.from(cartMessage_1.FindCartItemsWithPaginationRequest.encode(value).finish()),
        requestDeserialize: (value) => cartMessage_1.FindCartItemsWithPaginationRequest.decode(value),
        responseSerialize: (value) => Buffer.from(cartMessage_1.FindCartItemsWithPaginationResponse.encode(value).finish()),
        responseDeserialize: (value) => cartMessage_1.FindCartItemsWithPaginationResponse.decode(value),
    },
    updateCartItemById: {
        path: "/commerce_v1.CartService/UpdateCartItemById",
        requestStream: false,
        responseStream: false,
        requestSerialize: (value) => Buffer.from(cartMessage_1.UpdateCartItemByIdRequest.encode(value).finish()),
        requestDeserialize: (value) => cartMessage_1.UpdateCartItemByIdRequest.decode(value),
        responseSerialize: (value) => Buffer.from(cartMessage_1.UpdateCartItemByIdResponse.encode(value).finish()),
        responseDeserialize: (value) => cartMessage_1.UpdateCartItemByIdResponse.decode(value),
    },
    deleteCartItemById: {
        path: "/commerce_v1.CartService/DeleteCartItemById",
        requestStream: false,
        responseStream: false,
        requestSerialize: (value) => Buffer.from(cartMessage_1.DeleteCartItemByIdRequest.encode(value).finish()),
        requestDeserialize: (value) => cartMessage_1.DeleteCartItemByIdRequest.decode(value),
        responseSerialize: (value) => Buffer.from(cartMessage_1.DeleteCartItemByIdResponse.encode(value).finish()),
        responseDeserialize: (value) => cartMessage_1.DeleteCartItemByIdResponse.decode(value),
    },
};
exports.CartServiceClient = (0, grpc_js_1.makeGenericClientConstructor)(exports.CartServiceService, "commerce_v1.CartService");
//# sourceMappingURL=cartServices.js.map