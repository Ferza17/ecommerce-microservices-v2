package com.ferza17.ecommercemicroservicesv2.commerceservice.util;

import io.grpc.Status;
import org.springframework.http.HttpStatus;


public class ConvertStatus {
    public static HttpStatus mapGrpcStatusToHttp(Status.Code grpcStatus) {
        return switch (grpcStatus) {
            case OK -> HttpStatus.OK;
            case INVALID_ARGUMENT -> HttpStatus.BAD_REQUEST;
            case NOT_FOUND -> HttpStatus.NOT_FOUND;
            case ALREADY_EXISTS -> HttpStatus.CONFLICT;
            case PERMISSION_DENIED -> HttpStatus.FORBIDDEN;
            case UNAUTHENTICATED -> HttpStatus.UNAUTHORIZED;
            case RESOURCE_EXHAUSTED -> HttpStatus.TOO_MANY_REQUESTS;
            case FAILED_PRECONDITION -> HttpStatus.PRECONDITION_FAILED;
            case UNIMPLEMENTED -> HttpStatus.NOT_IMPLEMENTED;
            case UNAVAILABLE -> HttpStatus.SERVICE_UNAVAILABLE;
            case ABORTED -> HttpStatus.EXPECTATION_FAILED;
            default -> HttpStatus.INTERNAL_SERVER_ERROR;
        };
    }

    public static Status.Code mapHttpToGrpcStatus(HttpStatus httpStatus) {
        return switch (httpStatus) {
            case OK -> Status.Code.OK;
            case BAD_REQUEST -> Status.Code.INVALID_ARGUMENT;
            case NOT_FOUND -> Status.Code.NOT_FOUND;
            case CONFLICT -> Status.Code.ALREADY_EXISTS;
            case FORBIDDEN -> Status.Code.PERMISSION_DENIED;
            case UNAUTHORIZED -> Status.Code.UNAUTHENTICATED;
            case TOO_MANY_REQUESTS -> Status.Code.RESOURCE_EXHAUSTED;
            case PRECONDITION_FAILED -> Status.Code.FAILED_PRECONDITION;
            case NOT_IMPLEMENTED -> Status.Code.UNIMPLEMENTED;
            case SERVICE_UNAVAILABLE -> Status.Code.UNAVAILABLE;
            case EXPECTATION_FAILED -> Status.Code.ABORTED;
            default -> Status.Code.UNKNOWN;
        };
    }
}
