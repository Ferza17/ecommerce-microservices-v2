package com.ferza17.ecommercemicroservicesv2.commerceservice.module.cart;

import com.ferza17.ecommercemicroservicesv2.proto.v1.commerce.Response;
import com.ferza17.ecommercemicroservicesv2.proto.v1.commerce.Request;
import com.ferza17.ecommercemicroservicesv2.proto.v1.commerce.Model;
import com.ferza17.ecommercemicroservicesv2.proto.v1.commerce.CartServiceGrpc;


import io.grpc.ManagedChannel;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

@RestController
@RequestMapping("/api/v1/commerce/carts")
public class PresenterHttp {
    private final ManagedChannel grpcChannel;
    private final CartServiceGrpc.CartServiceBlockingStub blockingStub;

    public PresenterHttp(ManagedChannel grpcChannel, CartServiceGrpc.CartServiceBlockingStub blockingStub) {
        this.grpcChannel = grpcChannel;
        this.blockingStub = blockingStub;
    }

    @PostMapping("/items")
    public ResponseEntity<Response.CreateCartItemResponse> createCartItem(@RequestBody Request.CreateCartItemRequest request) {
       // TODO: Implement Me
        return null;
    }

    @GetMapping("/items/{id}")
    public ResponseEntity<Model.CartItem> findCartItemById(@PathVariable String id) {
        // TODO: Implement Me
        return null;
    }

    @GetMapping("/items")
    public ResponseEntity<Response.FindCartItemsWithPaginationResponse>findCartItemsWithPagination(@PathVariable("page") int page, @PathVariable int size) {
        // TODO: Implement Me
        return null;
    }

    @PutMapping("/items/{id}")
    public ResponseEntity<Response.UpdateCartItemByIdResponse> updateCartItemById(@PathVariable String id, @RequestBody Request.UpdateCartItemByIdRequest request) {
        // TODO: Implement Me
        return null;
    }

    @DeleteMapping("/items/{id}")
    public ResponseEntity<Response.DeleteCartItemByIdResponse> deleteCartItemById(@PathVariable String id) {
        // TODO: Implement Me
        return null;
    }
}
