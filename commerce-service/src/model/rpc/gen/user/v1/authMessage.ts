// Code generated by protoc-gen-ts_proto. DO NOT EDIT.
// versions:
//   protoc-gen-ts_proto  v2.6.1
//   protoc               unknown
// source: user/v1/authMessage.proto

/* eslint-disable */
import { BinaryReader, BinaryWriter } from "@bufbuild/protobuf/wire";

export const protobufPackage = "user_v1";

export interface UserLoginByEmailAndPasswordRequest {
  email: string;
  password: string;
}

export interface UserLogoutByTokenRequest {
  token: string;
}

export interface UserLogoutByTokenResponse {
}

export interface FindUserByTokenRequest {
  token: string;
}

export interface UserVerifyOtpRequest {
  otp: string;
}

export interface UserVerifyOtpResponse {
  accessToken: string;
  refreshToken: string;
}

function createBaseUserLoginByEmailAndPasswordRequest(): UserLoginByEmailAndPasswordRequest {
  return { email: "", password: "" };
}

export const UserLoginByEmailAndPasswordRequest: MessageFns<UserLoginByEmailAndPasswordRequest> = {
  encode(message: UserLoginByEmailAndPasswordRequest, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
    if (message.email !== "") {
      writer.uint32(10).string(message.email);
    }
    if (message.password !== "") {
      writer.uint32(18).string(message.password);
    }
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): UserLoginByEmailAndPasswordRequest {
    const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUserLoginByEmailAndPasswordRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1: {
          if (tag !== 10) {
            break;
          }

          message.email = reader.string();
          continue;
        }
        case 2: {
          if (tag !== 18) {
            break;
          }

          message.password = reader.string();
          continue;
        }
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skip(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): UserLoginByEmailAndPasswordRequest {
    return {
      email: isSet(object.email) ? globalThis.String(object.email) : "",
      password: isSet(object.password) ? globalThis.String(object.password) : "",
    };
  },

  toJSON(message: UserLoginByEmailAndPasswordRequest): unknown {
    const obj: any = {};
    if (message.email !== "") {
      obj.email = message.email;
    }
    if (message.password !== "") {
      obj.password = message.password;
    }
    return obj;
  },

  create(base?: DeepPartial<UserLoginByEmailAndPasswordRequest>): UserLoginByEmailAndPasswordRequest {
    return UserLoginByEmailAndPasswordRequest.fromPartial(base ?? {});
  },
  fromPartial(object: DeepPartial<UserLoginByEmailAndPasswordRequest>): UserLoginByEmailAndPasswordRequest {
    const message = createBaseUserLoginByEmailAndPasswordRequest();
    message.email = object.email ?? "";
    message.password = object.password ?? "";
    return message;
  },
};

function createBaseUserLogoutByTokenRequest(): UserLogoutByTokenRequest {
  return { token: "" };
}

export const UserLogoutByTokenRequest: MessageFns<UserLogoutByTokenRequest> = {
  encode(message: UserLogoutByTokenRequest, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
    if (message.token !== "") {
      writer.uint32(10).string(message.token);
    }
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): UserLogoutByTokenRequest {
    const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUserLogoutByTokenRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1: {
          if (tag !== 10) {
            break;
          }

          message.token = reader.string();
          continue;
        }
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skip(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): UserLogoutByTokenRequest {
    return { token: isSet(object.token) ? globalThis.String(object.token) : "" };
  },

  toJSON(message: UserLogoutByTokenRequest): unknown {
    const obj: any = {};
    if (message.token !== "") {
      obj.token = message.token;
    }
    return obj;
  },

  create(base?: DeepPartial<UserLogoutByTokenRequest>): UserLogoutByTokenRequest {
    return UserLogoutByTokenRequest.fromPartial(base ?? {});
  },
  fromPartial(object: DeepPartial<UserLogoutByTokenRequest>): UserLogoutByTokenRequest {
    const message = createBaseUserLogoutByTokenRequest();
    message.token = object.token ?? "";
    return message;
  },
};

function createBaseUserLogoutByTokenResponse(): UserLogoutByTokenResponse {
  return {};
}

export const UserLogoutByTokenResponse: MessageFns<UserLogoutByTokenResponse> = {
  encode(_: UserLogoutByTokenResponse, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): UserLogoutByTokenResponse {
    const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUserLogoutByTokenResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skip(tag & 7);
    }
    return message;
  },

  fromJSON(_: any): UserLogoutByTokenResponse {
    return {};
  },

  toJSON(_: UserLogoutByTokenResponse): unknown {
    const obj: any = {};
    return obj;
  },

  create(base?: DeepPartial<UserLogoutByTokenResponse>): UserLogoutByTokenResponse {
    return UserLogoutByTokenResponse.fromPartial(base ?? {});
  },
  fromPartial(_: DeepPartial<UserLogoutByTokenResponse>): UserLogoutByTokenResponse {
    const message = createBaseUserLogoutByTokenResponse();
    return message;
  },
};

function createBaseFindUserByTokenRequest(): FindUserByTokenRequest {
  return { token: "" };
}

export const FindUserByTokenRequest: MessageFns<FindUserByTokenRequest> = {
  encode(message: FindUserByTokenRequest, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
    if (message.token !== "") {
      writer.uint32(10).string(message.token);
    }
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): FindUserByTokenRequest {
    const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseFindUserByTokenRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1: {
          if (tag !== 10) {
            break;
          }

          message.token = reader.string();
          continue;
        }
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skip(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): FindUserByTokenRequest {
    return { token: isSet(object.token) ? globalThis.String(object.token) : "" };
  },

  toJSON(message: FindUserByTokenRequest): unknown {
    const obj: any = {};
    if (message.token !== "") {
      obj.token = message.token;
    }
    return obj;
  },

  create(base?: DeepPartial<FindUserByTokenRequest>): FindUserByTokenRequest {
    return FindUserByTokenRequest.fromPartial(base ?? {});
  },
  fromPartial(object: DeepPartial<FindUserByTokenRequest>): FindUserByTokenRequest {
    const message = createBaseFindUserByTokenRequest();
    message.token = object.token ?? "";
    return message;
  },
};

function createBaseUserVerifyOtpRequest(): UserVerifyOtpRequest {
  return { otp: "" };
}

export const UserVerifyOtpRequest: MessageFns<UserVerifyOtpRequest> = {
  encode(message: UserVerifyOtpRequest, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
    if (message.otp !== "") {
      writer.uint32(10).string(message.otp);
    }
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): UserVerifyOtpRequest {
    const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUserVerifyOtpRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1: {
          if (tag !== 10) {
            break;
          }

          message.otp = reader.string();
          continue;
        }
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skip(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): UserVerifyOtpRequest {
    return { otp: isSet(object.otp) ? globalThis.String(object.otp) : "" };
  },

  toJSON(message: UserVerifyOtpRequest): unknown {
    const obj: any = {};
    if (message.otp !== "") {
      obj.otp = message.otp;
    }
    return obj;
  },

  create(base?: DeepPartial<UserVerifyOtpRequest>): UserVerifyOtpRequest {
    return UserVerifyOtpRequest.fromPartial(base ?? {});
  },
  fromPartial(object: DeepPartial<UserVerifyOtpRequest>): UserVerifyOtpRequest {
    const message = createBaseUserVerifyOtpRequest();
    message.otp = object.otp ?? "";
    return message;
  },
};

function createBaseUserVerifyOtpResponse(): UserVerifyOtpResponse {
  return { accessToken: "", refreshToken: "" };
}

export const UserVerifyOtpResponse: MessageFns<UserVerifyOtpResponse> = {
  encode(message: UserVerifyOtpResponse, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
    if (message.accessToken !== "") {
      writer.uint32(10).string(message.accessToken);
    }
    if (message.refreshToken !== "") {
      writer.uint32(18).string(message.refreshToken);
    }
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): UserVerifyOtpResponse {
    const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUserVerifyOtpResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1: {
          if (tag !== 10) {
            break;
          }

          message.accessToken = reader.string();
          continue;
        }
        case 2: {
          if (tag !== 18) {
            break;
          }

          message.refreshToken = reader.string();
          continue;
        }
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skip(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): UserVerifyOtpResponse {
    return {
      accessToken: isSet(object.accessToken) ? globalThis.String(object.accessToken) : "",
      refreshToken: isSet(object.refreshToken) ? globalThis.String(object.refreshToken) : "",
    };
  },

  toJSON(message: UserVerifyOtpResponse): unknown {
    const obj: any = {};
    if (message.accessToken !== "") {
      obj.accessToken = message.accessToken;
    }
    if (message.refreshToken !== "") {
      obj.refreshToken = message.refreshToken;
    }
    return obj;
  },

  create(base?: DeepPartial<UserVerifyOtpResponse>): UserVerifyOtpResponse {
    return UserVerifyOtpResponse.fromPartial(base ?? {});
  },
  fromPartial(object: DeepPartial<UserVerifyOtpResponse>): UserVerifyOtpResponse {
    const message = createBaseUserVerifyOtpResponse();
    message.accessToken = object.accessToken ?? "";
    message.refreshToken = object.refreshToken ?? "";
    return message;
  },
};

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends globalThis.Array<infer U> ? globalThis.Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}

export interface MessageFns<T> {
  encode(message: T, writer?: BinaryWriter): BinaryWriter;
  decode(input: BinaryReader | Uint8Array, length?: number): T;
  fromJSON(object: any): T;
  toJSON(message: T): unknown;
  create(base?: DeepPartial<T>): T;
  fromPartial(object: DeepPartial<T>): T;
}
