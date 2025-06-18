"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.DeleteWishlistItemByIdResponse = exports.DeleteWishlistItemByIdRequest = exports.FindWishlistItemWithPaginationResponse = exports.FindWishlistItemWithPaginationRequest = exports.CreateWishlistItemResponse = exports.CreateWishlistItemRequest = exports.WishlistItem = exports.protobufPackage = void 0;
const wire_1 = require("@bufbuild/protobuf/wire");
exports.protobufPackage = "commerce_v1";
function createBaseWishlistItem() {
    return { id: "", productId: "", userId: "" };
}
exports.WishlistItem = {
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
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof wire_1.BinaryReader ? input : new wire_1.BinaryReader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseWishlistItem();
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
        return obj;
    },
    create(base) {
        return exports.WishlistItem.fromPartial(base ?? {});
    },
    fromPartial(object) {
        const message = createBaseWishlistItem();
        message.id = object.id ?? "";
        message.productId = object.productId ?? "";
        message.userId = object.userId ?? "";
        return message;
    },
};
function createBaseCreateWishlistItemRequest() {
    return { productId: "", userId: "" };
}
exports.CreateWishlistItemRequest = {
    encode(message, writer = new wire_1.BinaryWriter()) {
        if (message.productId !== "") {
            writer.uint32(10).string(message.productId);
        }
        if (message.userId !== "") {
            writer.uint32(18).string(message.userId);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof wire_1.BinaryReader ? input : new wire_1.BinaryReader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseCreateWishlistItemRequest();
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
        return obj;
    },
    create(base) {
        return exports.CreateWishlistItemRequest.fromPartial(base ?? {});
    },
    fromPartial(object) {
        const message = createBaseCreateWishlistItemRequest();
        message.productId = object.productId ?? "";
        message.userId = object.userId ?? "";
        return message;
    },
};
function createBaseCreateWishlistItemResponse() {
    return { id: "" };
}
exports.CreateWishlistItemResponse = {
    encode(message, writer = new wire_1.BinaryWriter()) {
        if (message.id !== "") {
            writer.uint32(10).string(message.id);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof wire_1.BinaryReader ? input : new wire_1.BinaryReader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseCreateWishlistItemResponse();
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
        return exports.CreateWishlistItemResponse.fromPartial(base ?? {});
    },
    fromPartial(object) {
        const message = createBaseCreateWishlistItemResponse();
        message.id = object.id ?? "";
        return message;
    },
};
function createBaseFindWishlistItemWithPaginationRequest() {
    return { userId: "", productIds: [], page: 0, limit: 0 };
}
exports.FindWishlistItemWithPaginationRequest = {
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
        const message = createBaseFindWishlistItemWithPaginationRequest();
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
        return exports.FindWishlistItemWithPaginationRequest.fromPartial(base ?? {});
    },
    fromPartial(object) {
        const message = createBaseFindWishlistItemWithPaginationRequest();
        message.userId = object.userId ?? "";
        message.productIds = object.productIds?.map((e) => e) || [];
        message.page = object.page ?? 0;
        message.limit = object.limit ?? 0;
        return message;
    },
};
function createBaseFindWishlistItemWithPaginationResponse() {
    return { items: [], page: 0, limit: 0 };
}
exports.FindWishlistItemWithPaginationResponse = {
    encode(message, writer = new wire_1.BinaryWriter()) {
        for (const v of message.items) {
            exports.WishlistItem.encode(v, writer.uint32(10).fork()).join();
        }
        if (message.page !== 0) {
            writer.uint32(16).int32(message.page);
        }
        if (message.limit !== 0) {
            writer.uint32(24).int32(message.limit);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof wire_1.BinaryReader ? input : new wire_1.BinaryReader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseFindWishlistItemWithPaginationResponse();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1: {
                    if (tag !== 10) {
                        break;
                    }
                    message.items.push(exports.WishlistItem.decode(reader, reader.uint32()));
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
            items: globalThis.Array.isArray(object?.items) ? object.items.map((e) => exports.WishlistItem.fromJSON(e)) : [],
            page: isSet(object.page) ? globalThis.Number(object.page) : 0,
            limit: isSet(object.limit) ? globalThis.Number(object.limit) : 0,
        };
    },
    toJSON(message) {
        const obj = {};
        if (message.items?.length) {
            obj.items = message.items.map((e) => exports.WishlistItem.toJSON(e));
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
        return exports.FindWishlistItemWithPaginationResponse.fromPartial(base ?? {});
    },
    fromPartial(object) {
        const message = createBaseFindWishlistItemWithPaginationResponse();
        message.items = object.items?.map((e) => exports.WishlistItem.fromPartial(e)) || [];
        message.page = object.page ?? 0;
        message.limit = object.limit ?? 0;
        return message;
    },
};
function createBaseDeleteWishlistItemByIdRequest() {
    return { id: "" };
}
exports.DeleteWishlistItemByIdRequest = {
    encode(message, writer = new wire_1.BinaryWriter()) {
        if (message.id !== "") {
            writer.uint32(10).string(message.id);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof wire_1.BinaryReader ? input : new wire_1.BinaryReader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseDeleteWishlistItemByIdRequest();
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
        return exports.DeleteWishlistItemByIdRequest.fromPartial(base ?? {});
    },
    fromPartial(object) {
        const message = createBaseDeleteWishlistItemByIdRequest();
        message.id = object.id ?? "";
        return message;
    },
};
function createBaseDeleteWishlistItemByIdResponse() {
    return { userId: "" };
}
exports.DeleteWishlistItemByIdResponse = {
    encode(message, writer = new wire_1.BinaryWriter()) {
        if (message.userId !== "") {
            writer.uint32(10).string(message.userId);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof wire_1.BinaryReader ? input : new wire_1.BinaryReader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseDeleteWishlistItemByIdResponse();
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
            }
            if ((tag & 7) === 4 || tag === 0) {
                break;
            }
            reader.skip(tag & 7);
        }
        return message;
    },
    fromJSON(object) {
        return { userId: isSet(object.userId) ? globalThis.String(object.userId) : "" };
    },
    toJSON(message) {
        const obj = {};
        if (message.userId !== "") {
            obj.userId = message.userId;
        }
        return obj;
    },
    create(base) {
        return exports.DeleteWishlistItemByIdResponse.fromPartial(base ?? {});
    },
    fromPartial(object) {
        const message = createBaseDeleteWishlistItemByIdResponse();
        message.userId = object.userId ?? "";
        return message;
    },
};
function isSet(value) {
    return value !== null && value !== undefined;
}
//# sourceMappingURL=wishlistMessage.js.map