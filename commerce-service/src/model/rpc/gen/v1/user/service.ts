// Code generated by protoc-gen-ts_proto. DO NOT EDIT.
// versions:
//   protoc-gen-ts_proto  v2.6.1
//   protoc               unknown
// source: v1/user/service.proto

/* eslint-disable */
import {
  type CallOptions,
  ChannelCredentials,
  Client,
  type ClientOptions,
  type ClientUnaryCall,
  type handleUnaryCall,
  makeGenericClientConstructor,
  Metadata,
  type ServiceError,
  type UntypedServiceImplementation,
} from "@grpc/grpc-js";
import { Empty } from "../../google/protobuf/empty";
import {
  AuthServiceVerifyIsExcludedRequest,
  AuthUserFindUserByTokenRequest,
  AuthUserLoginByEmailAndPasswordRequest,
  AuthUserLogoutByTokenRequest,
  AuthUserRegisterRequest,
  AuthUserVerifyAccessControlRequest,
  AuthUserVerifyOtpRequest,
  FindUserByEmailAndPasswordRequest,
  FindUserByIdRequest,
  UpdateUserByIdRequest,
} from "./request";
import {
  AuthServiceVerifyIsExcludedResponse,
  AuthUserFindUserByTokenResponse,
  AuthUserLogoutByTokenResponse,
  AuthUserVerifyAccessControlResponse,
  AuthUserVerifyOtpResponse,
  FindUserByEmailAndPasswordResponse,
  FindUserByIdResponse,
} from "./response";

export const protobufPackage = "user";

export type UserServiceService = typeof UserServiceService;
export const UserServiceService = {
  /** COMMAND */
  updateUserById: {
    path: "/user.UserService/UpdateUserById",
    requestStream: false,
    responseStream: false,
    requestSerialize: (value: UpdateUserByIdRequest) => Buffer.from(UpdateUserByIdRequest.encode(value).finish()),
    requestDeserialize: (value: Buffer) => UpdateUserByIdRequest.decode(value),
    responseSerialize: (value: Empty) => Buffer.from(Empty.encode(value).finish()),
    responseDeserialize: (value: Buffer) => Empty.decode(value),
  },
  /** QUERY */
  findUserById: {
    path: "/user.UserService/FindUserById",
    requestStream: false,
    responseStream: false,
    requestSerialize: (value: FindUserByIdRequest) => Buffer.from(FindUserByIdRequest.encode(value).finish()),
    requestDeserialize: (value: Buffer) => FindUserByIdRequest.decode(value),
    responseSerialize: (value: FindUserByIdResponse) => Buffer.from(FindUserByIdResponse.encode(value).finish()),
    responseDeserialize: (value: Buffer) => FindUserByIdResponse.decode(value),
  },
  findUserByEmailAndPassword: {
    path: "/user.UserService/FindUserByEmailAndPassword",
    requestStream: false,
    responseStream: false,
    requestSerialize: (value: FindUserByEmailAndPasswordRequest) =>
      Buffer.from(FindUserByEmailAndPasswordRequest.encode(value).finish()),
    requestDeserialize: (value: Buffer) => FindUserByEmailAndPasswordRequest.decode(value),
    responseSerialize: (value: FindUserByEmailAndPasswordResponse) =>
      Buffer.from(FindUserByEmailAndPasswordResponse.encode(value).finish()),
    responseDeserialize: (value: Buffer) => FindUserByEmailAndPasswordResponse.decode(value),
  },
} as const;

export interface UserServiceServer extends UntypedServiceImplementation {
  /** COMMAND */
  updateUserById: handleUnaryCall<UpdateUserByIdRequest, Empty>;
  /** QUERY */
  findUserById: handleUnaryCall<FindUserByIdRequest, FindUserByIdResponse>;
  findUserByEmailAndPassword: handleUnaryCall<FindUserByEmailAndPasswordRequest, FindUserByEmailAndPasswordResponse>;
}

