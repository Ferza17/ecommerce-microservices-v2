"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.AuthServiceClient = exports.AuthServiceService = exports.protobufPackage = void 0;
const grpc_js_1 = require("@grpc/grpc-js");
const authMessage_1 = require("./authMessage");
const userMessage_1 = require("./userMessage");
exports.protobufPackage = "user_v1";
exports.AuthServiceService = {
    userLogoutByToken: {
        path: "/user_v1.AuthService/UserLogoutByToken",
        requestStream: false,
        responseStream: false,
        requestSerialize: (value) => Buffer.from(authMessage_1.UserLogoutByTokenRequest.encode(value).finish()),
        requestDeserialize: (value) => authMessage_1.UserLogoutByTokenRequest.decode(value),
        responseSerialize: (value) => Buffer.from(authMessage_1.UserLogoutByTokenResponse.encode(value).finish()),
        responseDeserialize: (value) => authMessage_1.UserLogoutByTokenResponse.decode(value),
    },
    userVerifyOtp: {
        path: "/user_v1.AuthService/UserVerifyOtp",
        requestStream: false,
        responseStream: false,
        requestSerialize: (value) => Buffer.from(authMessage_1.UserVerifyOtpRequest.encode(value).finish()),
        requestDeserialize: (value) => authMessage_1.UserVerifyOtpRequest.decode(value),
        responseSerialize: (value) => Buffer.from(authMessage_1.UserVerifyOtpResponse.encode(value).finish()),
        responseDeserialize: (value) => authMessage_1.UserVerifyOtpResponse.decode(value),
    },
    findUserByToken: {
        path: "/user_v1.AuthService/FindUserByToken",
        requestStream: false,
        responseStream: false,
        requestSerialize: (value) => Buffer.from(authMessage_1.FindUserByTokenRequest.encode(value).finish()),
        requestDeserialize: (value) => authMessage_1.FindUserByTokenRequest.decode(value),
        responseSerialize: (value) => Buffer.from(userMessage_1.User.encode(value).finish()),
        responseDeserialize: (value) => userMessage_1.User.decode(value),
    },
};
exports.AuthServiceClient = (0, grpc_js_1.makeGenericClientConstructor)(exports.AuthServiceService, "user_v1.AuthService");
//# sourceMappingURL=authServices.js.map