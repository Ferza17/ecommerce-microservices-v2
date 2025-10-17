package com.ferza17.ecommercemicroservicesv2.commerceservice.module.cart;

import com.ferza17.ecommercemicroservicesv2.proto.v1.commerce.CartServiceGrpc;
import com.ferza17.ecommercemicroservicesv2.proto.v1.commerce.Request;
import com.ferza17.ecommercemicroservicesv2.proto.v1.commerce.Response;
import io.opentelemetry.api.trace.Span;
import io.opentelemetry.api.trace.Tracer;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.grpc.server.service.GrpcService;
import io.opentelemetry.context.Scope;


@GrpcService
public class CartGrpcService extends CartServiceGrpc.CartServiceImplBase {
    @Autowired
    private CartUseCase useCase;
    @Autowired
    private Tracer tracer;

    @Override
    public void addToCart(Request.AddToCartRequest request, io.grpc.stub.StreamObserver<Response.AddToCartResponse> responseObserver) {
        Span span = this.tracer.spanBuilder("CartGrpcService.addToCart").startSpan();
        try (Scope scope = span.makeCurrent()) {
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
        Span span = this.tracer.spanBuilder("CartGrpcService.findCartItemsWithPagination").startSpan();
        try (Scope scope = span.makeCurrent()) {
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
        Span span = this.tracer.spanBuilder("CartGrpcService.deleteCartItemById").startSpan();
        try (Scope scope = span.makeCurrent()) {
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
