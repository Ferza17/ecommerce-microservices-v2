package com.ferza17.ecommercemicroservicesv2.commerceservice.module.wishlist;

import com.ferza17.ecommercemicroservicesv2.proto.v1.commerce.WishlistServiceGrpc;
import com.ferza17.ecommercemicroservicesv2.proto.v1.commerce.Request;
import com.ferza17.ecommercemicroservicesv2.proto.v1.commerce.Response;
import io.opentelemetry.api.trace.Span;
import io.opentelemetry.api.trace.Tracer;
import io.opentelemetry.context.Scope;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.grpc.server.service.GrpcService;

@GrpcService
public class WishlistGrpcService extends WishlistServiceGrpc.WishlistServiceImplBase {
    @Autowired
    private WishlistUseCase useCase;
    @Autowired
    private Tracer tracer;


    @Override
    public void findWishlistItemWithPagination(Request.FindWishlistItemWithPaginationRequest request, io.grpc.stub.StreamObserver<Response.FindWishlistItemWithPaginationResponse> responseObserver) {
        Span span = this.tracer.spanBuilder("WishlistGrpcService.findWishlistItemWithPagination").startSpan();
        try (Scope scope = span.makeCurrent()) {
            // TODO:
            // 1. Trace Span
            // 2. Get RequestIDFilter
            // 3. Validate
            Response.FindWishlistItemWithPaginationResponse response = this.useCase.findWishlistItemWithPagination(request);
            responseObserver.onNext(response);
            responseObserver.onCompleted();
        } catch (Exception e) {
            span.recordException(e);
            responseObserver.onError(e);
        } finally {
            span.end();
        }
    }

    @Override
    public void addToWishlist(Request.AddToWishlistRequest request, io.grpc.stub.StreamObserver<Response.AddToWishlistResponse> responseObserver) {
        Span span = this.tracer.spanBuilder("WishlistGrpcService.addToWishlist").startSpan();
        try (Scope scope = span.makeCurrent()) {
            // TODO:
            // 1. Trace Span
            // 2. Get RequestIDFilter
            // 3. Validate
            Response.AddToWishlistResponse response = this.useCase.addToWishlist(request);
            responseObserver.onNext(response);
            responseObserver.onCompleted();
        } catch (Exception e) {
            span.recordException(e);
            responseObserver.onError(e);
        } finally {
            span.end();
        }
    }

    @Override
    public void deleteWishlistItemById(Request.DeleteWishlistItemByIdRequest request, io.grpc.stub.StreamObserver<Response.DeleteWishlistItemByIdResponse> responseObserver) {
        Span span = this.tracer.spanBuilder("WishlistGrpcService.deleteWishlistItemById").startSpan();
        try (Scope scope = span.makeCurrent()) {
            // TODO:
            // 1. Trace Span
            // 2. Get RequestIDFilter
            // 3. Validate
            Response.DeleteWishlistItemByIdResponse response = this.useCase.deleteCartItemById(request);
            responseObserver.onNext(response);
            responseObserver.onCompleted();
        } catch (Exception e) {
            span.recordException(e);
            responseObserver.onError(e);
        } finally {
            span.end();
        }
    }

}
