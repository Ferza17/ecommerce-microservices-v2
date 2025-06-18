import { BinaryReader, BinaryWriter } from "@bufbuild/protobuf/wire";
export declare const protobufPackage = "user_v1";
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
export declare const UserLoginByEmailAndPasswordRequest: MessageFns<UserLoginByEmailAndPasswordRequest>;
export declare const UserLogoutByTokenRequest: MessageFns<UserLogoutByTokenRequest>;
export declare const UserLogoutByTokenResponse: MessageFns<UserLogoutByTokenResponse>;
export declare const FindUserByTokenRequest: MessageFns<FindUserByTokenRequest>;
export declare const UserVerifyOtpRequest: MessageFns<UserVerifyOtpRequest>;
export declare const UserVerifyOtpResponse: MessageFns<UserVerifyOtpResponse>;
type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;
export type DeepPartial<T> = T extends Builtin ? T : T extends globalThis.Array<infer U> ? globalThis.Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export interface MessageFns<T> {
    encode(message: T, writer?: BinaryWriter): BinaryWriter;
    decode(input: BinaryReader | Uint8Array, length?: number): T;
    fromJSON(object: any): T;
    toJSON(message: T): unknown;
    create(base?: DeepPartial<T>): T;
    fromPartial(object: DeepPartial<T>): T;
}
export {};
