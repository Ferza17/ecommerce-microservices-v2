"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.SendOtpEmailNotificationRequest = exports.NotificationTemplate = exports.NotificationTypeEnum = exports.protobufPackage = void 0;
exports.notificationTypeEnumFromJSON = notificationTypeEnumFromJSON;
exports.notificationTypeEnumToJSON = notificationTypeEnumToJSON;
const wire_1 = require("@bufbuild/protobuf/wire");
const struct_1 = require("../../google/protobuf/struct");
exports.protobufPackage = "notification_v1";
var NotificationTypeEnum;
(function (NotificationTypeEnum) {
    NotificationTypeEnum[NotificationTypeEnum["NOTIFICATION_EMAIL_USER_OTP"] = 0] = "NOTIFICATION_EMAIL_USER_OTP";
    NotificationTypeEnum[NotificationTypeEnum["UNRECOGNIZED"] = -1] = "UNRECOGNIZED";
})(NotificationTypeEnum || (exports.NotificationTypeEnum = NotificationTypeEnum = {}));
function notificationTypeEnumFromJSON(object) {
    switch (object) {
        case 0:
        case "NOTIFICATION_EMAIL_USER_OTP":
            return NotificationTypeEnum.NOTIFICATION_EMAIL_USER_OTP;
        case -1:
        case "UNRECOGNIZED":
        default:
            return NotificationTypeEnum.UNRECOGNIZED;
    }
}
function notificationTypeEnumToJSON(object) {
    switch (object) {
        case NotificationTypeEnum.NOTIFICATION_EMAIL_USER_OTP:
            return "NOTIFICATION_EMAIL_USER_OTP";
        case NotificationTypeEnum.UNRECOGNIZED:
        default:
            return "UNRECOGNIZED";
    }
}
function createBaseNotificationTemplate() {
    return { id: "", type: "", template: "", templateVars: undefined };
}
exports.NotificationTemplate = {
    encode(message, writer = new wire_1.BinaryWriter()) {
        if (message.id !== "") {
            writer.uint32(10).string(message.id);
        }
        if (message.type !== "") {
            writer.uint32(18).string(message.type);
        }
        if (message.template !== "") {
            writer.uint32(26).string(message.template);
        }
        if (message.templateVars !== undefined) {
            struct_1.Struct.encode(struct_1.Struct.wrap(message.templateVars), writer.uint32(34).fork()).join();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof wire_1.BinaryReader ? input : new wire_1.BinaryReader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseNotificationTemplate();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1: {
                    if (tag !== 10) {
                        break;
                    }
                    message.id = reader.string();
                    continue;
                }
                case 2: {
                    if (tag !== 18) {
                        break;
                    }
                    message.type = reader.string();
                    continue;
                }
                case 3: {
                    if (tag !== 26) {
                        break;
                    }
                    message.template = reader.string();
                    continue;
                }
                case 4: {
                    if (tag !== 34) {
                        break;
                    }
                    message.templateVars = struct_1.Struct.unwrap(struct_1.Struct.decode(reader, reader.uint32()));
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
            id: isSet(object.id) ? globalThis.String(object.id) : "",
            type: isSet(object.type) ? globalThis.String(object.type) : "",
            template: isSet(object.template) ? globalThis.String(object.template) : "",
            templateVars: isObject(object.templateVars) ? object.templateVars : undefined,
        };
    },
    toJSON(message) {
        const obj = {};
        if (message.id !== "") {
            obj.id = message.id;
        }
        if (message.type !== "") {
            obj.type = message.type;
        }
        if (message.template !== "") {
            obj.template = message.template;
        }
        if (message.templateVars !== undefined) {
            obj.templateVars = message.templateVars;
        }
        return obj;
    },
    create(base) {
        return exports.NotificationTemplate.fromPartial(base ?? {});
    },
    fromPartial(object) {
        const message = createBaseNotificationTemplate();
        message.id = object.id ?? "";
        message.type = object.type ?? "";
        message.template = object.template ?? "";
        message.templateVars = object.templateVars ?? undefined;
        return message;
    },
};
function createBaseSendOtpEmailNotificationRequest() {
    return { otp: "", email: "", notificationType: 0 };
}
exports.SendOtpEmailNotificationRequest = {
    encode(message, writer = new wire_1.BinaryWriter()) {
        if (message.otp !== "") {
            writer.uint32(10).string(message.otp);
        }
        if (message.email !== "") {
            writer.uint32(18).string(message.email);
        }
        if (message.notificationType !== 0) {
            writer.uint32(24).int32(message.notificationType);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof wire_1.BinaryReader ? input : new wire_1.BinaryReader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseSendOtpEmailNotificationRequest();
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
                case 2: {
                    if (tag !== 18) {
                        break;
                    }
                    message.email = reader.string();
                    continue;
                }
                case 3: {
                    if (tag !== 24) {
                        break;
                    }
                    message.notificationType = reader.int32();
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
            otp: isSet(object.otp) ? globalThis.String(object.otp) : "",
            email: isSet(object.email) ? globalThis.String(object.email) : "",
            notificationType: isSet(object.notificationType) ? notificationTypeEnumFromJSON(object.notificationType) : 0,
        };
    },
    toJSON(message) {
        const obj = {};
        if (message.otp !== "") {
            obj.otp = message.otp;
        }
        if (message.email !== "") {
            obj.email = message.email;
        }
        if (message.notificationType !== 0) {
            obj.notificationType = notificationTypeEnumToJSON(message.notificationType);
        }
        return obj;
    },
    create(base) {
        return exports.SendOtpEmailNotificationRequest.fromPartial(base ?? {});
    },
    fromPartial(object) {
        const message = createBaseSendOtpEmailNotificationRequest();
        message.otp = object.otp ?? "";
        message.email = object.email ?? "";
        message.notificationType = object.notificationType ?? 0;
        return message;
    },
};
function isObject(value) {
    return typeof value === "object" && value !== null;
}
function isSet(value) {
    return value !== null && value !== undefined;
}
//# sourceMappingURL=notificationMessage.js.map