"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.DeleteProductByIdResponse = exports.DeleteProductByIdRequest = exports.UpdateProductByIdRequest = exports.CreateProductResponse = exports.CreateProductRequest = exports.FindProductByIdRequest = exports.FindProductsWithPaginationResponse = exports.FindProductsWithPaginationRequest = exports.Product = exports.protobufPackage = void 0;
const wire_1 = require("@bufbuild/protobuf/wire");
const timestamp_1 = require("../../google/protobuf/timestamp");
exports.protobufPackage = "product_v1";
function createBaseProduct() {
    return {
        id: "",
        name: "",
        description: "",
        uom: "",
        image: "",
        price: 0,
        stock: 0,
        createdAt: undefined,
        updatedAt: undefined,
        discardedAt: undefined,
    };
}
exports.Product = {
    encode(message, writer = new wire_1.BinaryWriter()) {
        if (message.id !== "") {
            writer.uint32(10).string(message.id);
        }
        if (message.name !== "") {
            writer.uint32(18).string(message.name);
        }
        if (message.description !== "") {
            writer.uint32(26).string(message.description);
        }
        if (message.uom !== "") {
            writer.uint32(34).string(message.uom);
        }
        if (message.image !== "") {
            writer.uint32(42).string(message.image);
        }
        if (message.price !== 0) {
            writer.uint32(49).double(message.price);
        }
        if (message.stock !== 0) {
            writer.uint32(56).int64(message.stock);
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
        const message = createBaseProduct();
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
                    if (tag !== 26) {
                        break;
                    }
                    message.description = reader.string();
                    continue;
                }
                case 4: {
                    if (tag !== 34) {
                        break;
                    }
                    message.uom = reader.string();
                    continue;
                }
                case 5: {
                    if (tag !== 42) {
                        break;
                    }
                    message.image = reader.string();
                    continue;
                }
                case 6: {
                    if (tag !== 49) {
                        break;
                    }
                    message.price = reader.double();
                    continue;
                }
                case 7: {
                    if (tag !== 56) {
                        break;
                    }
                    message.stock = longToNumber(reader.int64());
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
            name: isSet(object.name) ? globalThis.String(object.name) : "",
            description: isSet(object.description) ? globalThis.String(object.description) : "",
            uom: isSet(object.uom) ? globalThis.String(object.uom) : "",
            image: isSet(object.image) ? globalThis.String(object.image) : "",
            price: isSet(object.price) ? globalThis.Number(object.price) : 0,
            stock: isSet(object.stock) ? globalThis.Number(object.stock) : 0,
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
        if (message.description !== "") {
            obj.description = message.description;
        }
        if (message.uom !== "") {
            obj.uom = message.uom;
        }
        if (message.image !== "") {
            obj.image = message.image;
        }
        if (message.price !== 0) {
            obj.price = message.price;
        }
        if (message.stock !== 0) {
            obj.stock = Math.round(message.stock);
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
        return exports.Product.fromPartial(base ?? {});
    },
    fromPartial(object) {
        const message = createBaseProduct();
        message.id = object.id ?? "";
        message.name = object.name ?? "";
        message.description = object.description ?? "";
        message.uom = object.uom ?? "";
        message.image = object.image ?? "";
        message.price = object.price ?? 0;
        message.stock = object.stock ?? 0;
        message.createdAt = object.createdAt ?? undefined;
        message.updatedAt = object.updatedAt ?? undefined;
        message.discardedAt = object.discardedAt ?? undefined;
        return message;
    },
};
function createBaseFindProductsWithPaginationRequest() {
    return { ids: [], name: [], page: 0, limit: 0 };
}
exports.FindProductsWithPaginationRequest = {
    encode(message, writer = new wire_1.BinaryWriter()) {
        for (const v of message.ids) {
            writer.uint32(10).string(v);
        }
        for (const v of message.name) {
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
        const message = createBaseFindProductsWithPaginationRequest();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1: {
                    if (tag !== 10) {
                        break;
                    }
                    message.ids.push(reader.string());
                    continue;
                }
                case 2: {
                    if (tag !== 18) {
                        break;
                    }
                    message.name.push(reader.string());
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
            ids: globalThis.Array.isArray(object?.ids) ? object.ids.map((e) => globalThis.String(e)) : [],
            name: globalThis.Array.isArray(object?.name) ? object.name.map((e) => globalThis.String(e)) : [],
            page: isSet(object.page) ? globalThis.Number(object.page) : 0,
            limit: isSet(object.limit) ? globalThis.Number(object.limit) : 0,
        };
    },
    toJSON(message) {
        const obj = {};
        if (message.ids?.length) {
            obj.ids = message.ids;
        }
        if (message.name?.length) {
            obj.name = message.name;
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
        return exports.FindProductsWithPaginationRequest.fromPartial(base ?? {});
    },
    fromPartial(object) {
        const message = createBaseFindProductsWithPaginationRequest();
        message.ids = object.ids?.map((e) => e) || [];
        message.name = object.name?.map((e) => e) || [];
        message.page = object.page ?? 0;
        message.limit = object.limit ?? 0;
        return message;
    },
};
function createBaseFindProductsWithPaginationResponse() {
    return { data: [], limit: 0, page: 0, total: 0 };
}
exports.FindProductsWithPaginationResponse = {
    encode(message, writer = new wire_1.BinaryWriter()) {
        for (const v of message.data) {
            exports.Product.encode(v, writer.uint32(10).fork()).join();
        }
        if (message.limit !== 0) {
            writer.uint32(16).int32(message.limit);
        }
        if (message.page !== 0) {
            writer.uint32(24).int32(message.page);
        }
        if (message.total !== 0) {
            writer.uint32(32).int32(message.total);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof wire_1.BinaryReader ? input : new wire_1.BinaryReader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseFindProductsWithPaginationResponse();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1: {
                    if (tag !== 10) {
                        break;
                    }
                    message.data.push(exports.Product.decode(reader, reader.uint32()));
                    continue;
                }
                case 2: {
                    if (tag !== 16) {
                        break;
                    }
                    message.limit = reader.int32();
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
            data: globalThis.Array.isArray(object?.data) ? object.data.map((e) => exports.Product.fromJSON(e)) : [],
            limit: isSet(object.limit) ? globalThis.Number(object.limit) : 0,
            page: isSet(object.page) ? globalThis.Number(object.page) : 0,
            total: isSet(object.total) ? globalThis.Number(object.total) : 0,
        };
    },
    toJSON(message) {
        const obj = {};
        if (message.data?.length) {
            obj.data = message.data.map((e) => exports.Product.toJSON(e));
        }
        if (message.limit !== 0) {
            obj.limit = Math.round(message.limit);
        }
        if (message.page !== 0) {
            obj.page = Math.round(message.page);
        }
        if (message.total !== 0) {
            obj.total = Math.round(message.total);
        }
        return obj;
    },
    create(base) {
        return exports.FindProductsWithPaginationResponse.fromPartial(base ?? {});
    },
    fromPartial(object) {
        const message = createBaseFindProductsWithPaginationResponse();
        message.data = object.data?.map((e) => exports.Product.fromPartial(e)) || [];
        message.limit = object.limit ?? 0;
        message.page = object.page ?? 0;
        message.total = object.total ?? 0;
        return message;
    },
};
function createBaseFindProductByIdRequest() {
    return { id: "" };
}
exports.FindProductByIdRequest = {
    encode(message, writer = new wire_1.BinaryWriter()) {
        if (message.id !== "") {
            writer.uint32(10).string(message.id);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof wire_1.BinaryReader ? input : new wire_1.BinaryReader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseFindProductByIdRequest();
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
        return exports.FindProductByIdRequest.fromPartial(base ?? {});
    },
    fromPartial(object) {
        const message = createBaseFindProductByIdRequest();
        message.id = object.id ?? "";
        return message;
    },
};
function createBaseCreateProductRequest() {
    return { name: "", description: "", uom: "", image: "", price: 0, stock: 0 };
}
exports.CreateProductRequest = {
    encode(message, writer = new wire_1.BinaryWriter()) {
        if (message.name !== "") {
            writer.uint32(10).string(message.name);
        }
        if (message.description !== "") {
            writer.uint32(18).string(message.description);
        }
        if (message.uom !== "") {
            writer.uint32(26).string(message.uom);
        }
        if (message.image !== "") {
            writer.uint32(34).string(message.image);
        }
        if (message.price !== 0) {
            writer.uint32(41).double(message.price);
        }
        if (message.stock !== 0) {
            writer.uint32(48).int64(message.stock);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof wire_1.BinaryReader ? input : new wire_1.BinaryReader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseCreateProductRequest();
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
                case 2: {
                    if (tag !== 18) {
                        break;
                    }
                    message.description = reader.string();
                    continue;
                }
                case 3: {
                    if (tag !== 26) {
                        break;
                    }
                    message.uom = reader.string();
                    continue;
                }
                case 4: {
                    if (tag !== 34) {
                        break;
                    }
                    message.image = reader.string();
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
                    if (tag !== 48) {
                        break;
                    }
                    message.stock = longToNumber(reader.int64());
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
            name: isSet(object.name) ? globalThis.String(object.name) : "",
            description: isSet(object.description) ? globalThis.String(object.description) : "",
            uom: isSet(object.uom) ? globalThis.String(object.uom) : "",
            image: isSet(object.image) ? globalThis.String(object.image) : "",
            price: isSet(object.price) ? globalThis.Number(object.price) : 0,
            stock: isSet(object.stock) ? globalThis.Number(object.stock) : 0,
        };
    },
    toJSON(message) {
        const obj = {};
        if (message.name !== "") {
            obj.name = message.name;
        }
        if (message.description !== "") {
            obj.description = message.description;
        }
        if (message.uom !== "") {
            obj.uom = message.uom;
        }
        if (message.image !== "") {
            obj.image = message.image;
        }
        if (message.price !== 0) {
            obj.price = message.price;
        }
        if (message.stock !== 0) {
            obj.stock = Math.round(message.stock);
        }
        return obj;
    },
    create(base) {
        return exports.CreateProductRequest.fromPartial(base ?? {});
    },
    fromPartial(object) {
        const message = createBaseCreateProductRequest();
        message.name = object.name ?? "";
        message.description = object.description ?? "";
        message.uom = object.uom ?? "";
        message.image = object.image ?? "";
        message.price = object.price ?? 0;
        message.stock = object.stock ?? 0;
        return message;
    },
};
function createBaseCreateProductResponse() {
    return { id: "" };
}
exports.CreateProductResponse = {
    encode(message, writer = new wire_1.BinaryWriter()) {
        if (message.id !== "") {
            writer.uint32(10).string(message.id);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof wire_1.BinaryReader ? input : new wire_1.BinaryReader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseCreateProductResponse();
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
        return exports.CreateProductResponse.fromPartial(base ?? {});
    },
    fromPartial(object) {
        const message = createBaseCreateProductResponse();
        message.id = object.id ?? "";
        return message;
    },
};
function createBaseUpdateProductByIdRequest() {
    return {
        id: "",
        name: undefined,
        description: undefined,
        uom: undefined,
        image: undefined,
        price: undefined,
        stock: undefined,
    };
}
exports.UpdateProductByIdRequest = {
    encode(message, writer = new wire_1.BinaryWriter()) {
        if (message.id !== "") {
            writer.uint32(10).string(message.id);
        }
        if (message.name !== undefined) {
            writer.uint32(18).string(message.name);
        }
        if (message.description !== undefined) {
            writer.uint32(26).string(message.description);
        }
        if (message.uom !== undefined) {
            writer.uint32(34).string(message.uom);
        }
        if (message.image !== undefined) {
            writer.uint32(42).string(message.image);
        }
        if (message.price !== undefined) {
            writer.uint32(49).double(message.price);
        }
        if (message.stock !== undefined) {
            writer.uint32(56).int64(message.stock);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof wire_1.BinaryReader ? input : new wire_1.BinaryReader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseUpdateProductByIdRequest();
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
                    if (tag !== 26) {
                        break;
                    }
                    message.description = reader.string();
                    continue;
                }
                case 4: {
                    if (tag !== 34) {
                        break;
                    }
                    message.uom = reader.string();
                    continue;
                }
                case 5: {
                    if (tag !== 42) {
                        break;
                    }
                    message.image = reader.string();
                    continue;
                }
                case 6: {
                    if (tag !== 49) {
                        break;
                    }
                    message.price = reader.double();
                    continue;
                }
                case 7: {
                    if (tag !== 56) {
                        break;
                    }
                    message.stock = longToNumber(reader.int64());
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
            name: isSet(object.name) ? globalThis.String(object.name) : undefined,
            description: isSet(object.description) ? globalThis.String(object.description) : undefined,
            uom: isSet(object.uom) ? globalThis.String(object.uom) : undefined,
            image: isSet(object.image) ? globalThis.String(object.image) : undefined,
            price: isSet(object.price) ? globalThis.Number(object.price) : undefined,
            stock: isSet(object.stock) ? globalThis.Number(object.stock) : undefined,
        };
    },
    toJSON(message) {
        const obj = {};
        if (message.id !== "") {
            obj.id = message.id;
        }
        if (message.name !== undefined) {
            obj.name = message.name;
        }
        if (message.description !== undefined) {
            obj.description = message.description;
        }
        if (message.uom !== undefined) {
            obj.uom = message.uom;
        }
        if (message.image !== undefined) {
            obj.image = message.image;
        }
        if (message.price !== undefined) {
            obj.price = message.price;
        }
        if (message.stock !== undefined) {
            obj.stock = Math.round(message.stock);
        }
        return obj;
    },
    create(base) {
        return exports.UpdateProductByIdRequest.fromPartial(base ?? {});
    },
    fromPartial(object) {
        const message = createBaseUpdateProductByIdRequest();
        message.id = object.id ?? "";
        message.name = object.name ?? undefined;
        message.description = object.description ?? undefined;
        message.uom = object.uom ?? undefined;
        message.image = object.image ?? undefined;
        message.price = object.price ?? undefined;
        message.stock = object.stock ?? undefined;
        return message;
    },
};
function createBaseDeleteProductByIdRequest() {
    return { id: "" };
}
exports.DeleteProductByIdRequest = {
    encode(message, writer = new wire_1.BinaryWriter()) {
        if (message.id !== "") {
            writer.uint32(10).string(message.id);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof wire_1.BinaryReader ? input : new wire_1.BinaryReader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseDeleteProductByIdRequest();
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
        return exports.DeleteProductByIdRequest.fromPartial(base ?? {});
    },
    fromPartial(object) {
        const message = createBaseDeleteProductByIdRequest();
        message.id = object.id ?? "";
        return message;
    },
};
function createBaseDeleteProductByIdResponse() {
    return { message: "" };
}
exports.DeleteProductByIdResponse = {
    encode(message, writer = new wire_1.BinaryWriter()) {
        if (message.message !== "") {
            writer.uint32(10).string(message.message);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof wire_1.BinaryReader ? input : new wire_1.BinaryReader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseDeleteProductByIdResponse();
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
        return exports.DeleteProductByIdResponse.fromPartial(base ?? {});
    },
    fromPartial(object) {
        const message = createBaseDeleteProductByIdResponse();
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
function longToNumber(int64) {
    const num = globalThis.Number(int64.toString());
    if (num > globalThis.Number.MAX_SAFE_INTEGER) {
        throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
    }
    if (num < globalThis.Number.MIN_SAFE_INTEGER) {
        throw new globalThis.Error("Value is smaller than Number.MIN_SAFE_INTEGER");
    }
    return num;
}
function isSet(value) {
    return value !== null && value !== undefined;
}
//# sourceMappingURL=productMessage.js.map