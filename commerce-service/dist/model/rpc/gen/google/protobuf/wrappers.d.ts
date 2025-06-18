import { BinaryReader, BinaryWriter } from "@bufbuild/protobuf/wire";
export declare const protobufPackage = "google.protobuf";
export interface DoubleValue {
    value: number;
}
export interface FloatValue {
    value: number;
}
export interface Int64Value {
    value: number;
}
export interface UInt64Value {
    value: number;
}
export interface Int32Value {
    value: number;
}
export interface UInt32Value {
    value: number;
}
export interface BoolValue {
    value: boolean;
}
export interface StringValue {
    value: string;
}
export interface BytesValue {
    value: Buffer;
}
export declare const DoubleValue: MessageFns<DoubleValue>;
export declare const FloatValue: MessageFns<FloatValue>;
export declare const Int64Value: MessageFns<Int64Value>;
export declare const UInt64Value: MessageFns<UInt64Value>;
export declare const Int32Value: MessageFns<Int32Value>;
export declare const UInt32Value: MessageFns<UInt32Value>;
export declare const BoolValue: MessageFns<BoolValue>;
export declare const StringValue: MessageFns<StringValue>;
export declare const BytesValue: MessageFns<BytesValue>;
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
