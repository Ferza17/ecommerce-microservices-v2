import { Observable } from 'rxjs';
export declare const protobufPackage = "grpc.reflection.v1alpha";
export interface ServerReflectionRequest {
    host: string;
    fileByFilename: string | undefined;
    fileContainingSymbol: string | undefined;
    fileContainingExtension: ExtensionRequest | undefined;
    allExtensionNumbersOfType: string | undefined;
    listServices: string | undefined;
}
export interface ExtensionRequest {
    containingType: string;
    extensionNumber: number;
}
export interface ServerReflectionResponse {
    validHost: string;
    originalRequest: ServerReflectionRequest | undefined;
    fileDescriptorResponse: FileDescriptorResponse | undefined;
    allExtensionNumbersResponse: ExtensionNumberResponse | undefined;
    listServicesResponse: ListServiceResponse | undefined;
    errorResponse: ErrorResponse | undefined;
}
export interface FileDescriptorResponse {
    fileDescriptorProto: Uint8Array[];
}
export interface ExtensionNumberResponse {
    baseTypeName: string;
    extensionNumber: number[];
}
export interface ListServiceResponse {
    service: ServiceResponse[];
}
export interface ServiceResponse {
    name: string;
}
export interface ErrorResponse {
    errorCode: number;
    errorMessage: string;
}
export declare const GRPC_REFLECTION_V1ALPHA_PACKAGE_NAME = "grpc.reflection.v1alpha";
export interface ServerReflectionClient {
    serverReflectionInfo(request: Observable<ServerReflectionRequest>): Observable<ServerReflectionResponse>;
}
export interface ServerReflectionController {
    serverReflectionInfo(request: Observable<ServerReflectionRequest>): Observable<ServerReflectionResponse>;
}
export declare function ServerReflectionControllerMethods(): (constructor: Function) => void;
export declare const SERVER_REFLECTION_SERVICE_NAME = "ServerReflection";
