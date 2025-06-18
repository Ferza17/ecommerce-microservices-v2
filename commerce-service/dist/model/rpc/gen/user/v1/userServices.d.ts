import { type CallOptions, ChannelCredentials, Client, type ClientOptions, type ClientUnaryCall, type handleUnaryCall, Metadata, type ServiceError, type UntypedServiceImplementation } from "@grpc/grpc-js";
import { FindUserByEmailAndPasswordRequest, FindUserByIdRequest, User } from "./userMessage";
export declare const protobufPackage = "user_v1";
export type UserServiceService = typeof UserServiceService;
export declare const UserServiceService: {
    readonly findUserById: {
        readonly path: "/user_v1.UserService/FindUserById";
        readonly requestStream: false;
        readonly responseStream: false;
        readonly requestSerialize: (value: FindUserByIdRequest) => Buffer<ArrayBuffer>;
        readonly requestDeserialize: (value: Buffer) => FindUserByIdRequest;
        readonly responseSerialize: (value: User) => Buffer<ArrayBuffer>;
        readonly responseDeserialize: (value: Buffer) => User;
    };
    readonly findUserByEmailAndPassword: {
        readonly path: "/user_v1.UserService/FindUserByEmailAndPassword";
        readonly requestStream: false;
        readonly responseStream: false;
        readonly requestSerialize: (value: FindUserByEmailAndPasswordRequest) => Buffer<ArrayBuffer>;
        readonly requestDeserialize: (value: Buffer) => FindUserByEmailAndPasswordRequest;
        readonly responseSerialize: (value: User) => Buffer<ArrayBuffer>;
        readonly responseDeserialize: (value: Buffer) => User;
    };
};
export interface UserServiceServer extends UntypedServiceImplementation {
    findUserById: handleUnaryCall<FindUserByIdRequest, User>;
    findUserByEmailAndPassword: handleUnaryCall<FindUserByEmailAndPasswordRequest, User>;
}
export interface UserServiceClient extends Client {
    findUserById(request: FindUserByIdRequest, callback: (error: ServiceError | null, response: User) => void): ClientUnaryCall;
    findUserById(request: FindUserByIdRequest, metadata: Metadata, callback: (error: ServiceError | null, response: User) => void): ClientUnaryCall;
    findUserById(request: FindUserByIdRequest, metadata: Metadata, options: Partial<CallOptions>, callback: (error: ServiceError | null, response: User) => void): ClientUnaryCall;
    findUserByEmailAndPassword(request: FindUserByEmailAndPasswordRequest, callback: (error: ServiceError | null, response: User) => void): ClientUnaryCall;
    findUserByEmailAndPassword(request: FindUserByEmailAndPasswordRequest, metadata: Metadata, callback: (error: ServiceError | null, response: User) => void): ClientUnaryCall;
    findUserByEmailAndPassword(request: FindUserByEmailAndPasswordRequest, metadata: Metadata, options: Partial<CallOptions>, callback: (error: ServiceError | null, response: User) => void): ClientUnaryCall;
}
export declare const UserServiceClient: {
    new (address: string, credentials: ChannelCredentials, options?: Partial<ClientOptions>): UserServiceClient;
    service: typeof UserServiceService;
    serviceName: string;
};
