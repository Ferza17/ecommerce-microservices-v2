"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.ProductServiceClient = exports.ProductServiceService = exports.protobufPackage = void 0;
const grpc_js_1 = require("@grpc/grpc-js");
const productMessage_1 = require("./productMessage");
exports.protobufPackage = "product_v1";
exports.ProductServiceService = {
    findProductsWithPagination: {
        path: "/product_v1.ProductService/FindProductsWithPagination",
        requestStream: false,
        responseStream: false,
        requestSerialize: (value) => Buffer.from(productMessage_1.FindProductsWithPaginationRequest.encode(value).finish()),
        requestDeserialize: (value) => productMessage_1.FindProductsWithPaginationRequest.decode(value),
        responseSerialize: (value) => Buffer.from(productMessage_1.FindProductsWithPaginationResponse.encode(value).finish()),
        responseDeserialize: (value) => productMessage_1.FindProductsWithPaginationResponse.decode(value),
    },
    findProductById: {
        path: "/product_v1.ProductService/FindProductById",
        requestStream: false,
        responseStream: false,
        requestSerialize: (value) => Buffer.from(productMessage_1.FindProductByIdRequest.encode(value).finish()),
        requestDeserialize: (value) => productMessage_1.FindProductByIdRequest.decode(value),
        responseSerialize: (value) => Buffer.from(productMessage_1.Product.encode(value).finish()),
        responseDeserialize: (value) => productMessage_1.Product.decode(value),
    },
    createProduct: {
        path: "/product_v1.ProductService/CreateProduct",
        requestStream: false,
        responseStream: false,
        requestSerialize: (value) => Buffer.from(productMessage_1.CreateProductRequest.encode(value).finish()),
        requestDeserialize: (value) => productMessage_1.CreateProductRequest.decode(value),
        responseSerialize: (value) => Buffer.from(productMessage_1.CreateProductResponse.encode(value).finish()),
        responseDeserialize: (value) => productMessage_1.CreateProductResponse.decode(value),
    },
    updateProductById: {
        path: "/product_v1.ProductService/UpdateProductById",
        requestStream: false,
        responseStream: false,
        requestSerialize: (value) => Buffer.from(productMessage_1.UpdateProductByIdRequest.encode(value).finish()),
        requestDeserialize: (value) => productMessage_1.UpdateProductByIdRequest.decode(value),
        responseSerialize: (value) => Buffer.from(productMessage_1.Product.encode(value).finish()),
        responseDeserialize: (value) => productMessage_1.Product.decode(value),
    },
    deleteProductById: {
        path: "/product_v1.ProductService/DeleteProductById",
        requestStream: false,
        responseStream: false,
        requestSerialize: (value) => Buffer.from(productMessage_1.DeleteProductByIdRequest.encode(value).finish()),
        requestDeserialize: (value) => productMessage_1.DeleteProductByIdRequest.decode(value),
        responseSerialize: (value) => Buffer.from(productMessage_1.DeleteProductByIdResponse.encode(value).finish()),
        responseDeserialize: (value) => productMessage_1.DeleteProductByIdResponse.decode(value),
    },
};
exports.ProductServiceClient = (0, grpc_js_1.makeGenericClientConstructor)(exports.ProductServiceService, "product_v1.ProductService");
//# sourceMappingURL=productServices.js.map