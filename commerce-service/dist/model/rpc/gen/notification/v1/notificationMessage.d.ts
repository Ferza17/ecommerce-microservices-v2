import { BinaryReader, BinaryWriter } from "@bufbuild/protobuf/wire";
export declare const protobufPackage = "notification_v1";
export declare enum NotificationTypeEnum {
    NOTIFICATION_EMAIL_USER_OTP = 0,
    UNRECOGNIZED = -1
}
export declare function notificationTypeEnumFromJSON(object: any): NotificationTypeEnum;
export declare function notificationTypeEnumToJSON(object: NotificationTypeEnum): string;
export interface NotificationTemplate {
    id: string;
    type: string;
    template: string;
    templateVars: {
        [key: string]: any;
    } | undefined;
}
export interface SendOtpEmailNotificationRequest {
    otp: string;
    email: string;
    notificationType: NotificationTypeEnum;
}
export declare const NotificationTemplate: MessageFns<NotificationTemplate>;
export declare const SendOtpEmailNotificationRequest: MessageFns<SendOtpEmailNotificationRequest>;
type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;
export type DeepPartial<T> = T extends Builtin ? T : T extends globalThis.Array<infer U> ? globalThis.Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export interface MessageFns<T> {
    encode(message: T, writer?: BinaryWriter): BinaryWriter;
    decode(input: BinaryReader | Uint8Array, length?: number): T;
    fromJSON(object: any): T;
    toJSON(message: T): unknown;
    create(base?: DeepPartial<T>): T;
    fromPartial(object: DeepPartial<T>): T;
}
export {};
