package com.ferza17.ecommercemicroservicesv2.commerceservice.module.cart;

import com.ferza17.ecommercemicroservicesv2.proto.v1.commerce.CartServiceGrpc;
import com.ferza17.ecommercemicroservicesv2.proto.v1.commerce.Request;
import com.ferza17.ecommercemicroservicesv2.proto.v1.commerce.Response;
import io.opentelemetry.api.trace.Span;
import io.opentelemetry.sdk.OpenTelemetrySdk;

@org.springframework.grpc.server.service.GrpcService
public class CartGrpcService extends CartServiceGrpc.CartServiceImplBase {
    private final CartUseCase useCase;
    private final OpenTelemetrySdk openTelemetrySdk;

    public CartGrpcService(CartUseCase useCase, OpenTelemetrySdk openTelemetrySdk) {
        this.useCase = useCase;
        this.openTelemetrySdk = openTelemetrySdk;
    }

    @Override
    public void addToCart(Request.AddToCartRequest request, io.grpc.stub.StreamObserver<Response.AddToCartResponse> responseObserver) {
        Span span = this.openTelemetrySdk.getTracer(CartGrpcService.class.getSimpleName()).spanBuilder("addToCart").startSpan();
        try {
            // TODO:
            // 1. Trace Span
            // 2. Get RequestIDFilter
            // 3. Validate

            Response.AddToCartResponse createCartItemResponse = this.useCase.addToCart(request);
            responseObserver.onNext(createCartItemResponse);
            responseObserver.onCompleted();
        } catch (Exception ex) {
            span.recordException(ex);
            responseObserver.onError(ex);
        } finally {
            span.end();
        }
    }

    @Override
    public void findCartItemsWithPagination(Request.FindCartItemsWithPaginationRequest request, io.grpc.stub.StreamObserver<Response.FindCartItemsWithPaginationResponse> responseObserver) {
        Span span = this.openTelemetrySdk.getTracer(CartGrpcService.class.getSimpleName()).spanBuilder("findCartItemsWithPagination").startSpan();
        try {
            // TODO:
            // 1. Trace Span
            // 2. Get RequestIDFilter
            // 3. Validate

            Response.FindCartItemsWithPaginationResponse response = this.useCase.findCartItemsWithPagination(request);
            responseObserver.onNext(response);
            responseObserver.onCompleted();
        } catch (Exception ex) {
            span.recordException(ex);
            responseObserver.onError(ex);
        } finally {
            span.end();
        }
    }

    @Override
    public void deleteCartItemById(Request.DeleteCartItemByIdRequest request, io.grpc.stub.StreamObserver<Response.DeleteCartItemByIdResponse> responseObserver) {
        Span span = this.openTelemetrySdk.getTracer(CartGrpcService.class.getSimpleName()).spanBuilder("deleteCartItemById").startSpan();
        try {
            // TODO:
            // 1. Trace Span
            // 2. Get RequestIDFilter
            // 3. Validate

            Response.DeleteCartItemByIdResponse response = this.useCase.deleteCartItemById(request);
            responseObserver.onNext(response);
            responseObserver.onCompleted();
        } catch (Exception ex) {
            span.recordException(ex);
            responseObserver.onError(ex);
        } finally {
            span.end();
        }
    }
}
