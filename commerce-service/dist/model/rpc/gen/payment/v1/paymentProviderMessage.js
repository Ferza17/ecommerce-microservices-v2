"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.FindPaymentProviderByIdRequest = exports.FindPaymentProvidersResponse = exports.FindPaymentProvidersRequest = exports.Provider = exports.ProviderMethod = exports.protobufPackage = void 0;
exports.providerMethodFromJSON = providerMethodFromJSON;
exports.providerMethodToJSON = providerMethodToJSON;
const wire_1 = require("@bufbuild/protobuf/wire");
const timestamp_1 = require("../../google/protobuf/timestamp");
exports.protobufPackage = "payment_v1";
var ProviderMethod;
(function (ProviderMethod) {
    ProviderMethod[ProviderMethod["BANK"] = 0] = "BANK";
    ProviderMethod[ProviderMethod["CRYPTO_CURRENCY"] = 1] = "CRYPTO_CURRENCY";
    ProviderMethod[ProviderMethod["DEBIT"] = 2] = "DEBIT";
    ProviderMethod[ProviderMethod["CREDIT"] = 3] = "CREDIT";
    ProviderMethod[ProviderMethod["CASH_ON_DELIVERY"] = 4] = "CASH_ON_DELIVERY";
    ProviderMethod[ProviderMethod["UNRECOGNIZED"] = -1] = "UNRECOGNIZED";
})(ProviderMethod || (exports.ProviderMethod = ProviderMethod = {}));
function providerMethodFromJSON(object) {
    switch (object) {
        case 0:
        case "BANK":
            return ProviderMethod.BANK;
        case 1:
        case "CRYPTO_CURRENCY":
            return ProviderMethod.CRYPTO_CURRENCY;
        case 2:
        case "DEBIT":
            return ProviderMethod.DEBIT;
        case 3:
        case "CREDIT":
            return ProviderMethod.CREDIT;
        case 4:
        case "CASH_ON_DELIVERY":
            return ProviderMethod.CASH_ON_DELIVERY;
        case -1:
        case "UNRECOGNIZED":
        default:
            return ProviderMethod.UNRECOGNIZED;
    }
}
function providerMethodToJSON(object) {
    switch (object) {
        case ProviderMethod.BANK:
            return "BANK";
        case ProviderMethod.CRYPTO_CURRENCY:
            return "CRYPTO_CURRENCY";
        case ProviderMethod.DEBIT:
            return "DEBIT";
        case ProviderMethod.CREDIT:
            return "CREDIT";
        case ProviderMethod.CASH_ON_DELIVERY:
            return "CASH_ON_DELIVERY";
        case ProviderMethod.UNRECOGNIZED:
        default:
            return "UNRECOGNIZED";
    }
}
function createBaseProvider() {
    return { id: "", name: "", method: 0, createdAt: undefined, updatedAt: undefined, discardedAt: undefined };
}
exports.Provider = {
    encode(message, writer = new wire_1.BinaryWriter()) {
        if (message.id !== "") {
            writer.uint32(10).string(message.id);
        }
        if (message.name !== "") {
            writer.uint32(18).string(message.name);
        }
        if (message.method !== 0) {
            writer.uint32(24).int32(message.method);
        }
        if (message.createdAt !== undefined) {
            timestamp_1.Timestamp.encode(toTimestamp(message.createdAt), writer.uint32(34).fork()).join();
        }
        if (message.updatedAt !== undefined) {
            timestamp_1.Timestamp.encode(toTimestamp(message.updatedAt), writer.uint32(42).fork()).join();
        }
        if (message.discardedAt !== undefined) {
            timestamp_1.Timestamp.encode(toTimestamp(message.discardedAt), writer.uint32(50).fork()).join();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof wire_1.BinaryReader ? input : new wire_1.BinaryReader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseProvider();
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
                    message.name = reader.string();
                    continue;
                }
                case 3: {
                    if (tag !== 24) {
                        break;
                    }
                    message.method = reader.int32();
                    continue;
                }
                case 4: {
                    if (tag !== 34) {
                        break;
                    }
                    message.createdAt = fromTimestamp(timestamp_1.Timestamp.decode(reader, reader.uint32()));
                    continue;
                }
                case 5: {
                    if (tag !== 42) {
                        break;
                    }
                    message.updatedAt = fromTimestamp(timestamp_1.Timestamp.decode(reader, reader.uint32()));
                    continue;
                }
                case 6: {
                    if (tag !== 50) {
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
            name: isSet(object.name) ? globalThis.String(object.name) : "",
            method: isSet(object.method) ? providerMethodFromJSON(object.method) : 0,
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
        if (message.name !== "") {
            obj.name = message.name;
        }
        if (message.method !== 0) {
            obj.method = providerMethodToJSON(message.method);
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
        return exports.Provider.fromPartial(base ?? {});
    },
    fromPartial(object) {
        const message = createBaseProvider();
        message.id = object.id ?? "";
        message.name = object.name ?? "";
        message.method = object.method ?? 0;
        message.createdAt = object.createdAt ?? undefined;
        message.updatedAt = object.updatedAt ?? undefined;
        message.discardedAt = object.discardedAt ?? undefined;
        return message;
    },
};
function createBaseFindPaymentProvidersRequest() {
    return { name: "" };
}
exports.FindPaymentProvidersRequest = {
    encode(message, writer = new wire_1.BinaryWriter()) {
        if (message.name !== "") {
            writer.uint32(10).string(message.name);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof wire_1.BinaryReader ? input : new wire_1.BinaryReader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseFindPaymentProvidersRequest();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1: {
                    if (tag !== 10) {
                        break;
                    }
                    message.name = reader.string();
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
        return { name: isSet(object.name) ? globalThis.String(object.name) : "" };
    },
    toJSON(message) {
        const obj = {};
        if (message.name !== "") {
            obj.name = message.name;
        }
        return obj;
    },
    create(base) {
        return exports.FindPaymentProvidersRequest.fromPartial(base ?? {});
    },
    fromPartial(object) {
        const message = createBaseFindPaymentProvidersRequest();
        message.name = object.name ?? "";
        return message;
    },
};
function createBaseFindPaymentProvidersResponse() {
    return { data: [] };
}
exports.FindPaymentProvidersResponse = {
    encode(message, writer = new wire_1.BinaryWriter()) {
        for (const v of message.data) {
            exports.Provider.encode(v, writer.uint32(10).fork()).join();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof wire_1.BinaryReader ? input : new wire_1.BinaryReader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseFindPaymentProvidersResponse();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1: {
                    if (tag !== 10) {
                        break;
                    }
                    message.data.push(exports.Provider.decode(reader, reader.uint32()));
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
        return { data: globalThis.Array.isArray(object?.data) ? object.data.map((e) => exports.Provider.fromJSON(e)) : [] };
    },
    toJSON(message) {
        const obj = {};
        if (message.data?.length) {
            obj.data = message.data.map((e) => exports.Provider.toJSON(e));
        }
        return obj;
    },
    create(base) {
        return exports.FindPaymentProvidersResponse.fromPartial(base ?? {});
    },
    fromPartial(object) {
        const message = createBaseFindPaymentProvidersResponse();
        message.data = object.data?.map((e) => exports.Provider.fromPartial(e)) || [];
        return message;
    },
};
function createBaseFindPaymentProviderByIdRequest() {
    return { id: "" };
}
exports.FindPaymentProviderByIdRequest = {
    encode(message, writer = new wire_1.BinaryWriter()) {
        if (message.id !== "") {
            writer.uint32(10).string(message.id);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof wire_1.BinaryReader ? input : new wire_1.BinaryReader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseFindPaymentProviderByIdRequest();
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
        return exports.FindPaymentProviderByIdRequest.fromPartial(base ?? {});
    },
    fromPartial(object) {
        const message = createBaseFindPaymentProviderByIdRequest();
        message.id = object.id ?? "";
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
//# sourceMappingURL=paymentProviderMessage.js.map