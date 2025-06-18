import { BinaryReader, BinaryWriter } from "@bufbuild/protobuf/wire";
export declare const protobufPackage = "event_v1";
export interface EventStore {
    id: string;
    requestId: string;
    service: string;
    eventType: string;
    status: string;
    payload: {
        [key: string]: any;
    } | undefined;
    previousState?: {
        [key: string]: any;
    } | undefined;
    createdAt: Date | undefined;
    updatedAt: Date | undefined;
}
export interface CreateEventStoreResponse {
    id: string;
}
export declare const EventStore: MessageFns<EventStore>;
export declare const CreateEventStoreResponse: MessageFns<CreateEventStoreResponse>;
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
