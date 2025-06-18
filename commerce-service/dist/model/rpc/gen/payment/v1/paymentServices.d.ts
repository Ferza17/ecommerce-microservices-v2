import { type CallOptions, ChannelCredentials, Client, type ClientOptions, type ClientUnaryCall, type handleUnaryCall, Metadata, type ServiceError, type UntypedServiceImplementation } from "@grpc/grpc-js";
import { FindPaymentByIdRequest, FindPaymentByUserIdAndStatusRequest, Payment } from "./paymentMessage";
export declare const protobufPackage = "payment_v1";
export type PaymentServiceService = typeof PaymentServiceService;
export declare const PaymentServiceService: {
    readonly findPaymentById: {
        readonly path: "/payment_v1.PaymentService/FindPaymentById";
        readonly requestStream: false;
        readonly responseStream: false;
        readonly requestSerialize: (value: FindPaymentByIdRequest) => Buffer<ArrayBuffer>;
        readonly requestDeserialize: (value: Buffer) => FindPaymentByIdRequest;
        readonly responseSerialize: (value: Payment) => Buffer<ArrayBuffer>;
        readonly responseDeserialize: (value: Buffer) => Payment;
    };
    readonly findPaymentByUserIdAndStatus: {
        readonly path: "/payment_v1.PaymentService/FindPaymentByUserIdAndStatus";
        readonly requestStream: false;
        readonly responseStream: false;
        readonly requestSerialize: (value: FindPaymentByUserIdAndStatusRequest) => Buffer<ArrayBuffer>;
        readonly requestDeserialize: (value: Buffer) => FindPaymentByUserIdAndStatusRequest;
        readonly responseSerialize: (value: Payment) => Buffer<ArrayBuffer>;
        readonly responseDeserialize: (value: Buffer) => Payment;
    };
};
export interface PaymentServiceServer extends UntypedServiceImplementation {
    findPaymentById: handleUnaryCall<FindPaymentByIdRequest, Payment>;
    findPaymentByUserIdAndStatus: handleUnaryCall<FindPaymentByUserIdAndStatusRequest, Payment>;
}
export interface PaymentServiceClient extends Client {
    findPaymentById(request: FindPaymentByIdRequest, callback: (error: ServiceError | null, response: Payment) => void): ClientUnaryCall;
    findPaymentById(request: FindPaymentByIdRequest, metadata: Metadata, callback: (error: ServiceError | null, response: Payment) => void): ClientUnaryCall;
    findPaymentById(request: FindPaymentByIdRequest, metadata: Metadata, options: Partial<CallOptions>, callback: (error: ServiceError | null, response: Payment) => void): ClientUnaryCall;
    findPaymentByUserIdAndStatus(request: FindPaymentByUserIdAndStatusRequest, callback: (error: ServiceError | null, response: Payment) => void): ClientUnaryCall;
    findPaymentByUserIdAndStatus(request: FindPaymentByUserIdAndStatusRequest, metadata: Metadata, callback: (error: ServiceError | null, response: Payment) => void): ClientUnaryCall;
    findPaymentByUserIdAndStatus(request: FindPaymentByUserIdAndStatusRequest, metadata: Metadata, options: Partial<CallOptions>, callback: (error: ServiceError | null, response: Payment) => void): ClientUnaryCall;
}
export declare const PaymentServiceClient: {
    new (address: string, credentials: ChannelCredentials, options?: Partial<ClientOptions>): PaymentServiceClient;
    service: typeof PaymentServiceService;
    serviceName: string;
};
