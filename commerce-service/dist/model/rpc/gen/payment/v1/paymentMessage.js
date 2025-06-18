"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.FindPaymentByUserIdAndStatusRequest = exports.FindPaymentByIdRequest = exports.CallBackPaymentRequest = exports.PaymentOrderDelayedCancelledRequest = exports.CreatePaymentRequest = exports.Payment = exports.PaymentItem = exports.PaymentStatus = exports.protobufPackage = void 0;
exports.paymentStatusFromJSON = paymentStatusFromJSON;
exports.paymentStatusToJSON = paymentStatusToJSON;
const wire_1 = require("@bufbuild/protobuf/wire");
const timestamp_1 = require("../../google/protobuf/timestamp");
const paymentProviderMessage_1 = require("./paymentProviderMessage");
exports.protobufPackage = "payment_v1";
var PaymentStatus;
(function (PaymentStatus) {
    PaymentStatus[PaymentStatus["PENDING"] = 0] = "PENDING";
    PaymentStatus[PaymentStatus["PARTIAL"] = 1] = "PARTIAL";
    PaymentStatus[PaymentStatus["SUCCESS"] = 2] = "SUCCESS";
    PaymentStatus[PaymentStatus["FAILED"] = 3] = "FAILED";
    PaymentStatus[PaymentStatus["UNRECOGNIZED"] = -1] = "UNRECOGNIZED";
})(PaymentStatus || (exports.PaymentStatus = PaymentStatus = {}));
function paymentStatusFromJSON(object) {
    switch (object) {
        case 0:
        case "PENDING":
            return PaymentStatus.PENDING;
        case 1:
        case "PARTIAL":
            return PaymentStatus.PARTIAL;
        case 2:
        case "SUCCESS":
            return PaymentStatus.SUCCESS;
        case 3:
        case "FAILED":
            return PaymentStatus.FAILED;
        case -1:
        case "UNRECOGNIZED":
        default:
            return PaymentStatus.UNRECOGNIZED;
    }
}
function paymentStatusToJSON(object) {
    switch (object) {
        case PaymentStatus.PENDING:
            return "PENDING";
        case PaymentStatus.PARTIAL:
            return "PARTIAL";
        case PaymentStatus.SUCCESS:
            return "SUCCESS";
        case PaymentStatus.FAILED:
            return "FAILED";
        case PaymentStatus.UNRECOGNIZED:
        default:
            return "UNRECOGNIZED";
    }
}
function createBasePaymentItem() {
    return {
        id: "",
        productId: "",
        amount: 0,
        qty: 0,
        cratedAt: undefined,
        updatedAt: undefined,
        discardedAt: undefined,
    };
}
exports.PaymentItem = {
    encode(message, writer = new wire_1.BinaryWriter()) {
        if (message.id !== "") {
            writer.uint32(10).string(message.id);
        }
        if (message.productId !== "") {
            writer.uint32(18).string(message.productId);
        }
        if (message.amount !== 0) {
            writer.uint32(25).double(message.amount);
        }
        if (message.qty !== 0) {
            writer.uint32(32).int32(message.qty);
        }
        if (message.cratedAt !== undefined) {
            timestamp_1.Timestamp.encode(toTimestamp(message.cratedAt), writer.uint32(42).fork()).join();
        }
        if (message.updatedAt !== undefined) {
            timestamp_1.Timestamp.encode(toTimestamp(message.updatedAt), writer.uint32(50).fork()).join();
        }
        if (message.discardedAt !== undefined) {
            timestamp_1.Timestamp.encode(toTimestamp(message.discardedAt), writer.uint32(58).fork()).join();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof wire_1.BinaryReader ? input : new wire_1.BinaryReader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBasePaymentItem();
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
                    message.productId = reader.string();
                    continue;
                }
                case 3: {
                    if (tag !== 25) {
                        break;
                    }
                    message.amount = reader.double();
                    continue;
                }
                case 4: {
                    if (tag !== 32) {
                        break;
                    }
                    message.qty = reader.int32();
                    continue;
                }
                case 5: {
                    if (tag !== 42) {
                        break;
                    }
                    message.cratedAt = fromTimestamp(timestamp_1.Timestamp.decode(reader, reader.uint32()));
                    continue;
                }
                case 6: {
                    if (tag !== 50) {
                        break;
                    }
                    message.updatedAt = fromTimestamp(timestamp_1.Timestamp.decode(reader, reader.uint32()));
                    continue;
                }
                case 7: {
                    if (tag !== 58) {
                        break;
                    }
                    message.discardedAt = fromTimestamp(timestamp_1.Timestamp.decode(reader, reader.uint32()));
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
            productId: isSet(object.productId) ? globalThis.String(object.productId) : "",
            amount: isSet(object.amount) ? globalThis.Number(object.amount) : 0,
            qty: isSet(object.qty) ? globalThis.Number(object.qty) : 0,
            cratedAt: isSet(object.cratedAt) ? fromJsonTimestamp(object.cratedAt) : undefined,
            updatedAt: isSet(object.updatedAt) ? fromJsonTimestamp(object.updatedAt) : undefined,
            discardedAt: isSet(object.discardedAt) ? fromJsonTimestamp(object.discardedAt) : undefined,
        };
    },
    toJSON(message) {
        const obj = {};
        if (message.id !== "") {
            obj.id = message.id;
        }
        if (message.productId !== "") {
            obj.productId = message.productId;
        }
        if (message.amount !== 0) {
            obj.amount = message.amount;
        }
        if (message.qty !== 0) {
            obj.qty = Math.round(message.qty);
        }
        if (message.cratedAt !== undefined) {
            obj.cratedAt = message.cratedAt.toISOString();
        }
        if (message.updatedAt !== undefined) {
            obj.updatedAt = message.updatedAt.toISOString();
        }
        if (message.discardedAt !== undefined) {
            obj.discardedAt = message.discardedAt.toISOString();
        }
        return obj;
    },
    create(base) {
        return exports.PaymentItem.fromPartial(base ?? {});
    },
    fromPartial(object) {
        const message = createBasePaymentItem();
        message.id = object.id ?? "";
        message.productId = object.productId ?? "";
        message.amount = object.amount ?? 0;
        message.qty = object.qty ?? 0;
        message.cratedAt = object.cratedAt ?? undefined;
        message.updatedAt = object.updatedAt ?? undefined;
        message.discardedAt = object.discardedAt ?? undefined;
        return message;
    },
};
function createBasePayment() {
    return {
        id: "",
        code: "",
        Items: [],
        totalPrice: 0,
        status: 0,
        provider: undefined,
        userId: "",
        createdAt: undefined,
        updatedAt: undefined,
        discardedAt: undefined,
    };
}
exports.Payment = {
    encode(message, writer = new wire_1.BinaryWriter()) {
        if (message.id !== "") {
            writer.uint32(10).string(message.id);
        }
        if (message.code !== "") {
            writer.uint32(18).string(message.code);
        }
        for (const v of message.Items) {
            exports.PaymentItem.encode(v, writer.uint32(26).fork()).join();
        }
        if (message.totalPrice !== 0) {
            writer.uint32(33).double(message.totalPrice);
        }
        if (message.status !== 0) {
            writer.uint32(40).int32(message.status);
        }
        if (message.provider !== undefined) {
            paymentProviderMessage_1.Provider.encode(message.provider, writer.uint32(50).fork()).join();
        }
        if (message.userId !== "") {
            writer.uint32(58).string(message.userId);
        }
        if (message.createdAt !== undefined) {
            timestamp_1.Timestamp.encode(toTimestamp(message.createdAt), writer.uint32(66).fork()).join();
        }
        if (message.updatedAt !== undefined) {
            timestamp_1.Timestamp.encode(toTimestamp(message.updatedAt), writer.uint32(74).fork()).join();
        }
        if (message.discardedAt !== undefined) {
            timestamp_1.Timestamp.encode(toTimestamp(message.discardedAt), writer.uint32(82).fork()).join();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof wire_1.BinaryReader ? input : new wire_1.BinaryReader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBasePayment();
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
                    message.code = reader.string();
                    continue;
                }
                case 3: {
                    if (tag !== 26) {
                        break;
                    }
                    message.Items.push(exports.PaymentItem.decode(reader, reader.uint32()));
                    continue;
                }
                case 4: {
                    if (tag !== 33) {
                        break;
                    }
                    message.totalPrice = reader.double();
                    continue;
                }
                case 5: {
                    if (tag !== 40) {
                        break;
                    }
                    message.status = reader.int32();
                    continue;
                }
                case 6: {
                    if (tag !== 50) {
                        break;
                    }
                    message.provider = paymentProviderMessage_1.Provider.decode(reader, reader.uint32());
                    continue;
                }
                case 7: {
                    if (tag !== 58) {
                        break;
                    }
                    message.userId = reader.string();
                    continue;
                }
                case 8: {
                    if (tag !== 66) {
                        break;
                    }
                    message.createdAt = fromTimestamp(timestamp_1.Timestamp.decode(reader, reader.uint32()));
                    continue;
                }
                case 9: {
                    if (tag !== 74) {
                        break;
                    }
                    message.updatedAt = fromTimestamp(timestamp_1.Timestamp.decode(reader, reader.uint32()));
                    continue;
                }
                case 10: {
                    if (tag !== 82) {
                        break;
                    }
                    message.discardedAt = fromTimestamp(timestamp_1.Timestamp.decode(reader, reader.uint32()));
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
            code: isSet(object.code) ? globalThis.String(object.code) : "",
            Items: globalThis.Array.isArray(object?.Items) ? object.Items.map((e) => exports.PaymentItem.fromJSON(e)) : [],
            totalPrice: isSet(object.totalPrice) ? globalThis.Number(object.totalPrice) : 0,
            status: isSet(object.status) ? paymentStatusFromJSON(object.status) : 0,
            provider: isSet(object.provider) ? paymentProviderMessage_1.Provider.fromJSON(object.provider) : undefined,
            userId: isSet(object.userId) ? globalThis.String(object.userId) : "",
            createdAt: isSet(object.createdAt) ? fromJsonTimestamp(object.createdAt) : undefined,
            updatedAt: isSet(object.updatedAt) ? fromJsonTimestamp(object.updatedAt) : undefined,
            discardedAt: isSet(object.discardedAt) ? fromJsonTimestamp(object.discardedAt) : undefined,
        };
    },
    toJSON(message) {
        const obj = {};
        if (message.id !== "") {
            obj.id = message.id;
        }
        if (message.code !== "") {
            obj.code = message.code;
        }
        if (message.Items?.length) {
            obj.Items = message.Items.map((e) => exports.PaymentItem.toJSON(e));
        }
        if (message.totalPrice !== 0) {
            obj.totalPrice = message.totalPrice;
        }
        if (message.status !== 0) {
            obj.status = paymentStatusToJSON(message.status);
        }
        if (message.provider !== undefined) {
            obj.provider = paymentProviderMessage_1.Provider.toJSON(message.provider);
        }
        if (message.userId !== "") {
            obj.userId = message.userId;
        }
        if (message.createdAt !== undefined) {
            obj.createdAt = message.createdAt.toISOString();
        }
        if (message.updatedAt !== undefined) {
            obj.updatedAt = message.updatedAt.toISOString();
        }
        if (message.discardedAt !== undefined) {
            obj.discardedAt = message.discardedAt.toISOString();
        }
        return obj;
    },
    create(base) {
        return exports.Payment.fromPartial(base ?? {});
    },
    fromPartial(object) {
        const message = createBasePayment();
        message.id = object.id ?? "";
        message.code = object.code ?? "";
        message.Items = object.Items?.map((e) => exports.PaymentItem.fromPartial(e)) || [];
        message.totalPrice = object.totalPrice ?? 0;
        message.status = object.status ?? 0;
        message.provider = (object.provider !== undefined && object.provider !== null)
            ? paymentProviderMessage_1.Provider.fromPartial(object.provider)
            : undefined;
        message.userId = object.userId ?? "";
        message.createdAt = object.createdAt ?? undefined;
        message.updatedAt = object.updatedAt ?? undefined;
        message.discardedAt = object.discardedAt ?? undefined;
        return message;
    },
};
function createBaseCreatePaymentRequest() {
    return { items: [], userId: "", amount: 0, providerId: "" };
}
exports.CreatePaymentRequest = {
    encode(message, writer = new wire_1.BinaryWriter()) {
        for (const v of message.items) {
            exports.PaymentItem.encode(v, writer.uint32(10).fork()).join();
        }
        if (message.userId !== "") {
            writer.uint32(18).string(message.userId);
        }
        if (message.amount !== 0) {
            writer.uint32(25).double(message.amount);
        }
        if (message.providerId !== "") {
            writer.uint32(34).string(message.providerId);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof wire_1.BinaryReader ? input : new wire_1.BinaryReader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseCreatePaymentRequest();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1: {
                    if (tag !== 10) {
                        break;
                    }
                    message.items.push(exports.PaymentItem.decode(reader, reader.uint32()));
                    continue;
                }
                case 2: {
                    if (tag !== 18) {
                        break;
                    }
                    message.userId = reader.string();
                    continue;
                }
                case 3: {
                    if (tag !== 25) {
                        break;
                    }
                    message.amount = reader.double();
                    continue;
                }
                case 4: {
                    if (tag !== 34) {
                        break;
                    }
                    message.providerId = reader.string();
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
            items: globalThis.Array.isArray(object?.items) ? object.items.map((e) => exports.PaymentItem.fromJSON(e)) : [],
            userId: isSet(object.userId) ? globalThis.String(object.userId) : "",
            amount: isSet(object.amount) ? globalThis.Number(object.amount) : 0,
            providerId: isSet(object.providerId) ? globalThis.String(object.providerId) : "",
        };
    },
    toJSON(message) {
        const obj = {};
        if (message.items?.length) {
            obj.items = message.items.map((e) => exports.PaymentItem.toJSON(e));
        }
        if (message.userId !== "") {
            obj.userId = message.userId;
        }
        if (message.amount !== 0) {
            obj.amount = message.amount;
        }
        if (message.providerId !== "") {
            obj.providerId = message.providerId;
        }
        return obj;
    },
    create(base) {
        return exports.CreatePaymentRequest.fromPartial(base ?? {});
    },
    fromPartial(object) {
        const message = createBaseCreatePaymentRequest();
        message.items = object.items?.map((e) => exports.PaymentItem.fromPartial(e)) || [];
        message.userId = object.userId ?? "";
        message.amount = object.amount ?? 0;
        message.providerId = object.providerId ?? "";
        return message;
    },
};
function createBasePaymentOrderDelayedCancelledRequest() {
    return { id: "" };
}
exports.PaymentOrderDelayedCancelledRequest = {
    encode(message, writer = new wire_1.BinaryWriter()) {
        if (message.id !== "") {
            writer.uint32(10).string(message.id);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof wire_1.BinaryReader ? input : new wire_1.BinaryReader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBasePaymentOrderDelayedCancelledRequest();
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
            }
            if ((tag & 7) === 4 || tag === 0) {
                break;
            }
            reader.skip(tag & 7);
        }
        return message;
    },
    fromJSON(object) {
        return { id: isSet(object.id) ? globalThis.String(object.id) : "" };
    },
    toJSON(message) {
        const obj = {};
        if (message.id !== "") {
            obj.id = message.id;
        }
        return obj;
    },
    create(base) {
        return exports.PaymentOrderDelayedCancelledRequest.fromPartial(base ?? {});
    },
    fromPartial(object) {
        const message = createBasePaymentOrderDelayedCancelledRequest();
        message.id = object.id ?? "";
        return message;
    },
};
function createBaseCallBackPaymentRequest() {
    return { paymentId: "", amount: 0 };
}
exports.CallBackPaymentRequest = {
    encode(message, writer = new wire_1.BinaryWriter()) {
        if (message.paymentId !== "") {
            writer.uint32(10).string(message.paymentId);
        }
        if (message.amount !== 0) {
            writer.uint32(17).double(message.amount);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof wire_1.BinaryReader ? input : new wire_1.BinaryReader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseCallBackPaymentRequest();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1: {
                    if (tag !== 10) {
                        break;
                    }
                    message.paymentId = reader.string();
                    continue;
                }
                case 2: {
                    if (tag !== 17) {
                        break;
                    }
                    message.amount = reader.double();
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
            paymentId: isSet(object.paymentId) ? globalThis.String(object.paymentId) : "",
            amount: isSet(object.amount) ? globalThis.Number(object.amount) : 0,
        };
    },
    toJSON(message) {
        const obj = {};
        if (message.paymentId !== "") {
            obj.paymentId = message.paymentId;
        }
        if (message.amount !== 0) {
            obj.amount = message.amount;
        }
        return obj;
    },
    create(base) {
        return exports.CallBackPaymentRequest.fromPartial(base ?? {});
    },
    fromPartial(object) {
        const message = createBaseCallBackPaymentRequest();
        message.paymentId = object.paymentId ?? "";
        message.amount = object.amount ?? 0;
        return message;
    },
};
function createBaseFindPaymentByIdRequest() {
    return { id: "" };
}
exports.FindPaymentByIdRequest = {
    encode(message, writer = new wire_1.BinaryWriter()) {
        if (message.id !== "") {
            writer.uint32(10).string(message.id);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof wire_1.BinaryReader ? input : new wire_1.BinaryReader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseFindPaymentByIdRequest();
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
            }
            if ((tag & 7) === 4 || tag === 0) {
                break;
            }
            reader.skip(tag & 7);
        }
        return message;
    },
    fromJSON(object) {
        return { id: isSet(object.id) ? globalThis.String(object.id) : "" };
    },
    toJSON(message) {
        const obj = {};
        if (message.id !== "") {
            obj.id = message.id;
        }
        return obj;
    },
    create(base) {
        return exports.FindPaymentByIdRequest.fromPartial(base ?? {});
    },
    fromPartial(object) {
        const message = createBaseFindPaymentByIdRequest();
        message.id = object.id ?? "";
        return message;
    },
};
function createBaseFindPaymentByUserIdAndStatusRequest() {
    return { userId: "", status: 0 };
}
exports.FindPaymentByUserIdAndStatusRequest = {
    encode(message, writer = new wire_1.BinaryWriter()) {
        if (message.userId !== "") {
            writer.uint32(10).string(message.userId);
        }
        if (message.status !== 0) {
            writer.uint32(16).int32(message.status);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof wire_1.BinaryReader ? input : new wire_1.BinaryReader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseFindPaymentByUserIdAndStatusRequest();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1: {
                    if (tag !== 10) {
                        break;
                    }
                    message.userId = reader.string();
                    continue;
                }
                case 2: {
                    if (tag !== 16) {
                        break;
                    }
                    message.status = reader.int32();
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
            userId: isSet(object.userId) ? globalThis.String(object.userId) : "",
            status: isSet(object.status) ? paymentStatusFromJSON(object.status) : 0,
        };
    },
    toJSON(message) {
        const obj = {};
        if (message.userId !== "") {
            obj.userId = message.userId;
        }
        if (message.status !== 0) {
            obj.status = paymentStatusToJSON(message.status);
        }
        return obj;
    },
    create(base) {
        return exports.FindPaymentByUserIdAndStatusRequest.fromPartial(base ?? {});
    },
    fromPartial(object) {
        const message = createBaseFindPaymentByUserIdAndStatusRequest();
        message.userId = object.userId ?? "";
        message.status = object.status ?? 0;
        return message;
    },
};
function toTimestamp(date) {
    const seconds = Math.trunc(date.getTime() / 1_000);
    const nanos = (date.getTime() % 1_000) * 1_000_000;
    return { seconds, nanos };
}
function fromTimestamp(t) {
    let millis = (t.seconds || 0) * 1_000;
    millis += (t.nanos || 0) / 1_000_000;
    return new globalThis.Date(millis);
}
function fromJsonTimestamp(o) {
    if (o instanceof globalThis.Date) {
        return o;
    }
    else if (typeof o === "string") {
        return new globalThis.Date(o);
    }
    else {
        return fromTimestamp(timestamp_1.Timestamp.fromJSON(o));
    }
}
function isSet(value) {
    return value !== null && value !== undefined;
}
//# sourceMappingURL=paymentMessage.js.map