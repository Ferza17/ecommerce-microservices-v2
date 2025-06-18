"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.UserVerifyOtpResponse = exports.UserVerifyOtpRequest = exports.FindUserByTokenRequest = exports.UserLogoutByTokenResponse = exports.UserLogoutByTokenRequest = exports.UserLoginByEmailAndPasswordRequest = exports.protobufPackage = void 0;
const wire_1 = require("@bufbuild/protobuf/wire");
exports.protobufPackage = "user_v1";
function createBaseUserLoginByEmailAndPasswordRequest() {
    return { email: "", password: "" };
}
exports.UserLoginByEmailAndPasswordRequest = {
    encode(message, writer = new wire_1.BinaryWriter()) {
        if (message.email !== "") {
            writer.uint32(10).string(message.email);
        }
        if (message.password !== "") {
            writer.uint32(18).string(message.password);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof wire_1.BinaryReader ? input : new wire_1.BinaryReader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseUserLoginByEmailAndPasswordRequest();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1: {
                    if (tag !== 10) {
                        break;
                    }
                    message.email = reader.string();
                    continue;
                }
                case 2: {
                    if (tag !== 18) {
                        break;
                    }
                    message.password = reader.string();
                    continue;
                }
            }
            if ((tag & 7) === 4 || tag === 0) {
                break;
            }
            reader.skip(tag & 7);
        }
        return message;
    },
    fromJSON(object) {
        return {
            email: isSet(object.email) ? globalThis.String(object.email) : "",
            password: isSet(object.password) ? globalThis.String(object.password) : "",
        };
    },
    toJSON(message) {
        const obj = {};
        if (message.email !== "") {
            obj.email = message.email;
        }
        if (message.password !== "") {
            obj.password = message.password;
        }
        return obj;
    },
    create(base) {
        return exports.UserLoginByEmailAndPasswordRequest.fromPartial(base ?? {});
    },
    fromPartial(object) {
        const message = createBaseUserLoginByEmailAndPasswordRequest();
        message.email = object.email ?? "";
        message.password = object.password ?? "";
        return message;
    },
};
function createBaseUserLogoutByTokenRequest() {
    return { token: "" };
}
exports.UserLogoutByTokenRequest = {
    encode(message, writer = new wire_1.BinaryWriter()) {
        if (message.token !== "") {
            writer.uint32(10).string(message.token);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof wire_1.BinaryReader ? input : new wire_1.BinaryReader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseUserLogoutByTokenRequest();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1: {
                    if (tag !== 10) {
                        break;
                    }
                    message.token = reader.string();
                    continue;
                }
            }
            if ((tag & 7) === 4 || tag === 0) {
                break;
            }
            reader.skip(tag & 7);
        }
        return message;
    },
    fromJSON(object) {
        return { token: isSet(object.token) ? globalThis.String(object.token) : "" };
    },
    toJSON(message) {
        const obj = {};
        if (message.token !== "") {
            obj.token = message.token;
        }
        return obj;
    },
    create(base) {
        return exports.UserLogoutByTokenRequest.fromPartial(base ?? {});
    },
    fromPartial(object) {
        const message = createBaseUserLogoutByTokenRequest();
        message.token = object.token ?? "";
        return message;
    },
};
function createBaseUserLogoutByTokenResponse() {
    return {};
}
exports.UserLogoutByTokenResponse = {
    encode(_, writer = new wire_1.BinaryWriter()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof wire_1.BinaryReader ? input : new wire_1.BinaryReader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseUserLogoutByTokenResponse();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
            }
            if ((tag & 7) === 4 || tag === 0) {
                break;
            }
            reader.skip(tag & 7);
        }
        return message;
    },
    fromJSON(_) {
        return {};
    },
    toJSON(_) {
        const obj = {};
        return obj;
    },
    create(base) {
        return exports.UserLogoutByTokenResponse.fromPartial(base ?? {});
    },
    fromPartial(_) {
        const message = createBaseUserLogoutByTokenResponse();
        return message;
    },
};
function createBaseFindUserByTokenRequest() {
    return { token: "" };
}
exports.FindUserByTokenRequest = {
    encode(message, writer = new wire_1.BinaryWriter()) {
        if (message.token !== "") {
            writer.uint32(10).string(message.token);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof wire_1.BinaryReader ? input : new wire_1.BinaryReader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseFindUserByTokenRequest();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1: {
                    if (tag !== 10) {
                        break;
                    }
                    message.token = reader.string();
                    continue;
                }
            }
            if ((tag & 7) === 4 || tag === 0) {
                break;
            }
            reader.skip(tag & 7);
        }
        return message;
    },
    fromJSON(object) {
        return { token: isSet(object.token) ? globalThis.String(object.token) : "" };
    },
    toJSON(message) {
        const obj = {};
        if (message.token !== "") {
            obj.token = message.token;
        }
        return obj;
    },
    create(base) {
        return exports.FindUserByTokenRequest.fromPartial(base ?? {});
    },
    fromPartial(object) {
        const message = createBaseFindUserByTokenRequest();
        message.token = object.token ?? "";
        return message;
    },
};
function createBaseUserVerifyOtpRequest() {
    return { otp: "" };
}
exports.UserVerifyOtpRequest = {
    encode(message, writer = new wire_1.BinaryWriter()) {
        if (message.otp !== "") {
            writer.uint32(10).string(message.otp);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof wire_1.BinaryReader ? input : new wire_1.BinaryReader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseUserVerifyOtpRequest();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1: {
                    if (tag !== 10) {
                        break;
                    }
                    message.otp = reader.string();
                    continue;
                }
            }
            if ((tag & 7) === 4 || tag === 0) {
                break;
            }
            reader.skip(tag & 7);
        }
        return message;
    },
    fromJSON(object) {
        return { otp: isSet(object.otp) ? globalThis.String(object.otp) : "" };
    },
    toJSON(message) {
        const obj = {};
        if (message.otp !== "") {
            obj.otp = message.otp;
        }
        return obj;
    },
    create(base) {
        return exports.UserVerifyOtpRequest.fromPartial(base ?? {});
    },
    fromPartial(object) {
        const message = createBaseUserVerifyOtpRequest();
        message.otp = object.otp ?? "";
        return message;
    },
};
function createBaseUserVerifyOtpResponse() {
    return { accessToken: "", refreshToken: "" };
}
exports.UserVerifyOtpResponse = {
    encode(message, writer = new wire_1.BinaryWriter()) {
        if (message.accessToken !== "") {
            writer.uint32(10).string(message.accessToken);
        }
        if (message.refreshToken !== "") {
            writer.uint32(18).string(message.refreshToken);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof wire_1.BinaryReader ? input : new wire_1.BinaryReader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseUserVerifyOtpResponse();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1: {
                    if (tag !== 10) {
                        break;
                    }
                    message.accessToken = reader.string();
                    continue;
                }
                case 2: {
                    if (tag !== 18) {
                        break;
                    }
                    message.refreshToken = reader.string();
                    continue;
                }
            }
            if ((tag & 7) === 4 || tag === 0) {
                break;
            }
            reader.skip(tag & 7);
        }
        return message;
    },
    fromJSON(object) {
        return {
            accessToken: isSet(object.accessToken) ? globalThis.String(object.accessToken) : "",
            refreshToken: isSet(object.refreshToken) ? globalThis.String(object.refreshToken) : "",
        };
    },
    toJSON(message) {
        const obj = {};
        if (message.accessToken !== "") {
            obj.accessToken = message.accessToken;
        }
        if (message.refreshToken !== "") {
            obj.refreshToken = message.refreshToken;
        }
        return obj;
    },
    create(base) {
        return exports.UserVerifyOtpResponse.fromPartial(base ?? {});
    },
    fromPartial(object) {
        const message = createBaseUserVerifyOtpResponse();
        message.accessToken = object.accessToken ?? "";
        message.refreshToken = object.refreshToken ?? "";
        return message;
    },
};
function isSet(value) {
    return value !== null && value !== undefined;
}
//# sourceMappingURL=authMessage.js.map