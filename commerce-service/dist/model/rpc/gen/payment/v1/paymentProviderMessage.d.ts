import { BinaryReader, BinaryWriter } from "@bufbuild/protobuf/wire";
export declare const protobufPackage = "payment_v1";
export declare enum ProviderMethod {
    BANK = 0,
    CRYPTO_CURRENCY = 1,
    DEBIT = 2,
    CREDIT = 3,
    CASH_ON_DELIVERY = 4,
    UNRECOGNIZED = -1
}
export declare function providerMethodFromJSON(object: any): ProviderMethod;
export declare function providerMethodToJSON(object: ProviderMethod): string;
export interface Provider {
    id: string;
    name: string;
    method: ProviderMethod;
    createdAt: Date | undefined;
    updatedAt: Date | undefined;
    discardedAt: Date | undefined;
}
export interface FindPaymentProvidersRequest {
    name: string;
}
export interface FindPaymentProvidersResponse {
    data: Provider[];
}
export interface FindPaymentProviderByIdRequest {
    id: string;
}
export declare const Provider: MessageFns<Provider>;
export declare const FindPaymentProvidersRequest: MessageFns<FindPaymentProvidersRequest>;
export declare const FindPaymentProvidersResponse: MessageFns<FindPaymentProvidersResponse>;
export declare const FindPaymentProviderByIdRequest: MessageFns<FindPaymentProviderByIdRequest>;
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
