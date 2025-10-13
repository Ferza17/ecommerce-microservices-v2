package com.ferza17.ecommercemicroservicesv2.commerceservice.module.cart;

import com.ferza17.ecommercemicroservicesv2.proto.v1.commerce.CartServiceGrpc;
import com.ferza17.ecommercemicroservicesv2.proto.v1.commerce.Request;
import com.ferza17.ecommercemicroservicesv2.proto.v1.commerce.Response;
import com.ferza17.ecommercemicroservicesv2.proto.v1.commerce.Model;

@org.springframework.grpc.server.service.GrpcService
public class CartGrpcService extends CartServiceGrpc.CartServiceImplBase {

    private final CartUseCase useCase;

    public CartGrpcService(CartUseCase useCase) {
        this.useCase = useCase;
    }


    @Override
    public void createCartItem(Request.CreateCartItemRequest request, io.grpc.stub.StreamObserver<Response.CreateCartItemResponse> responseObserver) {
        try {
            // TODO:
            // 1. Trace Span
            // 2. Get RequestIDFilter
            // 3. Validate
            Response.CreateCartItemResponse createCartItemResponse = this.useCase.createCartItem(request);
            responseObserver.onNext(createCartItemResponse);
            responseObserver.onCompleted();
        } catch (Exception ex) {
            responseObserver.onError(ex);
        }
    }

    @Override
    public void findCartItemById(Request.FindCartItemByIdRequest request, io.grpc.stub.StreamObserver<Model.CartItem> responseObserver) {
        try {
            // TODO:
            // 1. Trace Span
            // 2. Get RequestIDFilter
            // 3. Validate
            Model.CartItem cartItem = this.useCase.findCartItemById(request);
            responseObserver.onNext(cartItem);
            responseObserver.onCompleted();
        } catch (Exception ex) {
            responseObserver.onError(ex);
        }

    }

    @Override
    public void findCartItemsWithPagination(Request.FindCartItemsWithPaginationRequest request, io.grpc.stub.StreamObserver<Response.FindCartItemsWithPaginationResponse> responseObserver) {
        try {
            // TODO:
            // 1. Trace Span
            // 2. Get RequestIDFilter
            // 3. Validate
            Response.FindCartItemsWithPaginationResponse response = this.useCase.findCartItemsWithPagination(request);
            responseObserver.onNext(response);
            responseObserver.onCompleted();
        } catch (Exception ex) {
            responseObserver.onError(ex);
        }
    }

    @Override
    public void updateCartItemById(Request.UpdateCartItemByIdRequest request, io.grpc.stub.StreamObserver<Response.UpdateCartItemByIdResponse> responseObserver) {
        try {
            // TODO:
            // 1. Trace Span
            // 2. Get RequestIDFilter
            // 3. Validate
            Response.UpdateCartItemByIdResponse response = this.useCase.updateCartItemById(request);
            responseObserver.onNext(response);
            responseObserver.onCompleted();
        } catch (Exception ex) {
            responseObserver.onError(ex);
        }
    }

    @Override
    public void deleteCartItemById(Request.DeleteCartItemByIdRequest request, io.grpc.stub.StreamObserver<Response.DeleteCartItemByIdResponse> responseObserver) {
        try {
            // TODO:
            // 1. Trace Span
            // 2. Get RequestIDFilter
            // 3. Validate
            Response.DeleteCartItemByIdResponse response = this.useCase.deleteCartItemById(request);
            responseObserver.onNext(response);
            responseObserver.onCompleted();
        } catch (Exception e) {
            responseObserver.onError(e);
        }
    }
}
