"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.CreateEventStoreResponse = exports.EventStore = exports.protobufPackage = void 0;
const wire_1 = require("@bufbuild/protobuf/wire");
const struct_1 = require("../../google/protobuf/struct");
const timestamp_1 = require("../../google/protobuf/timestamp");
exports.protobufPackage = "event_v1";
function createBaseEventStore() {
    return {
        id: "",
        requestId: "",
        service: "",
        eventType: "",
        status: "",
        payload: undefined,
        previousState: undefined,
        createdAt: undefined,
        updatedAt: undefined,
    };
}
exports.EventStore = {
    encode(message, writer = new wire_1.BinaryWriter()) {
        if (message.id !== "") {
            writer.uint32(10).string(message.id);
        }
        if (message.requestId !== "") {
            writer.uint32(18).string(message.requestId);
        }
        if (message.service !== "") {
            writer.uint32(26).string(message.service);
        }
        if (message.eventType !== "") {
            writer.uint32(34).string(message.eventType);
        }
        if (message.status !== "") {
            writer.uint32(42).string(message.status);
        }
        if (message.payload !== undefined) {
            struct_1.Struct.encode(struct_1.Struct.wrap(message.payload), writer.uint32(50).fork()).join();
        }
        if (message.previousState !== undefined) {
            struct_1.Struct.encode(struct_1.Struct.wrap(message.previousState), writer.uint32(58).fork()).join();
        }
        if (message.createdAt !== undefined) {
            timestamp_1.Timestamp.encode(toTimestamp(message.createdAt), writer.uint32(66).fork()).join();
        }
        if (message.updatedAt !== undefined) {
            timestamp_1.Timestamp.encode(toTimestamp(message.updatedAt), writer.uint32(74).fork()).join();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof wire_1.BinaryReader ? input : new wire_1.BinaryReader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseEventStore();
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
                    message.requestId = reader.string();
                    continue;
                }
                case 3: {
                    if (tag !== 26) {
                        break;
                    }
                    message.service = reader.string();
                    continue;
                }
                case 4: {
                    if (tag !== 34) {
                        break;
                    }
                    message.eventType = reader.string();
                    continue;
                }
                case 5: {
                    if (tag !== 42) {
                        break;
                    }
                    message.status = reader.string();
                    continue;
                }
                case 6: {
                    if (tag !== 50) {
                        break;
                    }
                    message.payload = struct_1.Struct.unwrap(struct_1.Struct.decode(reader, reader.uint32()));
                    continue;
                }
                case 7: {
                    if (tag !== 58) {
                        break;
                    }
                    message.previousState = struct_1.Struct.unwrap(struct_1.Struct.decode(reader, reader.uint32()));
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
            requestId: isSet(object.requestId) ? globalThis.String(object.requestId) : "",
            service: isSet(object.service) ? globalThis.String(object.service) : "",
            eventType: isSet(object.eventType) ? globalThis.String(object.eventType) : "",
            status: isSet(object.status) ? globalThis.String(object.status) : "",
            payload: isObject(object.payload) ? object.payload : undefined,
            previousState: isObject(object.previousState) ? object.previousState : undefined,
            createdAt: isSet(object.createdAt) ? fromJsonTimestamp(object.createdAt) : undefined,
            updatedAt: isSet(object.updatedAt) ? fromJsonTimestamp(object.updatedAt) : undefined,
        };
    },
    toJSON(message) {
        const obj = {};
        if (message.id !== "") {
            obj.id = message.id;
        }
        if (message.requestId !== "") {
            obj.requestId = message.requestId;
        }
        if (message.service !== "") {
            obj.service = message.service;
        }
        if (message.eventType !== "") {
            obj.eventType = message.eventType;
        }
        if (message.status !== "") {
            obj.status = message.status;
        }
        if (message.payload !== undefined) {
            obj.payload = message.payload;
        }
        if (message.previousState !== undefined) {
            obj.previousState = message.previousState;
        }
        if (message.createdAt !== undefined) {
            obj.createdAt = message.createdAt.toISOString();
        }
        if (message.updatedAt !== undefined) {
            obj.updatedAt = message.updatedAt.toISOString();
        }
        return obj;
    },
    create(base) {
        return exports.EventStore.fromPartial(base ?? {});
    },
    fromPartial(object) {
        const message = createBaseEventStore();
        message.id = object.id ?? "";
        message.requestId = object.requestId ?? "";
        message.service = object.service ?? "";
        message.eventType = object.eventType ?? "";
        message.status = object.status ?? "";
        message.payload = object.payload ?? undefined;
        message.previousState = object.previousState ?? undefined;
        message.createdAt = object.createdAt ?? undefined;
        message.updatedAt = object.updatedAt ?? undefined;
        return message;
    },
};
function createBaseCreateEventStoreResponse() {
    return { id: "" };
}
exports.CreateEventStoreResponse = {
    encode(message, writer = new wire_1.BinaryWriter()) {
        if (message.id !== "") {
            writer.uint32(10).string(message.id);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof wire_1.BinaryReader ? input : new wire_1.BinaryReader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseCreateEventStoreResponse();
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
        return exports.CreateEventStoreResponse.fromPartial(base ?? {});
    },
    fromPartial(object) {
        const message = createBaseCreateEventStoreResponse();
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
function isObject(value) {
    return typeof value === "object" && value !== null;
}
function isSet(value) {
    return value !== null && value !== undefined;
}
//# sourceMappingURL=eventStoreMessage.js.map