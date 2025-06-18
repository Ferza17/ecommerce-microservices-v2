import { type CallOptions, ChannelCredentials, Client, type ClientOptions, type ClientUnaryCall, type handleUnaryCall, Metadata, type ServiceError, type UntypedServiceImplementation } from "@grpc/grpc-js";
import { CartItem, CreateCartItemRequest, CreateCartItemResponse, DeleteCartItemByIdRequest, DeleteCartItemByIdResponse, FindCartItemByIdRequest, FindCartItemsWithPaginationRequest, FindCartItemsWithPaginationResponse, UpdateCartItemByIdRequest, UpdateCartItemByIdResponse } from "./cartMessage";
export declare const protobufPackage = "commerce_v1";
export type CartServiceService = typeof CartServiceService;
export declare const CartServiceService: {
    readonly createCartItem: {
        readonly path: "/commerce_v1.CartService/CreateCartItem";
        readonly requestStream: false;
        readonly responseStream: false;
        readonly requestSerialize: (value: CreateCartItemRequest) => Buffer<ArrayBuffer>;
        readonly requestDeserialize: (value: Buffer) => CreateCartItemRequest;
        readonly responseSerialize: (value: CreateCartItemResponse) => Buffer<ArrayBuffer>;
        readonly responseDeserialize: (value: Buffer) => CreateCartItemResponse;
    };
    readonly findCartItemById: {
        readonly path: "/commerce_v1.CartService/FindCartItemById";
        readonly requestStream: false;
        readonly responseStream: false;
        readonly requestSerialize: (value: FindCartItemByIdRequest) => Buffer<ArrayBuffer>;
        readonly requestDeserialize: (value: Buffer) => FindCartItemByIdRequest;
        readonly responseSerialize: (value: CartItem) => Buffer<ArrayBuffer>;
        readonly responseDeserialize: (value: Buffer) => CartItem;
    };
    readonly findCartItemsWithPagination: {
        readonly path: "/commerce_v1.CartService/FindCartItemsWithPagination";
        readonly requestStream: false;
        readonly responseStream: false;
        readonly requestSerialize: (value: FindCartItemsWithPaginationRequest) => Buffer<ArrayBuffer>;
        readonly requestDeserialize: (value: Buffer) => FindCartItemsWithPaginationRequest;
        readonly responseSerialize: (value: FindCartItemsWithPaginationResponse) => Buffer<ArrayBuffer>;
        readonly responseDeserialize: (value: Buffer) => FindCartItemsWithPaginationResponse;
    };
    readonly updateCartItemById: {
        readonly path: "/commerce_v1.CartService/UpdateCartItemById";
        readonly requestStream: false;
        readonly responseStream: false;
        readonly requestSerialize: (value: UpdateCartItemByIdRequest) => Buffer<ArrayBuffer>;
        readonly requestDeserialize: (value: Buffer) => UpdateCartItemByIdRequest;
        readonly responseSerialize: (value: UpdateCartItemByIdResponse) => Buffer<ArrayBuffer>;
        readonly responseDeserialize: (value: Buffer) => UpdateCartItemByIdResponse;
    };
    readonly deleteCartItemById: {
        readonly path: "/commerce_v1.CartService/DeleteCartItemById";
        readonly requestStream: false;
        readonly responseStream: false;
        readonly requestSerialize: (value: DeleteCartItemByIdRequest) => Buffer<ArrayBuffer>;
        readonly requestDeserialize: (value: Buffer) => DeleteCartItemByIdRequest;
        readonly responseSerialize: (value: DeleteCartItemByIdResponse) => Buffer<ArrayBuffer>;
        readonly responseDeserialize: (value: Buffer) => DeleteCartItemByIdResponse;
    };
};
export interface CartServiceServer extends UntypedServiceImplementation {
    createCartItem: handleUnaryCall<CreateCartItemRequest, CreateCartItemResponse>;
    findCartItemById: handleUnaryCall<FindCartItemByIdRequest, CartItem>;
    findCartItemsWithPagination: handleUnaryCall<FindCartItemsWithPaginationRequest, FindCartItemsWithPaginationResponse>;
    updateCartItemById: handleUnaryCall<UpdateCartItemByIdRequest, UpdateCartItemByIdResponse>;
    deleteCartItemById: handleUnaryCall<DeleteCartItemByIdRequest, DeleteCartItemByIdResponse>;
}
export interface CartServiceClient extends Client {
    createCartItem(request: CreateCartItemRequest, callback: (error: ServiceError | null, response: CreateCartItemResponse) => void): ClientUnaryCall;
    createCartItem(request: CreateCartItemRequest, metadata: Metadata, callback: (error: ServiceError | null, response: CreateCartItemResponse) => void): ClientUnaryCall;
    createCartItem(request: CreateCartItemRequest, metadata: Metadata, options: Partial<CallOptions>, callback: (error: ServiceError | null, response: CreateCartItemResponse) => void): ClientUnaryCall;
    findCartItemById(request: FindCartItemByIdRequest, callback: (error: ServiceError | null, response: CartItem) => void): ClientUnaryCall;
    findCartItemById(request: FindCartItemByIdRequest, metadata: Metadata, callback: (error: ServiceError | null, response: CartItem) => void): ClientUnaryCall;
    findCartItemById(request: FindCartItemByIdRequest, metadata: Metadata, options: Partial<CallOptions>, callback: (error: ServiceError | null, response: CartItem) => void): ClientUnaryCall;
    findCartItemsWithPagination(request: FindCartItemsWithPaginationRequest, callback: (error: ServiceError | null, response: FindCartItemsWithPaginationResponse) => void): ClientUnaryCall;
    findCartItemsWithPagination(request: FindCartItemsWithPaginationRequest, metadata: Metadata, callback: (error: ServiceError | null, response: FindCartItemsWithPaginationResponse) => void): ClientUnaryCall;
    findCartItemsWithPagination(request: FindCartItemsWithPaginationRequest, metadata: Metadata, options: Partial<CallOptions>, callback: (error: ServiceError | null, response: FindCartItemsWithPaginationResponse) => void): ClientUnaryCall;
    updateCartItemById(request: UpdateCartItemByIdRequest, callback: (error: ServiceError | null, response: UpdateCartItemByIdResponse) => void): ClientUnaryCall;
    updateCartItemById(request: UpdateCartItemByIdRequest, metadata: Metadata, callback: (error: ServiceError | null, response: UpdateCartItemByIdResponse) => void): ClientUnaryCall;
    updateCartItemById(request: UpdateCartItemByIdRequest, metadata: Metadata, options: Partial<CallOptions>, callback: (error: ServiceError | null, response: UpdateCartItemByIdResponse) => void): ClientUnaryCall;
    deleteCartItemById(request: DeleteCartItemByIdRequest, callback: (error: ServiceError | null, response: DeleteCartItemByIdResponse) => void): ClientUnaryCall;
    deleteCartItemById(request: DeleteCartItemByIdRequest, metadata: Metadata, callback: (error: ServiceError | null, response: DeleteCartItemByIdResponse) => void): ClientUnaryCall;
    deleteCartItemById(request: DeleteCartItemByIdRequest, metadata: Metadata, options: Partial<CallOptions>, callback: (error: ServiceError | null, response: DeleteCartItemByIdResponse) => void): ClientUnaryCall;
}
export declare const CartServiceClient: {
    new (address: string, credentials: ChannelCredentials, options?: Partial<ClientOptions>): CartServiceClient;
    service: typeof CartServiceService;
    serviceName: string;
};
