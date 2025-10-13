package com.ferza17.ecommercemicroservicesv2.commerceservice.module.wishlist;

import com.ferza17.ecommercemicroservicesv2.proto.v1.commerce.WishlistServiceGrpc;
import com.ferza17.ecommercemicroservicesv2.proto.v1.commerce.Request.FindWishlistItemWithPaginationRequest;
import com.ferza17.ecommercemicroservicesv2.proto.v1.commerce.Response.FindWishlistItemWithPaginationResponse;
import com.ferza17.ecommercemicroservicesv2.proto.v1.commerce.Request.CreateWishlistItemRequest;
import com.ferza17.ecommercemicroservicesv2.proto.v1.commerce.Response.CreateWishlistItemResponse;
import com.ferza17.ecommercemicroservicesv2.proto.v1.commerce.Request.DeleteWishlistItemByIdRequest;
import com.ferza17.ecommercemicroservicesv2.proto.v1.commerce.Response.DeleteWishlistItemByIdResponse;

public class Presenter extends WishlistServiceGrpc.WishlistServiceImplBase {

    @Override
    public void findWishlistItemWithPagination(FindWishlistItemWithPaginationRequest request, io.grpc.stub.StreamObserver<FindWishlistItemWithPaginationResponse> responseObserver) {
        //TODO: Implement Me
       responseObserver.onCompleted();
    }

    @Override
    public void createWishlistItem(CreateWishlistItemRequest request, io.grpc.stub.StreamObserver<CreateWishlistItemResponse> responseObserver) {
        //TODO: Implement Me
        responseObserver.onCompleted();
    }

    @Override
    public void deleteWishlistItemById(DeleteWishlistItemByIdRequest request, io.grpc.stub.StreamObserver<DeleteWishlistItemByIdResponse> responseObserver) {
        //TODO: Implement Me
        responseObserver.onCompleted();
    }

}
