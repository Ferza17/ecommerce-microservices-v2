import { type CallOptions, ChannelCredentials, Client, type ClientOptions, type ClientUnaryCall, type handleUnaryCall, Metadata, type ServiceError, type UntypedServiceImplementation } from "@grpc/grpc-js";
import { FindPaymentProviderByIdRequest, FindPaymentProvidersRequest, FindPaymentProvidersResponse, Provider } from "./paymentProviderMessage";
export declare const protobufPackage = "payment_v1";
export type PaymentProviderServiceService = typeof PaymentProviderServiceService;
export declare const PaymentProviderServiceService: {
    readonly findPaymentProviders: {
        readonly path: "/payment_v1.PaymentProviderService/FindPaymentProviders";
        readonly requestStream: false;
        readonly responseStream: false;
        readonly requestSerialize: (value: FindPaymentProvidersRequest) => Buffer<ArrayBuffer>;
        readonly requestDeserialize: (value: Buffer) => FindPaymentProvidersRequest;
        readonly responseSerialize: (value: FindPaymentProvidersResponse) => Buffer<ArrayBuffer>;
        readonly responseDeserialize: (value: Buffer) => FindPaymentProvidersResponse;
    };
    readonly findPaymentProviderById: {
        readonly path: "/payment_v1.PaymentProviderService/FindPaymentProviderById";
        readonly requestStream: false;
        readonly responseStream: false;
        readonly requestSerialize: (value: FindPaymentProviderByIdRequest) => Buffer<ArrayBuffer>;
        readonly requestDeserialize: (value: Buffer) => FindPaymentProviderByIdRequest;
        readonly responseSerialize: (value: Provider) => Buffer<ArrayBuffer>;
        readonly responseDeserialize: (value: Buffer) => Provider;
    };
};
export interface PaymentProviderServiceServer extends UntypedServiceImplementation {
    findPaymentProviders: handleUnaryCall<FindPaymentProvidersRequest, FindPaymentProvidersResponse>;
    findPaymentProviderById: handleUnaryCall<FindPaymentProviderByIdRequest, Provider>;
}
export interface PaymentProviderServiceClient extends Client {
    findPaymentProviders(request: FindPaymentProvidersRequest, callback: (error: ServiceError | null, response: FindPaymentProvidersResponse) => void): ClientUnaryCall;
    findPaymentProviders(request: FindPaymentProvidersRequest, metadata: Metadata, callback: (error: ServiceError | null, response: FindPaymentProvidersResponse) => void): ClientUnaryCall;
    findPaymentProviders(request: FindPaymentProvidersRequest, metadata: Metadata, options: Partial<CallOptions>, callback: (error: ServiceError | null, response: FindPaymentProvidersResponse) => void): ClientUnaryCall;
    findPaymentProviderById(request: FindPaymentProviderByIdRequest, callback: (error: ServiceError | null, response: Provider) => void): ClientUnaryCall;
    findPaymentProviderById(request: FindPaymentProviderByIdRequest, metadata: Metadata, callback: (error: ServiceError | null, response: Provider) => void): ClientUnaryCall;
    findPaymentProviderById(request: FindPaymentProviderByIdRequest, metadata: Metadata, options: Partial<CallOptions>, callback: (error: ServiceError | null, response: Provider) => void): ClientUnaryCall;
}
export declare const PaymentProviderServiceClient: {
    new (address: string, credentials: ChannelCredentials, options?: Partial<ClientOptions>): PaymentProviderServiceClient;
    service: typeof PaymentProviderServiceService;
    serviceName: string;
};
