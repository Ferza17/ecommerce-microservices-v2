import { BinaryReader, BinaryWriter } from "@bufbuild/protobuf/wire";
export declare const protobufPackage = "commerce_v1";
export interface WishlistItem {
    id: string;
    productId: string;
    userId: string;
}
export interface CreateWishlistItemRequest {
    productId: string;
    userId: string;
}
export interface CreateWishlistItemResponse {
    id: string;
}
export interface FindWishlistItemWithPaginationRequest {
    userId: string;
    productIds: string[];
    page: number;
    limit: number;
}
export interface FindWishlistItemWithPaginationResponse {
    items: WishlistItem[];
    page: number;
    limit: number;
}
export interface DeleteWishlistItemByIdRequest {
    id: string;
}
export interface DeleteWishlistItemByIdResponse {
    userId: string;
}
export declare const WishlistItem: MessageFns<WishlistItem>;
export declare const CreateWishlistItemRequest: MessageFns<CreateWishlistItemRequest>;
export declare const CreateWishlistItemResponse: MessageFns<CreateWishlistItemResponse>;
export declare const FindWishlistItemWithPaginationRequest: MessageFns<FindWishlistItemWithPaginationRequest>;
export declare const FindWishlistItemWithPaginationResponse: MessageFns<FindWishlistItemWithPaginationResponse>;
export declare const DeleteWishlistItemByIdRequest: MessageFns<DeleteWishlistItemByIdRequest>;
export declare const DeleteWishlistItemByIdResponse: MessageFns<DeleteWishlistItemByIdResponse>;
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
