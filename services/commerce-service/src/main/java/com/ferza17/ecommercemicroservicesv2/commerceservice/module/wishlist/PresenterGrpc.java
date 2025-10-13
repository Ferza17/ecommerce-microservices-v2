package com.ferza17.ecommercemicroservicesv2.commerceservice.module.wishlist;

import com.ferza17.ecommercemicroservicesv2.proto.v1.commerce.WishlistServiceGrpc;
import com.ferza17.ecommercemicroservicesv2.proto.v1.commerce.Request;
import com.ferza17.ecommercemicroservicesv2.proto.v1.commerce.Response;
import org.springframework.grpc.server.service.GrpcService;

@GrpcService
public class PresenterGrpc extends WishlistServiceGrpc.WishlistServiceImplBase {

    @Override
    public void findWishlistItemWithPagination(Request.FindWishlistItemWithPaginationRequest request, io.grpc.stub.StreamObserver<Response.FindWishlistItemWithPaginationResponse> responseObserver) {
        //TODO: Implement Me
       responseObserver.onCompleted();
    }

    @Override
    public void createWishlistItem(Request.CreateWishlistItemRequest request, io.grpc.stub.StreamObserver<Response.CreateWishlistItemResponse> responseObserver) {
        //TODO: Implement Me
        responseObserver.onCompleted();
    }

    @Override
    public void deleteWishlistItemById(Request.DeleteWishlistItemByIdRequest request, io.grpc.stub.StreamObserver<Response.DeleteWishlistItemByIdResponse> responseObserver) {
        //TODO: Implement Me
        responseObserver.onCompleted();
    }

}
