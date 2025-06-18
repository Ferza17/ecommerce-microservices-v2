import { BinaryReader, BinaryWriter } from "@bufbuild/protobuf/wire";
export declare const protobufPackage = "commerce_v1";
export interface OrderItem {
    id: string;
    productId: string;
    qty: number;
    price: number;
    cratedAt: Date | undefined;
    updatedAt: Date | undefined;
    userId: string;
}
export interface CreateOrderItemRequest {
    productId: string;
    userId: string;
    qty: number;
    price: number;
}
export interface UpdateOrderItemByIdRequest {
    id: string;
    productId: string;
    userId: string;
    qty: number;
    price: number;
}
export interface UpdateOrderItemByIdResponse {
    id: string;
}
export interface FindOrderWithPaginationRequest {
    userId: string;
    productIds: string[];
    page: number;
    limit: number;
}
export interface FindOrderItemsWithPaginationResponse {
    items: OrderItem[];
    page: number;
    limit: number;
    total: number;
}
export interface FindOrderItemByIdRequest {
    id: string;
}
export interface DeleteOrderItemByIdRequest {
    id: string;
}
export interface DeleteOrderItemByIdResponse {
    message: string;
}
export declare const OrderItem: MessageFns<OrderItem>;
export declare const CreateOrderItemRequest: MessageFns<CreateOrderItemRequest>;
export declare const UpdateOrderItemByIdRequest: MessageFns<UpdateOrderItemByIdRequest>;
export declare const UpdateOrderItemByIdResponse: MessageFns<UpdateOrderItemByIdResponse>;
export declare const FindOrderWithPaginationRequest: MessageFns<FindOrderWithPaginationRequest>;
export declare const FindOrderItemsWithPaginationResponse: MessageFns<FindOrderItemsWithPaginationResponse>;
export declare const FindOrderItemByIdRequest: MessageFns<FindOrderItemByIdRequest>;
export declare const DeleteOrderItemByIdRequest: MessageFns<DeleteOrderItemByIdRequest>;
export declare const DeleteOrderItemByIdResponse: MessageFns<DeleteOrderItemByIdResponse>;
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
