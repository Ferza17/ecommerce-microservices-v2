import { type CallOptions, ChannelCredentials, Client, type ClientOptions, type ClientUnaryCall, type handleUnaryCall, Metadata, type ServiceError, type UntypedServiceImplementation } from "@grpc/grpc-js";
import { CreateProductRequest, CreateProductResponse, DeleteProductByIdRequest, DeleteProductByIdResponse, FindProductByIdRequest, FindProductsWithPaginationRequest, FindProductsWithPaginationResponse, Product, UpdateProductByIdRequest } from "./productMessage";
export declare const protobufPackage = "product_v1";
export type ProductServiceService = typeof ProductServiceService;
export declare const ProductServiceService: {
    readonly findProductsWithPagination: {
        readonly path: "/product_v1.ProductService/FindProductsWithPagination";
        readonly requestStream: false;
        readonly responseStream: false;
        readonly requestSerialize: (value: FindProductsWithPaginationRequest) => Buffer<ArrayBuffer>;
        readonly requestDeserialize: (value: Buffer) => FindProductsWithPaginationRequest;
        readonly responseSerialize: (value: FindProductsWithPaginationResponse) => Buffer<ArrayBuffer>;
        readonly responseDeserialize: (value: Buffer) => FindProductsWithPaginationResponse;
    };
    readonly findProductById: {
        readonly path: "/product_v1.ProductService/FindProductById";
        readonly requestStream: false;
        readonly responseStream: false;
        readonly requestSerialize: (value: FindProductByIdRequest) => Buffer<ArrayBuffer>;
        readonly requestDeserialize: (value: Buffer) => FindProductByIdRequest;
        readonly responseSerialize: (value: Product) => Buffer<ArrayBuffer>;
        readonly responseDeserialize: (value: Buffer) => Product;
    };
    readonly createProduct: {
        readonly path: "/product_v1.ProductService/CreateProduct";
        readonly requestStream: false;
        readonly responseStream: false;
        readonly requestSerialize: (value: CreateProductRequest) => Buffer<ArrayBuffer>;
        readonly requestDeserialize: (value: Buffer) => CreateProductRequest;
        readonly responseSerialize: (value: CreateProductResponse) => Buffer<ArrayBuffer>;
        readonly responseDeserialize: (value: Buffer) => CreateProductResponse;
    };
    readonly updateProductById: {
        readonly path: "/product_v1.ProductService/UpdateProductById";
        readonly requestStream: false;
        readonly responseStream: false;
        readonly requestSerialize: (value: UpdateProductByIdRequest) => Buffer<ArrayBuffer>;
        readonly requestDeserialize: (value: Buffer) => UpdateProductByIdRequest;
        readonly responseSerialize: (value: Product) => Buffer<ArrayBuffer>;
        readonly responseDeserialize: (value: Buffer) => Product;
    };
    readonly deleteProductById: {
        readonly path: "/product_v1.ProductService/DeleteProductById";
        readonly requestStream: false;
        readonly responseStream: false;
        readonly requestSerialize: (value: DeleteProductByIdRequest) => Buffer<ArrayBuffer>;
        readonly requestDeserialize: (value: Buffer) => DeleteProductByIdRequest;
        readonly responseSerialize: (value: DeleteProductByIdResponse) => Buffer<ArrayBuffer>;
        readonly responseDeserialize: (value: Buffer) => DeleteProductByIdResponse;
    };
};
export interface ProductServiceServer extends UntypedServiceImplementation {
    findProductsWithPagination: handleUnaryCall<FindProductsWithPaginationRequest, FindProductsWithPaginationResponse>;
    findProductById: handleUnaryCall<FindProductByIdRequest, Product>;
    createProduct: handleUnaryCall<CreateProductRequest, CreateProductResponse>;
    updateProductById: handleUnaryCall<UpdateProductByIdRequest, Product>;
    deleteProductById: handleUnaryCall<DeleteProductByIdRequest, DeleteProductByIdResponse>;
}
export interface ProductServiceClient extends Client {
    findProductsWithPagination(request: FindProductsWithPaginationRequest, callback: (error: ServiceError | null, response: FindProductsWithPaginationResponse) => void): ClientUnaryCall;
    findProductsWithPagination(request: FindProductsWithPaginationRequest, metadata: Metadata, callback: (error: ServiceError | null, response: FindProductsWithPaginationResponse) => void): ClientUnaryCall;
    findProductsWithPagination(request: FindProductsWithPaginationRequest, metadata: Metadata, options: Partial<CallOptions>, callback: (error: ServiceError | null, response: FindProductsWithPaginationResponse) => void): ClientUnaryCall;
    findProductById(request: FindProductByIdRequest, callback: (error: ServiceError | null, response: Product) => void): ClientUnaryCall;
    findProductById(request: FindProductByIdRequest, metadata: Metadata, callback: (error: ServiceError | null, response: Product) => void): ClientUnaryCall;
    findProductById(request: FindProductByIdRequest, metadata: Metadata, options: Partial<CallOptions>, callback: (error: ServiceError | null, response: Product) => void): ClientUnaryCall;
    createProduct(request: CreateProductRequest, callback: (error: ServiceError | null, response: CreateProductResponse) => void): ClientUnaryCall;
    createProduct(request: CreateProductRequest, metadata: Metadata, callback: (error: ServiceError | null, response: CreateProductResponse) => void): ClientUnaryCall;
    createProduct(request: CreateProductRequest, metadata: Metadata, options: Partial<CallOptions>, callback: (error: ServiceError | null, response: CreateProductResponse) => void): ClientUnaryCall;
    updateProductById(request: UpdateProductByIdRequest, callback: (error: ServiceError | null, response: Product) => void): ClientUnaryCall;
    updateProductById(request: UpdateProductByIdRequest, metadata: Metadata, callback: (error: ServiceError | null, response: Product) => void): ClientUnaryCall;
    updateProductById(request: UpdateProductByIdRequest, metadata: Metadata, options: Partial<CallOptions>, callback: (error: ServiceError | null, response: Product) => void): ClientUnaryCall;
    deleteProductById(request: DeleteProductByIdRequest, callback: (error: ServiceError | null, response: DeleteProductByIdResponse) => void): ClientUnaryCall;
    deleteProductById(request: DeleteProductByIdRequest, metadata: Metadata, callback: (error: ServiceError | null, response: DeleteProductByIdResponse) => void): ClientUnaryCall;
    deleteProductById(request: DeleteProductByIdRequest, metadata: Metadata, options: Partial<CallOptions>, callback: (error: ServiceError | null, response: DeleteProductByIdResponse) => void): ClientUnaryCall;
}
export declare const ProductServiceClient: {
    new (address: string, credentials: ChannelCredentials, options?: Partial<ClientOptions>): ProductServiceClient;
    service: typeof ProductServiceService;
    serviceName: string;
};
