package com.ferza17.ecommercemicroservicesv2.commerceservice.pkg.exception;

import io.grpc.Status;
import lombok.Getter;

public enum BaseErrorCode {
    UNAUTHENTICATED(Status.UNAUTHENTICATED.getCode(), "Unauthenticated"),
    INVALID_ARGUMENT(Status.INVALID_ARGUMENT.getCode(), "Invalid argument"),
    NOT_FOUND(Status.NOT_FOUND.getCode(), "Resource not found"),
    INTERNAL_ERROR(Status.UNKNOWN.getCode(), "Internal server error"),
    BUSINESS_ERROR(Status.ABORTED.getCode(), "Business rule violated");

    @Getter
    private final Status.Code code;
    @Getter
    private final String message;

    BaseErrorCode(Status.Code code, String message) {
        this.code = code;
        this.message = message;
    }
}
