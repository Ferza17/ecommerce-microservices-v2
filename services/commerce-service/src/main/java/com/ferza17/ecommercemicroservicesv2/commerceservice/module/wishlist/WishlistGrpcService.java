package com.ferza17.ecommercemicroservicesv2.commerceservice.module.wishlist;

import com.ferza17.ecommercemicroservicesv2.proto.v1.commerce.WishlistServiceGrpc;
import com.ferza17.ecommercemicroservicesv2.proto.v1.commerce.Request;
import com.ferza17.ecommercemicroservicesv2.proto.v1.commerce.Response;

@org.springframework.grpc.server.service.GrpcService
public class WishlistGrpcService extends WishlistServiceGrpc.WishlistServiceImplBase {

    private final WishlistUseCase useCase;

    public WishlistGrpcService(WishlistUseCase useCase) {
        this.useCase = useCase;
    }


    @Override
    public void findWishlistItemWithPagination(Request.FindWishlistItemWithPaginationRequest request, io.grpc.stub.StreamObserver<Response.FindWishlistItemWithPaginationResponse> responseObserver) {
        try {
            // TODO:
            // 1. Trace Span
            // 2. Get RequestIDFilter
            // 3. Validate
            Response.FindWishlistItemWithPaginationResponse response = this.useCase.findWishlistItemWithPagination(request);
            responseObserver.onNext(response);
            responseObserver.onCompleted();
        } catch (Exception e) {
            responseObserver.onError(e);
        }
    }

    @Override
    public void addToWishlist(Request.AddToWishlistRequest request, io.grpc.stub.StreamObserver<Response.AddToWishlistResponse> responseObserver) {
        try {
            // TODO:
            // 1. Trace Span
            // 2. Get RequestIDFilter
            // 3. Validate
            Response.AddToWishlistResponse response = this.useCase.addToWishlist(request);
            responseObserver.onNext(response);
            responseObserver.onCompleted();
        } catch (Exception e) {
            responseObserver.onError(e);
        }
    }

    @Override
    public void deleteWishlistItemById(Request.DeleteWishlistItemByIdRequest request, io.grpc.stub.StreamObserver<Response.DeleteWishlistItemByIdResponse> responseObserver) {
        try {
            // TODO:
            // 1. Trace Span
            // 2. Get RequestIDFilter
            // 3. Validate
            Response.DeleteWishlistItemByIdResponse response = this.useCase.deleteCartItemById(request);
            responseObserver.onNext(response);
            responseObserver.onCompleted();
        } catch (Exception e) {
            responseObserver.onError(e);
        }
    }

}