export interface UserServiceClient extends Client {
  /** COMMAND */
  updateUserById(
    request: UpdateUserByIdRequest,
    callback: (error: ServiceError | null, response: Empty) => void,
  ): ClientUnaryCall;
  updateUserById(
    request: UpdateUserByIdRequest,
    metadata: Metadata,
    callback: (error: ServiceError | null, response: Empty) => void,
  ): ClientUnaryCall;
  updateUserById(
    request: UpdateUserByIdRequest,
    metadata: Metadata,
    options: Partial<CallOptions>,
    callback: (error: ServiceError | null, response: Empty) => void,
  ): ClientUnaryCall;
  /** QUERY */
  findUserById(
    request: FindUserByIdRequest,
    callback: (error: ServiceError | null, response: FindUserByIdResponse) => void,
  ): ClientUnaryCall;
  findUserById(
    request: FindUserByIdRequest,
    metadata: Metadata,
    callback: (error: ServiceError | null, response: FindUserByIdResponse) => void,
  ): ClientUnaryCall;
  findUserById(
    request: FindUserByIdRequest,
    metadata: Metadata,
    options: Partial<CallOptions>,
    callback: (error: ServiceError | null, response: FindUserByIdResponse) => void,
  ): ClientUnaryCall;
  findUserByEmailAndPassword(
    request: FindUserByEmailAndPasswordRequest,
    callback: (error: ServiceError | null, response: FindUserByEmailAndPasswordResponse) => void,
  ): ClientUnaryCall;
  findUserByEmailAndPassword(
    request: FindUserByEmailAndPasswordRequest,
    metadata: Metadata,
    callback: (error: ServiceError | null, response: FindUserByEmailAndPasswordResponse) => void,
  ): ClientUnaryCall;
  findUserByEmailAndPassword(
    request: FindUserByEmailAndPasswordRequest,
    metadata: Metadata,
    options: Partial<CallOptions>,
    callback: (error: ServiceError | null, response: FindUserByEmailAndPasswordResponse) => void,
  ): ClientUnaryCall;
}

export const UserServiceClient = makeGenericClientConstructor(UserServiceService, "user.UserService") as unknown as {
  new (address: string, credentials: ChannelCredentials, options?: Partial<ClientOptions>): UserServiceClient;
  service: typeof UserServiceService;
  serviceName: string;
};

export type AuthServiceService = typeof AuthServiceService;
export const AuthServiceService = {
  /** COMMAND */
  authUserRegister: {
    path: "/user.AuthService/AuthUserRegister",
    requestStream: false,
    responseStream: false,
    requestSerialize: (value: AuthUserRegisterRequest) => Buffer.from(AuthUserRegisterRequest.encode(value).finish()),
    requestDeserialize: (value: Buffer) => AuthUserRegisterRequest.decode(value),
    responseSerialize: (value: Empty) => Buffer.from(Empty.encode(value).finish()),
    responseDeserialize: (value: Buffer) => Empty.decode(value),
  },
  authUserLoginByEmailAndPassword: {
    path: "/user.AuthService/AuthUserLoginByEmailAndPassword",
    requestStream: false,
    responseStream: false,
    requestSerialize: (value: AuthUserLoginByEmailAndPasswordRequest) =>
      Buffer.from(AuthUserLoginByEmailAndPasswordRequest.encode(value).finish()),
    requestDeserialize: (value: Buffer) => AuthUserLoginByEmailAndPasswordRequest.decode(value),
    responseSerialize: (value: Empty) => Buffer.from(Empty.encode(value).finish()),
    responseDeserialize: (value: Buffer) => Empty.decode(value),
  },
  authUserVerifyOtp: {
    path: "/user.AuthService/AuthUserVerifyOtp",
    requestStream: false,
    responseStream: false,
    requestSerialize: (value: AuthUserVerifyOtpRequest) => Buffer.from(AuthUserVerifyOtpRequest.encode(value).finish()),
    requestDeserialize: (value: Buffer) => AuthUserVerifyOtpRequest.decode(value),
    responseSerialize: (value: AuthUserVerifyOtpResponse) =>
      Buffer.from(AuthUserVerifyOtpResponse.encode(value).finish()),
    responseDeserialize: (value: Buffer) => AuthUserVerifyOtpResponse.decode(value),
  },
  authUserLogoutByToken: {
    path: "/user.AuthService/AuthUserLogoutByToken",
    requestStream: false,
    responseStream: false,
    requestSerialize: (value: AuthUserLogoutByTokenRequest) =>
      Buffer.from(AuthUserLogoutByTokenRequest.encode(value).finish()),
    requestDeserialize: (value: Buffer) => AuthUserLogoutByTokenRequest.decode(value),
    responseSerialize: (value: AuthUserLogoutByTokenResponse) =>
      Buffer.from(AuthUserLogoutByTokenResponse.encode(value).finish()),
    responseDeserialize: (value: Buffer) => AuthUserLogoutByTokenResponse.decode(value),
  },
  authUserVerifyAccessControl: {
    path: "/user.AuthService/AuthUserVerifyAccessControl",
    requestStream: false,
    responseStream: false,
    requestSerialize: (value: AuthUserVerifyAccessControlRequest) =>
      Buffer.from(AuthUserVerifyAccessControlRequest.encode(value).finish()),
    requestDeserialize: (value: Buffer) => AuthUserVerifyAccessControlRequest.decode(value),
    responseSerialize: (value: AuthUserVerifyAccessControlResponse) =>
      Buffer.from(AuthUserVerifyAccessControlResponse.encode(value).finish()),
    responseDeserialize: (value: Buffer) => AuthUserVerifyAccessControlResponse.decode(value),
  },
  authServiceVerifyIsExcluded: {
    path: "/user.AuthService/AuthServiceVerifyIsExcluded",
    requestStream: false,
    responseStream: false,
    requestSerialize: (value: AuthServiceVerifyIsExcludedRequest) =>
      Buffer.from(AuthServiceVerifyIsExcludedRequest.encode(value).finish()),
    requestDeserialize: (value: Buffer) => AuthServiceVerifyIsExcludedRequest.decode(value),
    responseSerialize: (value: AuthServiceVerifyIsExcludedResponse) =>
      Buffer.from(AuthServiceVerifyIsExcludedResponse.encode(value).finish()),
    responseDeserialize: (value: Buffer) => AuthServiceVerifyIsExcludedResponse.decode(value),
  },
  /** QUERY */
  authUserFindUserByToken: {
    path: "/user.AuthService/AuthUserFindUserByToken",
    requestStream: false,
    responseStream: false,
    requestSerialize: (value: AuthUserFindUserByTokenRequest) =>
      Buffer.from(AuthUserFindUserByTokenRequest.encode(value).finish()),
    requestDeserialize: (value: Buffer) => AuthUserFindUserByTokenRequest.decode(value),
    responseSerialize: (value: AuthUserFindUserByTokenResponse) =>
      Buffer.from(AuthUserFindUserByTokenResponse.encode(value).finish()),
    responseDeserialize: (value: Buffer) => AuthUserFindUserByTokenResponse.decode(value),
  },
} as const;

