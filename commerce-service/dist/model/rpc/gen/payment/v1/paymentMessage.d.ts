import { BinaryReader, BinaryWriter } from "@bufbuild/protobuf/wire";
import { Provider } from "./paymentProviderMessage";
export declare const protobufPackage = "payment_v1";
export declare enum PaymentStatus {
    PENDING = 0,
    PARTIAL = 1,
    SUCCESS = 2,
    FAILED = 3,
    UNRECOGNIZED = -1
}
export declare function paymentStatusFromJSON(object: any): PaymentStatus;
export declare function paymentStatusToJSON(object: PaymentStatus): string;
export interface PaymentItem {
    id: string;
    productId: string;
    amount: number;
    qty: number;
    cratedAt: Date | undefined;
    updatedAt: Date | undefined;
    discardedAt: Date | undefined;
}
export interface Payment {
    id: string;
    code: string;
    Items: PaymentItem[];
    totalPrice: number;
    status: PaymentStatus;
    provider: Provider | undefined;
    userId: string;
    createdAt: Date | undefined;
    updatedAt: Date | undefined;
    discardedAt: Date | undefined;
}
export interface CreatePaymentRequest {
    items: PaymentItem[];
    userId: string;
    amount: number;
    providerId: string;
}
export interface PaymentOrderDelayedCancelledRequest {
    id: string;
}
export interface CallBackPaymentRequest {
    paymentId: string;
    amount: number;
}
export interface FindPaymentByIdRequest {
    id: string;
}
export interface FindPaymentByUserIdAndStatusRequest {
    userId: string;
    status: PaymentStatus;
}
export declare const PaymentItem: MessageFns<PaymentItem>;
export declare const Payment: MessageFns<Payment>;
export declare const CreatePaymentRequest: MessageFns<CreatePaymentRequest>;
export declare const PaymentOrderDelayedCancelledRequest: MessageFns<PaymentOrderDelayedCancelledRequest>;
export declare const CallBackPaymentRequest: MessageFns<CallBackPaymentRequest>;
export declare const FindPaymentByIdRequest: MessageFns<FindPaymentByIdRequest>;
export declare const FindPaymentByUserIdAndStatusRequest: MessageFns<FindPaymentByUserIdAndStatusRequest>;
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
