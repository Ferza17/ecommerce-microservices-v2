"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.DeleteCartItemByIdResponse = exports.DeleteCartItemByIdRequest = exports.FindCartItemByIdRequest = exports.FindCartItemsWithPaginationResponse = exports.FindCartItemsWithPaginationRequest = exports.UpdateCartItemByIdResponse = exports.UpdateCartItemByIdRequest = exports.CreateCartItemResponse = exports.CreateCartItemRequest = exports.CartItem = exports.protobufPackage = void 0;
const wire_1 = require("@bufbuild/protobuf/wire");
const timestamp_1 = require("../../google/protobuf/timestamp");
exports.protobufPackage = "commerce_v1";
function createBaseCartItem() {
    return { id: "", productId: "", userId: "", qty: 0, price: 0, cratedAt: undefined, updatedAt: undefined };
}
exports.CartItem = {
    encode(message, writer = new wire_1.BinaryWriter()) {
        if (message.id !== "") {
            writer.uint32(10).string(message.id);
        }
        if (message.productId !== "") {
            writer.uint32(18).string(message.productId);
        }
        if (message.userId !== "") {
            writer.uint32(26).string(message.userId);
        }
        if (message.qty !== 0) {
            writer.uint32(32).int32(message.qty);
        }
        if (message.price !== 0) {
            writer.uint32(41).double(message.price);
        }
        if (message.cratedAt !== undefined) {
            timestamp_1.Timestamp.encode(toTimestamp(message.cratedAt), writer.uint32(50).fork()).join();
        }
        if (message.updatedAt !== undefined) {
            timestamp_1.Timestamp.encode(toTimestamp(message.updatedAt), writer.uint32(58).fork()).join();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof wire_1.BinaryReader ? input : new wire_1.BinaryReader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseCartItem();
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
                    if (tag !== 26) {
                        break;
                    }
                    message.userId = reader.string();
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
                    if (tag !== 41) {
                        break;
                    }
                    message.price = reader.double();
                    continue;
                }
                case 6: {
                    if (tag !== 50) {
                        break;
                    }
                    message.cratedAt = fromTimestamp(timestamp_1.Timestamp.decode(reader, reader.uint32()));
                    continue;
                }
                case 7: {
                    if (tag !== 58) {
                        break;
                    }
                    message.updatedAt = fromTimestamp(timestamp_1.Timestamp.decode(reader, reader.uint32()));
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
            userId: isSet(object.userId) ? globalThis.String(object.userId) : "",
            qty: isSet(object.qty) ? globalThis.Number(object.qty) : 0,
            price: isSet(object.price) ? globalThis.Number(object.price) : 0,
            cratedAt: isSet(object.cratedAt) ? fromJsonTimestamp(object.cratedAt) : undefined,
            updatedAt: isSet(object.updatedAt) ? fromJsonTimestamp(object.updatedAt) : undefined,
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
        if (message.userId !== "") {
            obj.userId = message.userId;
        }
        if (message.qty !== 0) {
            obj.qty = Math.round(message.qty);
        }
        if (message.price !== 0) {
            obj.price = message.price;
        }
        if (message.cratedAt !== undefined) {
            obj.cratedAt = message.cratedAt.toISOString();
        }
        if (message.updatedAt !== undefined) {
            obj.updatedAt = message.updatedAt.toISOString();
        }
        return obj;
    },
    create(base) {
        return exports.CartItem.fromPartial(base ?? {});
    },
    fromPartial(object) {
        const message = createBaseCartItem();
        message.id = object.id ?? "";
        message.productId = object.productId ?? "";
        message.userId = object.userId ?? "";
        message.qty = object.qty ?? 0;
        message.price = object.price ?? 0;
        message.cratedAt = object.cratedAt ?? undefined;
        message.updatedAt = object.updatedAt ?? undefined;
        return message;
    },
};
function createBaseCreateCartItemRequest() {
    return { productId: "", userId: "", qty: 0, price: 0 };
}
exports.CreateCartItemRequest = {
    encode(message, writer = new wire_1.BinaryWriter()) {
        if (message.productId !== "") {
            writer.uint32(10).string(message.productId);
        }
        if (message.userId !== "") {
            writer.uint32(18).string(message.userId);
        }
        if (message.qty !== 0) {
            writer.uint32(24).int32(message.qty);
        }
        if (message.price !== 0) {
            writer.uint32(33).double(message.price);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof wire_1.BinaryReader ? input : new wire_1.BinaryReader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseCreateCartItemRequest();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1: {
                    if (tag !== 10) {
                        break;
                    }
                    message.productId = reader.string();
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
                    if (tag !== 24) {
                        break;
                    }
                    message.qty = reader.int32();
                    continue;
                }
                case 4: {
                    if (tag !== 33) {
                        break;
                    }
                    message.price = reader.double();
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
            productId: isSet(object.productId) ? globalThis.String(object.productId) : "",
            userId: isSet(object.userId) ? globalThis.String(object.userId) : "",
            qty: isSet(object.qty) ? globalThis.Number(object.qty) : 0,
            price: isSet(object.price) ? globalThis.Number(object.price) : 0,
        };
    },
    toJSON(message) {
        const obj = {};
        if (message.productId !== "") {
            obj.productId = message.productId;
        }
        if (message.userId !== "") {
            obj.userId = message.userId;
        }
        if (message.qty !== 0) {
            obj.qty = Math.round(message.qty);
        }
        if (message.price !== 0) {
            obj.price = message.price;
        }
        return obj;
    },
    create(base) {
        return exports.CreateCartItemRequest.fromPartial(base ?? {});
    },
    fromPartial(object) {
        const message = createBaseCreateCartItemRequest();
        message.productId = object.productId ?? "";
        message.userId = object.userId ?? "";
        message.qty = object.qty ?? 0;
        message.price = object.price ?? 0;
        return message;
    },
};
function createBaseCreateCartItemResponse() {
    return { id: "" };
}
exports.CreateCartItemResponse = {
    encode(message, writer = new wire_1.BinaryWriter()) {
        if (message.id !== "") {
            writer.uint32(10).string(message.id);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof wire_1.BinaryReader ? input : new wire_1.BinaryReader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseCreateCartItemResponse();
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
        return exports.CreateCartItemResponse.fromPartial(base ?? {});
    },
    fromPartial(object) {
        const message = createBaseCreateCartItemResponse();
        message.id = object.id ?? "";
        return message;
    },
};
function createBaseUpdateCartItemByIdRequest() {
    return { id: "", productId: "", userId: "", qty: 0, price: 0 };
}
exports.UpdateCartItemByIdRequest = {
    encode(message, writer = new wire_1.BinaryWriter()) {
        if (message.id !== "") {
            writer.uint32(10).string(message.id);
        }
        if (message.productId !== "") {
            writer.uint32(18).string(message.productId);
        }
        if (message.userId !== "") {
            writer.uint32(26).string(message.userId);
        }
        if (message.qty !== 0) {
            writer.uint32(32).int32(message.qty);
        }
        if (message.price !== 0) {
            writer.uint32(41).double(message.price);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof wire_1.BinaryReader ? input : new wire_1.BinaryReader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseUpdateCartItemByIdRequest();
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
                    if (tag !== 26) {
                        break;
                    }
                    message.userId = reader.string();
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
                    if (tag !== 41) {
                        break;
                    }
                    message.price = reader.double();
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
            userId: isSet(object.userId) ? globalThis.String(object.userId) : "",
            qty: isSet(object.qty) ? globalThis.Number(object.qty) : 0,
            price: isSet(object.price) ? globalThis.Number(object.price) : 0,
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
        if (message.userId !== "") {
            obj.userId = message.userId;
        }
        if (message.qty !== 0) {
            obj.qty = Math.round(message.qty);
        }
        if (message.price !== 0) {
            obj.price = message.price;
        }
        return obj;
    },
    create(base) {
        return exports.UpdateCartItemByIdRequest.fromPartial(base ?? {});
    },
    fromPartial(object) {
        const message = createBaseUpdateCartItemByIdRequest();
        message.id = object.id ?? "";
        message.productId = object.productId ?? "";
        message.userId = object.userId ?? "";
        message.qty = object.qty ?? 0;
        message.price = object.price ?? 0;
        return message;
    },
};
function createBaseUpdateCartItemByIdResponse() {
    return { id: "" };
}
exports.UpdateCartItemByIdResponse = {
    encode(message, writer = new wire_1.BinaryWriter()) {
        if (message.id !== "") {
            writer.uint32(10).string(message.id);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof wire_1.BinaryReader ? input : new wire_1.BinaryReader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseUpdateCartItemByIdResponse();
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
        return exports.UpdateCartItemByIdResponse.fromPartial(base ?? {});
    },
    fromPartial(object) {
        const message = createBaseUpdateCartItemByIdResponse();
        message.id = object.id ?? "";
        return message;
    },
};
function createBaseFindCartItemsWithPaginationRequest() {
    return { userId: "", productIds: [], page: 0, limit: 0 };
}
exports.FindCartItemsWithPaginationRequest = {
    encode(message, writer = new wire_1.BinaryWriter()) {
        if (message.userId !== "") {
            writer.uint32(10).string(message.userId);
        }
        for (const v of message.productIds) {
            writer.uint32(18).string(v);
        }
        if (message.page !== 0) {
            writer.uint32(24).int32(message.page);
        }
        if (message.limit !== 0) {
            writer.uint32(32).int32(message.limit);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof wire_1.BinaryReader ? input : new wire_1.BinaryReader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseFindCartItemsWithPaginationRequest();
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
                    if (tag !== 18) {
                        break;
                    }
                    message.productIds.push(reader.string());
                    continue;
                }
                case 3: {
                    if (tag !== 24) {
                        break;
                    }
                    message.page = reader.int32();
                    continue;
                }
                case 4: {
                    if (tag !== 32) {
                        break;
                    }
                    message.limit = reader.int32();
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
            productIds: globalThis.Array.isArray(object?.productIds)
                ? object.productIds.map((e) => globalThis.String(e))
                : [],
            page: isSet(object.page) ? globalThis.Number(object.page) : 0,
            limit: isSet(object.limit) ? globalThis.Number(object.limit) : 0,
        };
    },
    toJSON(message) {
        const obj = {};
        if (message.userId !== "") {
            obj.userId = message.userId;
        }
        if (message.productIds?.length) {
            obj.productIds = message.productIds;
        }
        if (message.page !== 0) {
            obj.page = Math.round(message.page);
        }
        if (message.limit !== 0) {
            obj.limit = Math.round(message.limit);
        }
        return obj;
    },
    create(base) {
        return exports.FindCartItemsWithPaginationRequest.fromPartial(base ?? {});
    },
    fromPartial(object) {
        const message = createBaseFindCartItemsWithPaginationRequest();
        message.userId = object.userId ?? "";
        message.productIds = object.productIds?.map((e) => e) || [];
        message.page = object.page ?? 0;
        message.limit = object.limit ?? 0;
        return message;
    },
};
function createBaseFindCartItemsWithPaginationResponse() {
    return { items: [], page: 0, limit: 0, total: 0 };
}
exports.FindCartItemsWithPaginationResponse = {
    encode(message, writer = new wire_1.BinaryWriter()) {
        for (const v of message.items) {
            exports.CartItem.encode(v, writer.uint32(10).fork()).join();
        }
        if (message.page !== 0) {
            writer.uint32(16).int32(message.page);
        }
        if (message.limit !== 0) {
            writer.uint32(24).int32(message.limit);
        }
        if (message.total !== 0) {
            writer.uint32(32).int32(message.total);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof wire_1.BinaryReader ? input : new wire_1.BinaryReader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseFindCartItemsWithPaginationResponse();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1: {
                    if (tag !== 10) {
                        break;
                    }
                    message.items.push(exports.CartItem.decode(reader, reader.uint32()));
                    continue;
                }
                case 2: {
                    if (tag !== 16) {
                        break;
                    }
                    message.page = reader.int32();
                    continue;
                }
                case 3: {
                    if (tag !== 24) {
                        break;
                    }
                    message.limit = reader.int32();
                    continue;
                }
                case 4: {
                    if (tag !== 32) {
                        break;
                    }
                    message.total = reader.int32();
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
            items: globalThis.Array.isArray(object?.items) ? object.items.map((e) => exports.CartItem.fromJSON(e)) : [],
            page: isSet(object.page) ? globalThis.Number(object.page) : 0,
            limit: isSet(object.limit) ? globalThis.Number(object.limit) : 0,
            total: isSet(object.total) ? globalThis.Number(object.total) : 0,
        };
    },
    toJSON(message) {
        const obj = {};
        if (message.items?.length) {
            obj.items = message.items.map((e) => exports.CartItem.toJSON(e));
        }
        if (message.page !== 0) {
            obj.page = Math.round(message.page);
        }
        if (message.limit !== 0) {
            obj.limit = Math.round(message.limit);
        }
        if (message.total !== 0) {
            obj.total = Math.round(message.total);
        }
        return obj;
    },
    create(base) {
        return exports.FindCartItemsWithPaginationResponse.fromPartial(base ?? {});
    },
    fromPartial(object) {
        const message = createBaseFindCartItemsWithPaginationResponse();
        message.items = object.items?.map((e) => exports.CartItem.fromPartial(e)) || [];
        message.page = object.page ?? 0;
        message.limit = object.limit ?? 0;
        message.total = object.total ?? 0;
        return message;
    },
};
function createBaseFindCartItemByIdRequest() {
    return { id: "" };
}
exports.FindCartItemByIdRequest = {
    encode(message, writer = new wire_1.BinaryWriter()) {
        if (message.id !== "") {
            writer.uint32(10).string(message.id);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof wire_1.BinaryReader ? input : new wire_1.BinaryReader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseFindCartItemByIdRequest();
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
        return exports.FindCartItemByIdRequest.fromPartial(base ?? {});
    },
    fromPartial(object) {
        const message = createBaseFindCartItemByIdRequest();
        message.id = object.id ?? "";
        return message;
    },
};
function createBaseDeleteCartItemByIdRequest() {
    return { id: "" };
}
exports.DeleteCartItemByIdRequest = {
    encode(message, writer = new wire_1.BinaryWriter()) {
        if (message.id !== "") {
            writer.uint32(10).string(message.id);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof wire_1.BinaryReader ? input : new wire_1.BinaryReader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseDeleteCartItemByIdRequest();
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
        return exports.DeleteCartItemByIdRequest.fromPartial(base ?? {});
    },
    fromPartial(object) {
        const message = createBaseDeleteCartItemByIdRequest();
        message.id = object.id ?? "";
        return message;
    },
};
function createBaseDeleteCartItemByIdResponse() {
    return { message: "" };
}
exports.DeleteCartItemByIdResponse = {
    encode(message, writer = new wire_1.BinaryWriter()) {
        if (message.message !== "") {
            writer.uint32(10).string(message.message);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof wire_1.BinaryReader ? input : new wire_1.BinaryReader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseDeleteCartItemByIdResponse();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1: {
                    if (tag !== 10) {
                        break;
                    }
                    message.message = reader.string();
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
        return { message: isSet(object.message) ? globalThis.String(object.message) : "" };
    },
    toJSON(message) {
        const obj = {};
        if (message.message !== "") {
            obj.message = message.message;
        }
        return obj;
    },
    create(base) {
        return exports.DeleteCartItemByIdResponse.fromPartial(base ?? {});
    },
    fromPartial(object) {
        const message = createBaseDeleteCartItemByIdResponse();
        message.message = object.message ?? "";
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
//# sourceMappingURL=cartMessage.js.map