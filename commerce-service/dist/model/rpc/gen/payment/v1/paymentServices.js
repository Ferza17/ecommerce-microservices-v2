"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.PaymentServiceClient = exports.PaymentServiceService = exports.protobufPackage = void 0;
const grpc_js_1 = require("@grpc/grpc-js");
const paymentMessage_1 = require("./paymentMessage");
exports.protobufPackage = "payment_v1";
exports.PaymentServiceService = {
    findPaymentById: {
        path: "/payment_v1.PaymentService/FindPaymentById",
        requestStream: false,
        responseStream: false,
        requestSerialize: (value) => Buffer.from(paymentMessage_1.FindPaymentByIdRequest.encode(value).finish()),
        requestDeserialize: (value) => paymentMessage_1.FindPaymentByIdRequest.decode(value),
        responseSerialize: (value) => Buffer.from(paymentMessage_1.Payment.encode(value).finish()),
        responseDeserialize: (value) => paymentMessage_1.Payment.decode(value),
    },
    findPaymentByUserIdAndStatus: {
        path: "/payment_v1.PaymentService/FindPaymentByUserIdAndStatus",
        requestStream: false,
        responseStream: false,
        requestSerialize: (value) => Buffer.from(paymentMessage_1.FindPaymentByUserIdAndStatusRequest.encode(value).finish()),
        requestDeserialize: (value) => paymentMessage_1.FindPaymentByUserIdAndStatusRequest.decode(value),
        responseSerialize: (value) => Buffer.from(paymentMessage_1.Payment.encode(value).finish()),
        responseDeserialize: (value) => paymentMessage_1.Payment.decode(value),
    },
};
exports.PaymentServiceClient = (0, grpc_js_1.makeGenericClientConstructor)(exports.PaymentServiceService, "payment_v1.PaymentService");
//# sourceMappingURL=paymentServices.js.map