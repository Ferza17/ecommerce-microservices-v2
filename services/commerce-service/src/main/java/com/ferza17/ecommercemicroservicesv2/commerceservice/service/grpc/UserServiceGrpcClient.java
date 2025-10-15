package com.ferza17.ecommercemicroservicesv2.commerceservice.service.grpc;

import com.ferza17.ecommercemicroservicesv2.proto.v1.user.*;
import io.opentelemetry.sdk.OpenTelemetrySdk;

@org.springframework.stereotype.Service
public class UserServiceGrpcClient {
    private final UserServiceGrpc.UserServiceBlockingStub userServiceBlockingStub;
    private final AuthServiceGrpc.AuthServiceBlockingStub authServiceBlockingStub;
    private final OpenTelemetrySdk openTelemetrySdk;


    public UserServiceGrpcClient(UserServiceGrpc.UserServiceBlockingStub userServiceBlockingStub, AuthServiceGrpc.AuthServiceBlockingStub authServiceBlockingStub, OpenTelemetrySdk openTelemetrySdk) {
        this.userServiceBlockingStub = userServiceBlockingStub;
        this.authServiceBlockingStub = authServiceBlockingStub;
        this.openTelemetrySdk = openTelemetrySdk;
    }
}
