package com.ferza17.ecommercemicroservicesv2.commerceservice.module.cart;

import com.ferza17.ecommercemicroservicesv2.proto.v1.commerce.CartServiceGrpc;
import com.ferza17.ecommercemicroservicesv2.proto.v1.commerce.Request.CreateCartItemRequest;
import com.ferza17.ecommercemicroservicesv2.proto.v1.commerce.Response.CreateCartItemResponse;
import com.ferza17.ecommercemicroservicesv2.proto.v1.commerce.Request.FindCartItemByIdRequest;
import com.ferza17.ecommercemicroservicesv2.proto.v1.commerce.Model.CartItem;
import com.ferza17.ecommercemicroservicesv2.proto.v1.commerce.Request.FindCartItemsWithPaginationRequest;
import com.ferza17.ecommercemicroservicesv2.proto.v1.commerce.Response.FindCartItemsWithPaginationResponse;
import com.ferza17.ecommercemicroservicesv2.proto.v1.commerce.Request.UpdateCartItemByIdRequest;
import com.ferza17.ecommercemicroservicesv2.proto.v1.commerce.Response.UpdateCartItemByIdResponse;
import com.ferza17.ecommercemicroservicesv2.proto.v1.commerce.Request.DeleteCartItemByIdRequest;
import com.ferza17.ecommercemicroservicesv2.proto.v1.commerce.Response.DeleteCartItemByIdResponse;

public class Presenter extends CartServiceGrpc.CartServiceImplBase {

    @Override
    public void createCartItem(CreateCartItemRequest request, io.grpc.stub.StreamObserver<CreateCartItemResponse> responseObserver) {
        //TODO: Implement Me
        responseObserver.onCompleted();

    }

    @Override
    public void findCartItemById(FindCartItemByIdRequest request, io.grpc.stub.StreamObserver<CartItem> responseObserver) {
        //TODO: Implement Me
        responseObserver.onCompleted();

    }

    @Override
    public void findCartItemsWithPagination(FindCartItemsWithPaginationRequest request, io.grpc.stub.StreamObserver<FindCartItemsWithPaginationResponse> responseObserver) {
        //TODO: Implement Me
        responseObserver.onCompleted();

    }

    @Override
    public void updateCartItemById(UpdateCartItemByIdRequest request, io.grpc.stub.StreamObserver<UpdateCartItemByIdResponse> responseObserver) {
        //TODO: Implement Me
        responseObserver.onCompleted();

    }

    @Override
    public void deleteCartItemById(DeleteCartItemByIdRequest request, io.grpc.stub.StreamObserver<DeleteCartItemByIdResponse> responseObserver) {
        //TODO: Implement Me
        responseObserver.onCompleted();
    }
}
