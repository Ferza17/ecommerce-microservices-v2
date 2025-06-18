"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.PaymentProviderServiceClient = exports.PaymentProviderServiceService = exports.protobufPackage = void 0;
const grpc_js_1 = require("@grpc/grpc-js");
const paymentProviderMessage_1 = require("./paymentProviderMessage");
exports.protobufPackage = "payment_v1";
exports.PaymentProviderServiceService = {
    findPaymentProviders: {
        path: "/payment_v1.PaymentProviderService/FindPaymentProviders",
        requestStream: false,
        responseStream: false,
        requestSerialize: (value) => Buffer.from(paymentProviderMessage_1.FindPaymentProvidersRequest.encode(value).finish()),
        requestDeserialize: (value) => paymentProviderMessage_1.FindPaymentProvidersRequest.decode(value),
        responseSerialize: (value) => Buffer.from(paymentProviderMessage_1.FindPaymentProvidersResponse.encode(value).finish()),
        responseDeserialize: (value) => paymentProviderMessage_1.FindPaymentProvidersResponse.decode(value),
    },
    findPaymentProviderById: {
        path: "/payment_v1.PaymentProviderService/FindPaymentProviderById",
        requestStream: false,
        responseStream: false,
        requestSerialize: (value) => Buffer.from(paymentProviderMessage_1.FindPaymentProviderByIdRequest.encode(value).finish()),
        requestDeserialize: (value) => paymentProviderMessage_1.FindPaymentProviderByIdRequest.decode(value),
        responseSerialize: (value) => Buffer.from(paymentProviderMessage_1.Provider.encode(value).finish()),
        responseDeserialize: (value) => paymentProviderMessage_1.Provider.decode(value),
    },
};
exports.PaymentProviderServiceClient = (0, grpc_js_1.makeGenericClientConstructor)(exports.PaymentProviderServiceService, "payment_v1.PaymentProviderService");
//# sourceMappingURL=paymentProviderServices.js.map