"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.BytesValue = exports.StringValue = exports.BoolValue = exports.UInt32Value = exports.Int32Value = exports.UInt64Value = exports.Int64Value = exports.FloatValue = exports.DoubleValue = exports.protobufPackage = void 0;
const wire_1 = require("@bufbuild/protobuf/wire");
exports.protobufPackage = "google.protobuf";
function createBaseDoubleValue() {
    return { value: 0 };
}
exports.DoubleValue = {
    encode(message, writer = new wire_1.BinaryWriter()) {
        if (message.value !== 0) {
            writer.uint32(9).double(message.value);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof wire_1.BinaryReader ? input : new wire_1.BinaryReader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseDoubleValue();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1: {
                    if (tag !== 9) {
                        break;
                    }
                    message.value = reader.double();
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
        return { value: isSet(object.value) ? globalThis.Number(object.value) : 0 };
    },
    toJSON(message) {
        const obj = {};
        if (message.value !== 0) {
            obj.value = message.value;
        }
        return obj;
    },
    create(base) {
        return exports.DoubleValue.fromPartial(base ?? {});
    },
    fromPartial(object) {
        const message = createBaseDoubleValue();
        message.value = object.value ?? 0;
        return message;
    },
};
function createBaseFloatValue() {
    return { value: 0 };
}
exports.FloatValue = {
    encode(message, writer = new wire_1.BinaryWriter()) {
        if (message.value !== 0) {
            writer.uint32(13).float(message.value);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof wire_1.BinaryReader ? input : new wire_1.BinaryReader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseFloatValue();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1: {
                    if (tag !== 13) {
                        break;
                    }
                    message.value = reader.float();
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
        return { value: isSet(object.value) ? globalThis.Number(object.value) : 0 };
    },
    toJSON(message) {
        const obj = {};
        if (message.value !== 0) {
            obj.value = message.value;
        }
        return obj;
    },
    create(base) {
        return exports.FloatValue.fromPartial(base ?? {});
    },
    fromPartial(object) {
        const message = createBaseFloatValue();
        message.value = object.value ?? 0;
        return message;
    },
};
function createBaseInt64Value() {
    return { value: 0 };
}
exports.Int64Value = {
    encode(message, writer = new wire_1.BinaryWriter()) {
        if (message.value !== 0) {
            writer.uint32(8).int64(message.value);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof wire_1.BinaryReader ? input : new wire_1.BinaryReader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseInt64Value();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1: {
                    if (tag !== 8) {
                        break;
                    }
                    message.value = longToNumber(reader.int64());
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
        return { value: isSet(object.value) ? globalThis.Number(object.value) : 0 };
    },
    toJSON(message) {
        const obj = {};
        if (message.value !== 0) {
            obj.value = Math.round(message.value);
        }
        return obj;
    },
    create(base) {
        return exports.Int64Value.fromPartial(base ?? {});
    },
    fromPartial(object) {
        const message = createBaseInt64Value();
        message.value = object.value ?? 0;
        return message;
    },
};
function createBaseUInt64Value() {
    return { value: 0 };
}
exports.UInt64Value = {
    encode(message, writer = new wire_1.BinaryWriter()) {
        if (message.value !== 0) {
            writer.uint32(8).uint64(message.value);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof wire_1.BinaryReader ? input : new wire_1.BinaryReader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseUInt64Value();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1: {
                    if (tag !== 8) {
                        break;
                    }
                    message.value = longToNumber(reader.uint64());
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
        return { value: isSet(object.value) ? globalThis.Number(object.value) : 0 };
    },
    toJSON(message) {
        const obj = {};
        if (message.value !== 0) {
            obj.value = Math.round(message.value);
        }
        return obj;
    },
    create(base) {
        return exports.UInt64Value.fromPartial(base ?? {});
    },
    fromPartial(object) {
        const message = createBaseUInt64Value();
        message.value = object.value ?? 0;
        return message;
    },
};
function createBaseInt32Value() {
    return { value: 0 };
}
exports.Int32Value = {
    encode(message, writer = new wire_1.BinaryWriter()) {
        if (message.value !== 0) {
            writer.uint32(8).int32(message.value);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof wire_1.BinaryReader ? input : new wire_1.BinaryReader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseInt32Value();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1: {
                    if (tag !== 8) {
                        break;
                    }
                    message.value = reader.int32();
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
        return { value: isSet(object.value) ? globalThis.Number(object.value) : 0 };
    },
    toJSON(message) {
        const obj = {};
        if (message.value !== 0) {
            obj.value = Math.round(message.value);
        }
        return obj;
    },
    create(base) {
        return exports.Int32Value.fromPartial(base ?? {});
    },
    fromPartial(object) {
        const message = createBaseInt32Value();
        message.value = object.value ?? 0;
        return message;
    },
};
function createBaseUInt32Value() {
    return { value: 0 };
}
exports.UInt32Value = {
    encode(message, writer = new wire_1.BinaryWriter()) {
        if (message.value !== 0) {
            writer.uint32(8).uint32(message.value);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof wire_1.BinaryReader ? input : new wire_1.BinaryReader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseUInt32Value();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1: {
                    if (tag !== 8) {
                        break;
                    }
                    message.value = reader.uint32();
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
        return { value: isSet(object.value) ? globalThis.Number(object.value) : 0 };
    },
    toJSON(message) {
        const obj = {};
        if (message.value !== 0) {
            obj.value = Math.round(message.value);
        }
        return obj;
    },
    create(base) {
        return exports.UInt32Value.fromPartial(base ?? {});
    },
    fromPartial(object) {
        const message = createBaseUInt32Value();
        message.value = object.value ?? 0;
        return message;
    },
};
function createBaseBoolValue() {
    return { value: false };
}
exports.BoolValue = {
    encode(message, writer = new wire_1.BinaryWriter()) {
        if (message.value !== false) {
            writer.uint32(8).bool(message.value);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof wire_1.BinaryReader ? input : new wire_1.BinaryReader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseBoolValue();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1: {
                    if (tag !== 8) {
                        break;
                    }
                    message.value = reader.bool();
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
        return { value: isSet(object.value) ? globalThis.Boolean(object.value) : false };
    },
    toJSON(message) {
        const obj = {};
        if (message.value !== false) {
            obj.value = message.value;
        }
        return obj;
    },
    create(base) {
        return exports.BoolValue.fromPartial(base ?? {});
    },
    fromPartial(object) {
        const message = createBaseBoolValue();
        message.value = object.value ?? false;
        return message;
    },
};
function createBaseStringValue() {
    return { value: "" };
}
exports.StringValue = {
    encode(message, writer = new wire_1.BinaryWriter()) {
        if (message.value !== "") {
            writer.uint32(10).string(message.value);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof wire_1.BinaryReader ? input : new wire_1.BinaryReader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseStringValue();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1: {
                    if (tag !== 10) {
                        break;
                    }
                    message.value = reader.string();
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
        return { value: isSet(object.value) ? globalThis.String(object.value) : "" };
    },
    toJSON(message) {
        const obj = {};
        if (message.value !== "") {
            obj.value = message.value;
        }
        return obj;
    },
    create(base) {
        return exports.StringValue.fromPartial(base ?? {});
    },
    fromPartial(object) {
        const message = createBaseStringValue();
        message.value = object.value ?? "";
        return message;
    },
};
function createBaseBytesValue() {
    return { value: Buffer.alloc(0) };
}
exports.BytesValue = {
    encode(message, writer = new wire_1.BinaryWriter()) {
        if (message.value.length !== 0) {
            writer.uint32(10).bytes(message.value);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof wire_1.BinaryReader ? input : new wire_1.BinaryReader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseBytesValue();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1: {
                    if (tag !== 10) {
                        break;
                    }
                    message.value = Buffer.from(reader.bytes());
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
        return { value: isSet(object.value) ? Buffer.from(bytesFromBase64(object.value)) : Buffer.alloc(0) };
    },
    toJSON(message) {
        const obj = {};
        if (message.value.length !== 0) {
            obj.value = base64FromBytes(message.value);
        }
        return obj;
    },
    create(base) {
        return exports.BytesValue.fromPartial(base ?? {});
    },
    fromPartial(object) {
        const message = createBaseBytesValue();
        message.value = object.value ?? Buffer.alloc(0);
        return message;
    },
};
function bytesFromBase64(b64) {
    return Uint8Array.from(globalThis.Buffer.from(b64, "base64"));
}
function base64FromBytes(arr) {
    return globalThis.Buffer.from(arr).toString("base64");
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
//# sourceMappingURL=wrappers.js.map