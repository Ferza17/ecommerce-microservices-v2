import { BinaryReader, BinaryWriter } from "@bufbuild/protobuf/wire";
export declare const protobufPackage = "commerce_v1";
export interface CartItem {
    id: string;
    productId: string;
    userId: string;
    qty: number;
    price: number;
    cratedAt: Date | undefined;
    updatedAt: Date | undefined;
}
export interface CreateCartItemRequest {
    productId: string;
    userId: string;
    qty: number;
    price: number;
}
export interface CreateCartItemResponse {
    id: string;
}
export interface UpdateCartItemByIdRequest {
    id: string;
    productId: string;
    userId: string;
    qty: number;
    price: number;
}
export interface UpdateCartItemByIdResponse {
    id: string;
}
export interface FindCartItemsWithPaginationRequest {
    userId: string;
    productIds: string[];
    page: number;
    limit: number;
}
export interface FindCartItemsWithPaginationResponse {
    items: CartItem[];
    page: number;
    limit: number;
    total: number;
}
export interface FindCartItemByIdRequest {
    id: string;
}
export interface DeleteCartItemByIdRequest {
    id: string;
}
export interface DeleteCartItemByIdResponse {
    message: string;
}
export declare const CartItem: MessageFns<CartItem>;
export declare const CreateCartItemRequest: MessageFns<CreateCartItemRequest>;
export declare const CreateCartItemResponse: MessageFns<CreateCartItemResponse>;
export declare const UpdateCartItemByIdRequest: MessageFns<UpdateCartItemByIdRequest>;
export declare const UpdateCartItemByIdResponse: MessageFns<UpdateCartItemByIdResponse>;
export declare const FindCartItemsWithPaginationRequest: MessageFns<FindCartItemsWithPaginationRequest>;
export declare const FindCartItemsWithPaginationResponse: MessageFns<FindCartItemsWithPaginationResponse>;
export declare const FindCartItemByIdRequest: MessageFns<FindCartItemByIdRequest>;
export declare const DeleteCartItemByIdRequest: MessageFns<DeleteCartItemByIdRequest>;
export declare const DeleteCartItemByIdResponse: MessageFns<DeleteCartItemByIdResponse>;
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
