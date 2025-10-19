package com.ferza17.ecommercemicroservicesv2.commerceservice.pkg.exception;

import lombok.Getter;

public class BaseException extends RuntimeException{
    @Getter
    private final BaseErrorCode baseErrorCode;

    public BaseException(BaseErrorCode baseErrorCode) {
        super(baseErrorCode.getMessage());
        this.baseErrorCode = baseErrorCode;
    }

    public BaseException(BaseErrorCode baseErrorCode, String message) {
        super(message);
        this.baseErrorCode = baseErrorCode;
    }
}
