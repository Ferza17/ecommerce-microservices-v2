import { type CallOptions, ChannelCredentials, Client, type ClientOptions, type ClientUnaryCall, type handleUnaryCall, Metadata, type ServiceError, type UntypedServiceImplementation } from "@grpc/grpc-js";
import { FindUserByTokenRequest, UserLogoutByTokenRequest, UserLogoutByTokenResponse, UserVerifyOtpRequest, UserVerifyOtpResponse } from "./authMessage";
import { User } from "./userMessage";
export declare const protobufPackage = "user_v1";
export type AuthServiceService = typeof AuthServiceService;
export declare const AuthServiceService: {
    readonly userLogoutByToken: {
        readonly path: "/user_v1.AuthService/UserLogoutByToken";
        readonly requestStream: false;
        readonly responseStream: false;
        readonly requestSerialize: (value: UserLogoutByTokenRequest) => Buffer<ArrayBuffer>;
        readonly requestDeserialize: (value: Buffer) => UserLogoutByTokenRequest;
        readonly responseSerialize: (value: UserLogoutByTokenResponse) => Buffer<ArrayBuffer>;
        readonly responseDeserialize: (value: Buffer) => UserLogoutByTokenResponse;
    };
    readonly userVerifyOtp: {
        readonly path: "/user_v1.AuthService/UserVerifyOtp";
        readonly requestStream: false;
        readonly responseStream: false;
        readonly requestSerialize: (value: UserVerifyOtpRequest) => Buffer<ArrayBuffer>;
        readonly requestDeserialize: (value: Buffer) => UserVerifyOtpRequest;
        readonly responseSerialize: (value: UserVerifyOtpResponse) => Buffer<ArrayBuffer>;
        readonly responseDeserialize: (value: Buffer) => UserVerifyOtpResponse;
    };
    readonly findUserByToken: {
        readonly path: "/user_v1.AuthService/FindUserByToken";
        readonly requestStream: false;
        readonly responseStream: false;
        readonly requestSerialize: (value: FindUserByTokenRequest) => Buffer<ArrayBuffer>;
        readonly requestDeserialize: (value: Buffer) => FindUserByTokenRequest;
        readonly responseSerialize: (value: User) => Buffer<ArrayBuffer>;
        readonly responseDeserialize: (value: Buffer) => User;
    };
};
export interface AuthServiceServer extends UntypedServiceImplementation {
    userLogoutByToken: handleUnaryCall<UserLogoutByTokenRequest, UserLogoutByTokenResponse>;
    userVerifyOtp: handleUnaryCall<UserVerifyOtpRequest, UserVerifyOtpResponse>;
    findUserByToken: handleUnaryCall<FindUserByTokenRequest, User>;
}
export interface AuthServiceClient extends Client {
    userLogoutByToken(request: UserLogoutByTokenRequest, callback: (error: ServiceError | null, response: UserLogoutByTokenResponse) => void): ClientUnaryCall;
    userLogoutByToken(request: UserLogoutByTokenRequest, metadata: Metadata, callback: (error: ServiceError | null, response: UserLogoutByTokenResponse) => void): ClientUnaryCall;
    userLogoutByToken(request: UserLogoutByTokenRequest, metadata: Metadata, options: Partial<CallOptions>, callback: (error: ServiceError | null, response: UserLogoutByTokenResponse) => void): ClientUnaryCall;
    userVerifyOtp(request: UserVerifyOtpRequest, callback: (error: ServiceError | null, response: UserVerifyOtpResponse) => void): ClientUnaryCall;
    userVerifyOtp(request: UserVerifyOtpRequest, metadata: Metadata, callback: (error: ServiceError | null, response: UserVerifyOtpResponse) => void): ClientUnaryCall;
    userVerifyOtp(request: UserVerifyOtpRequest, metadata: Metadata, options: Partial<CallOptions>, callback: (error: ServiceError | null, response: UserVerifyOtpResponse) => void): ClientUnaryCall;
    findUserByToken(request: FindUserByTokenRequest, callback: (error: ServiceError | null, response: User) => void): ClientUnaryCall;
    findUserByToken(request: FindUserByTokenRequest, metadata: Metadata, callback: (error: ServiceError | null, response: User) => void): ClientUnaryCall;
    findUserByToken(request: FindUserByTokenRequest, metadata: Metadata, options: Partial<CallOptions>, callback: (error: ServiceError | null, response: User) => void): ClientUnaryCall;
}
export declare const AuthServiceClient: {
    new (address: string, credentials: ChannelCredentials, options?: Partial<ClientOptions>): AuthServiceClient;
    service: typeof AuthServiceService;
    serviceName: string;
};
