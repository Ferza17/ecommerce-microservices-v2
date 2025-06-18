import { type CallOptions, ChannelCredentials, Client, type ClientOptions, type ClientUnaryCall, type handleUnaryCall, Metadata, type ServiceError, type UntypedServiceImplementation } from "@grpc/grpc-js";
import { CreateWishlistItemRequest, CreateWishlistItemResponse, DeleteWishlistItemByIdRequest, DeleteWishlistItemByIdResponse, FindWishlistItemWithPaginationRequest, FindWishlistItemWithPaginationResponse } from "./wishlistMessage";
export declare const protobufPackage = "commerce_v1";
export type WishlistServiceService = typeof WishlistServiceService;
export declare const WishlistServiceService: {
    readonly createWishlistItem: {
        readonly path: "/commerce_v1.WishlistService/CreateWishlistItem";
        readonly requestStream: false;
        readonly responseStream: false;
        readonly requestSerialize: (value: CreateWishlistItemRequest) => Buffer<ArrayBuffer>;
        readonly requestDeserialize: (value: Buffer) => CreateWishlistItemRequest;
        readonly responseSerialize: (value: CreateWishlistItemResponse) => Buffer<ArrayBuffer>;
        readonly responseDeserialize: (value: Buffer) => CreateWishlistItemResponse;
    };
    readonly findWishlistItemWithPagination: {
        readonly path: "/commerce_v1.WishlistService/FindWishlistItemWithPagination";
        readonly requestStream: false;
        readonly responseStream: false;
        readonly requestSerialize: (value: FindWishlistItemWithPaginationRequest) => Buffer<ArrayBuffer>;
        readonly requestDeserialize: (value: Buffer) => FindWishlistItemWithPaginationRequest;
        readonly responseSerialize: (value: FindWishlistItemWithPaginationResponse) => Buffer<ArrayBuffer>;
        readonly responseDeserialize: (value: Buffer) => FindWishlistItemWithPaginationResponse;
    };
    readonly deleteWishlistItemById: {
        readonly path: "/commerce_v1.WishlistService/DeleteWishlistItemById";
        readonly requestStream: false;
        readonly responseStream: false;
        readonly requestSerialize: (value: DeleteWishlistItemByIdRequest) => Buffer<ArrayBuffer>;
        readonly requestDeserialize: (value: Buffer) => DeleteWishlistItemByIdRequest;
        readonly responseSerialize: (value: DeleteWishlistItemByIdResponse) => Buffer<ArrayBuffer>;
        readonly responseDeserialize: (value: Buffer) => DeleteWishlistItemByIdResponse;
    };
};
export interface WishlistServiceServer extends UntypedServiceImplementation {
    createWishlistItem: handleUnaryCall<CreateWishlistItemRequest, CreateWishlistItemResponse>;
    findWishlistItemWithPagination: handleUnaryCall<FindWishlistItemWithPaginationRequest, FindWishlistItemWithPaginationResponse>;
    deleteWishlistItemById: handleUnaryCall<DeleteWishlistItemByIdRequest, DeleteWishlistItemByIdResponse>;
}
export interface WishlistServiceClient extends Client {
    createWishlistItem(request: CreateWishlistItemRequest, callback: (error: ServiceError | null, response: CreateWishlistItemResponse) => void): ClientUnaryCall;
    createWishlistItem(request: CreateWishlistItemRequest, metadata: Metadata, callback: (error: ServiceError | null, response: CreateWishlistItemResponse) => void): ClientUnaryCall;
    createWishlistItem(request: CreateWishlistItemRequest, metadata: Metadata, options: Partial<CallOptions>, callback: (error: ServiceError | null, response: CreateWishlistItemResponse) => void): ClientUnaryCall;
    findWishlistItemWithPagination(request: FindWishlistItemWithPaginationRequest, callback: (error: ServiceError | null, response: FindWishlistItemWithPaginationResponse) => void): ClientUnaryCall;
    findWishlistItemWithPagination(request: FindWishlistItemWithPaginationRequest, metadata: Metadata, callback: (error: ServiceError | null, response: FindWishlistItemWithPaginationResponse) => void): ClientUnaryCall;
    findWishlistItemWithPagination(request: FindWishlistItemWithPaginationRequest, metadata: Metadata, options: Partial<CallOptions>, callback: (error: ServiceError | null, response: FindWishlistItemWithPaginationResponse) => void): ClientUnaryCall;
    deleteWishlistItemById(request: DeleteWishlistItemByIdRequest, callback: (error: ServiceError | null, response: DeleteWishlistItemByIdResponse) => void): ClientUnaryCall;
    deleteWishlistItemById(request: DeleteWishlistItemByIdRequest, metadata: Metadata, callback: (error: ServiceError | null, response: DeleteWishlistItemByIdResponse) => void): ClientUnaryCall;
    deleteWishlistItemById(request: DeleteWishlistItemByIdRequest, metadata: Metadata, options: Partial<CallOptions>, callback: (error: ServiceError | null, response: DeleteWishlistItemByIdResponse) => void): ClientUnaryCall;
}
export declare const WishlistServiceClient: {
    new (address: string, credentials: ChannelCredentials, options?: Partial<ClientOptions>): WishlistServiceClient;
    service: typeof WishlistServiceService;
    serviceName: string;
};
