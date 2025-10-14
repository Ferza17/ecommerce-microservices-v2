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
    public void addToCart(Request.AddToCartRequest request,
                          io.grpc.stub.StreamObserver<Response.AddToCartResponse> responseObserver) {
        try {
            // TODO:
            // 1. Trace Span
            // 2. Get RequestIDFilter
            // 3. Validate
            Response.AddToCartResponse createCartItemResponse = this.useCase.addToCart(request);
            responseObserver.onNext(createCartItemResponse);
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
