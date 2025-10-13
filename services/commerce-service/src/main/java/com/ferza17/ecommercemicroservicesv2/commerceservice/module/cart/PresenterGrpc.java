package com.ferza17.ecommercemicroservicesv2.commerceservice.module.cart;

import com.ferza17.ecommercemicroservicesv2.proto.v1.commerce.CartServiceGrpc;
import com.ferza17.ecommercemicroservicesv2.proto.v1.commerce.Request;
import com.ferza17.ecommercemicroservicesv2.proto.v1.commerce.Response;
import com.ferza17.ecommercemicroservicesv2.proto.v1.commerce.Model;
import org.springframework.grpc.server.service.GrpcService;

@GrpcService
public class PresenterGrpc extends CartServiceGrpc.CartServiceImplBase {

    @Override
    public void createCartItem(Request.CreateCartItemRequest request, io.grpc.stub.StreamObserver<Response.CreateCartItemResponse> responseObserver) {
        //TODO: Implement Me
        responseObserver.onCompleted();

    }

    @Override
    public void findCartItemById(Request.FindCartItemByIdRequest request, io.grpc.stub.StreamObserver<Model.CartItem> responseObserver) {
        //TODO: Implement Me
        responseObserver.onCompleted();

    }

    @Override
    public void findCartItemsWithPagination(Request.FindCartItemsWithPaginationRequest request, io.grpc.stub.StreamObserver<Response.FindCartItemsWithPaginationResponse> responseObserver) {
        //TODO: Implement Me
        responseObserver.onCompleted();

    }

    @Override
    public void updateCartItemById(Request.UpdateCartItemByIdRequest request, io.grpc.stub.StreamObserver<Response.UpdateCartItemByIdResponse> responseObserver) {
        //TODO: Implement Me
        responseObserver.onCompleted();

    }

    @Override
    public void deleteCartItemById(Request.DeleteCartItemByIdRequest request, io.grpc.stub.StreamObserver<Response.DeleteCartItemByIdResponse> responseObserver) {
        //TODO: Implement Me
        responseObserver.onCompleted();
    }
}
