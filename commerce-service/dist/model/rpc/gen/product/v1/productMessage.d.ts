import { BinaryReader, BinaryWriter } from "@bufbuild/protobuf/wire";
export declare const protobufPackage = "product_v1";
export interface Product {
    id: string;
    name: string;
    description: string;
    uom: string;
    image: string;
    price: number;
    stock: number;
    createdAt: Date | undefined;
    updatedAt: Date | undefined;
    discardedAt: Date | undefined;
}
export interface FindProductsWithPaginationRequest {
    ids: string[];
    name: string[];
    page: number;
    limit: number;
}
export interface FindProductsWithPaginationResponse {
    data: Product[];
    limit: number;
    page: number;
    total: number;
}
export interface FindProductByIdRequest {
    id: string;
}
export interface CreateProductRequest {
    name: string;
    description: string;
    uom: string;
    image: string;
    price: number;
    stock: number;
}
export interface CreateProductResponse {
    id: string;
}
export interface UpdateProductByIdRequest {
    id: string;
    name?: string | undefined;
    description?: string | undefined;
    uom?: string | undefined;
    image?: string | undefined;
    price?: number | undefined;
    stock?: number | undefined;
}
export interface DeleteProductByIdRequest {
    id: string;
}
export interface DeleteProductByIdResponse {
    message: string;
}
export declare const Product: MessageFns<Product>;
export declare const FindProductsWithPaginationRequest: MessageFns<FindProductsWithPaginationRequest>;
export declare const FindProductsWithPaginationResponse: MessageFns<FindProductsWithPaginationResponse>;
export declare const FindProductByIdRequest: MessageFns<FindProductByIdRequest>;
export declare const CreateProductRequest: MessageFns<CreateProductRequest>;
export declare const CreateProductResponse: MessageFns<CreateProductResponse>;
export declare const UpdateProductByIdRequest: MessageFns<UpdateProductByIdRequest>;
export declare const DeleteProductByIdRequest: MessageFns<DeleteProductByIdRequest>;
export declare const DeleteProductByIdResponse: MessageFns<DeleteProductByIdResponse>;
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
