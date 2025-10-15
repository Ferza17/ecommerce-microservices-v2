package com.ferza17.ecommercemicroservicesv2.commerceservice.service.grpc;

import com.ferza17.ecommercemicroservicesv2.proto.v1.product.*;
import io.opentelemetry.sdk.OpenTelemetrySdk;

@org.springframework.stereotype.Service
public class ProductServiceGrpcClient {
    private final ProductServiceGrpc.ProductServiceBlockingStub productServiceBlockingStub;
    private final OpenTelemetrySdk openTelemetrySdk;


    public ProductServiceGrpcClient(ProductServiceGrpc.ProductServiceBlockingStub productServiceBlockingStub, OpenTelemetrySdk openTelemetrySdk) {
        this.productServiceBlockingStub = productServiceBlockingStub;
        this.openTelemetrySdk = openTelemetrySdk;
    }

    // TODO:
    // 1. Every GRPC Call send Metadata X-REQUEST-ID
    // 2. Every GRPC Call send Metadata AUTHORIZATION if token exists
    // 3. Every GRPC Call send Metadata TRACEPARENT
}
