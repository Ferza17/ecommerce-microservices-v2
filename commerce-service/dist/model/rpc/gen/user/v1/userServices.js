"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.UserServiceClient = exports.UserServiceService = exports.protobufPackage = void 0;
const grpc_js_1 = require("@grpc/grpc-js");
const userMessage_1 = require("./userMessage");
exports.protobufPackage = "user_v1";
exports.UserServiceService = {
    findUserById: {
        path: "/user_v1.UserService/FindUserById",
        requestStream: false,
        responseStream: false,
        requestSerialize: (value) => Buffer.from(userMessage_1.FindUserByIdRequest.encode(value).finish()),
        requestDeserialize: (value) => userMessage_1.FindUserByIdRequest.decode(value),
        responseSerialize: (value) => Buffer.from(userMessage_1.User.encode(value).finish()),
        responseDeserialize: (value) => userMessage_1.User.decode(value),
    },
    findUserByEmailAndPassword: {
        path: "/user_v1.UserService/FindUserByEmailAndPassword",
        requestStream: false,
        responseStream: false,
        requestSerialize: (value) => Buffer.from(userMessage_1.FindUserByEmailAndPasswordRequest.encode(value).finish()),
        requestDeserialize: (value) => userMessage_1.FindUserByEmailAndPasswordRequest.decode(value),
        responseSerialize: (value) => Buffer.from(userMessage_1.User.encode(value).finish()),
        responseDeserialize: (value) => userMessage_1.User.decode(value),
    },
};
exports.UserServiceClient = (0, grpc_js_1.makeGenericClientConstructor)(exports.UserServiceService, "user_v1.UserService");
//# sourceMappingURL=userServices.js.map