import { BinaryReader, BinaryWriter } from "@bufbuild/protobuf/wire";
export declare const protobufPackage = "user_v1";
export interface User {
    id: string;
    name: string;
    email: string;
    password: string;
    createdAt?: Date | undefined;
    updatedAt?: Date | undefined;
    discardedAt?: Date | undefined;
}
export interface FindUserByIdRequest {
    id: string;
}
export interface CreateUserRequest {
    name: string;
    email: string;
    password: string;
}
export interface CreateUserResponse {
    id: string;
}
export interface UpdateUserByIdRequest {
    id: string;
    name?: string | undefined;
    email?: string | undefined;
    password?: string | undefined;
}
export interface UpdateUserByIdResponse {
    id: string;
}
export interface FindUserByEmailAndPasswordRequest {
    email: string;
    password: string;
}
export declare const User: MessageFns<User>;
export declare const FindUserByIdRequest: MessageFns<FindUserByIdRequest>;
export declare const CreateUserRequest: MessageFns<CreateUserRequest>;
export declare const CreateUserResponse: MessageFns<CreateUserResponse>;
export declare const UpdateUserByIdRequest: MessageFns<UpdateUserByIdRequest>;
export declare const UpdateUserByIdResponse: MessageFns<UpdateUserByIdResponse>;
export declare const FindUserByEmailAndPasswordRequest: MessageFns<FindUserByEmailAndPasswordRequest>;
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