export interface AuthServiceServer extends UntypedServiceImplementation {
  /** COMMAND */
  authUserRegister: handleUnaryCall<AuthUserRegisterRequest, Empty>;
  authUserLoginByEmailAndPassword: handleUnaryCall<AuthUserLoginByEmailAndPasswordRequest, Empty>;
  authUserVerifyOtp: handleUnaryCall<AuthUserVerifyOtpRequest, AuthUserVerifyOtpResponse>;
  authUserLogoutByToken: handleUnaryCall<AuthUserLogoutByTokenRequest, AuthUserLogoutByTokenResponse>;
  authUserVerifyAccessControl: handleUnaryCall<AuthUserVerifyAccessControlRequest, AuthUserVerifyAccessControlResponse>;
  authServiceVerifyIsExcluded: handleUnaryCall<AuthServiceVerifyIsExcludedRequest, AuthServiceVerifyIsExcludedResponse>;
  /** QUERY */
  authUserFindUserByToken: handleUnaryCall<AuthUserFindUserByTokenRequest, AuthUserFindUserByTokenResponse>;
}

export interface AuthServiceClient extends Client {
  /** COMMAND */
  authUserRegister(
    request: AuthUserRegisterRequest,
    callback: (error: ServiceError | null, response: Empty) => void,
  ): ClientUnaryCall;
  authUserRegister(
    request: AuthUserRegisterRequest,
    metadata: Metadata,
    callback: (error: ServiceError | null, response: Empty) => void,
  ): ClientUnaryCall;
  authUserRegister(
    request: AuthUserRegisterRequest,
    metadata: Metadata,
    options: Partial<CallOptions>,
    callback: (error: ServiceError | null, response: Empty) => void,
  ): ClientUnaryCall;
  authUserLoginByEmailAndPassword(
    request: AuthUserLoginByEmailAndPasswordRequest,
    callback: (error: ServiceError | null, response: Empty) => void,
  ): ClientUnaryCall;
  authUserLoginByEmailAndPassword(
    request: AuthUserLoginByEmailAndPasswordRequest,
    metadata: Metadata,
    callback: (error: ServiceError | null, response: Empty) => void,
  ): ClientUnaryCall;
  authUserLoginByEmailAndPassword(
    request: AuthUserLoginByEmailAndPasswordRequest,
    metadata: Metadata,
    options: Partial<CallOptions>,
    callback: (error: ServiceError | null, response: Empty) => void,
  ): ClientUnaryCall;
  authUserVerifyOtp(
    request: AuthUserVerifyOtpRequest,
    callback: (error: ServiceError | null, response: AuthUserVerifyOtpResponse) => void,
  ): ClientUnaryCall;
  authUserVerifyOtp(
    request: AuthUserVerifyOtpRequest,
    metadata: Metadata,
    callback: (error: ServiceError | null, response: AuthUserVerifyOtpResponse) => void,
  ): ClientUnaryCall;
  authUserVerifyOtp(
    request: AuthUserVerifyOtpRequest,
    metadata: Metadata,
    options: Partial<CallOptions>,
    callback: (error: ServiceError | null, response: AuthUserVerifyOtpResponse) => void,
  ): ClientUnaryCall;
  authUserLogoutByToken(
    request: AuthUserLogoutByTokenRequest,
    callback: (error: ServiceError | null, response: AuthUserLogoutByTokenResponse) => void,
  ): ClientUnaryCall;
  authUserLogoutByToken(
    request: AuthUserLogoutByTokenRequest,
    metadata: Metadata,
    callback: (error: ServiceError | null, response: AuthUserLogoutByTokenResponse) => void,
  ): ClientUnaryCall;
  authUserLogoutByToken(
    request: AuthUserLogoutByTokenRequest,
    metadata: Metadata,
    options: Partial<CallOptions>,
    callback: (error: ServiceError | null, response: AuthUserLogoutByTokenResponse) => void,
  ): ClientUnaryCall;
  authUserVerifyAccessControl(
    request: AuthUserVerifyAccessControlRequest,
    callback: (error: ServiceError | null, response: AuthUserVerifyAccessControlResponse) => void,
  ): ClientUnaryCall;
  authUserVerifyAccessControl(
    request: AuthUserVerifyAccessControlRequest,
    metadata: Metadata,
    callback: (error: ServiceError | null, response: AuthUserVerifyAccessControlResponse) => void,
  ): ClientUnaryCall;
  authUserVerifyAccessControl(
    request: AuthUserVerifyAccessControlRequest,
    metadata: Metadata,
    options: Partial<CallOptions>,
    callback: (error: ServiceError | null, response: AuthUserVerifyAccessControlResponse) => void,
  ): ClientUnaryCall;
  authServiceVerifyIsExcluded(
    request: AuthServiceVerifyIsExcludedRequest,
    callback: (error: ServiceError | null, response: AuthServiceVerifyIsExcludedResponse) => void,
  ): ClientUnaryCall;
  authServiceVerifyIsExcluded(
    request: AuthServiceVerifyIsExcludedRequest,
    metadata: Metadata,
    callback: (error: ServiceError | null, response: AuthServiceVerifyIsExcludedResponse) => void,
  ): ClientUnaryCall;
  authServiceVerifyIsExcluded(
    request: AuthServiceVerifyIsExcludedRequest,
    metadata: Metadata,
    options: Partial<CallOptions>,
    callback: (error: ServiceError | null, response: AuthServiceVerifyIsExcludedResponse) => void,
  ): ClientUnaryCall;
  /** QUERY */
  authUserFindUserByToken(
    request: AuthUserFindUserByTokenRequest,
    callback: (error: ServiceError | null, response: AuthUserFindUserByTokenResponse) => void,
  ): ClientUnaryCall;
  authUserFindUserByToken(
    request: AuthUserFindUserByTokenRequest,
    metadata: Metadata,
    callback: (error: ServiceError | null, response: AuthUserFindUserByTokenResponse) => void,
  ): ClientUnaryCall;
  authUserFindUserByToken(
    request: AuthUserFindUserByTokenRequest,
    metadata: Metadata,
    options: Partial<CallOptions>,
    callback: (error: ServiceError | null, response: AuthUserFindUserByTokenResponse) => void,
  ): ClientUnaryCall;
}

export const AuthServiceClient = makeGenericClientConstructor(AuthServiceService, "user.AuthService") as unknown as {
  new (address: string, credentials: ChannelCredentials, options?: Partial<ClientOptions>): AuthServiceClient;
  service: typeof AuthServiceService;
  serviceName: string;
};
