package com.ferza17.ecommercemicroservicesv2.commerceservice.module.cart;

import com.ferza17.ecommercemicroservicesv2.proto.v1.commerce.Response;
import com.ferza17.ecommercemicroservicesv2.proto.v1.commerce.Request;
import com.ferza17.ecommercemicroservicesv2.proto.v1.commerce.Model;
import com.ferza17.ecommercemicroservicesv2.proto.v1.commerce.CartServiceGrpc;


import io.grpc.ManagedChannel;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

@org.springframework.web.bind.annotation.RestController
@RequestMapping("/api/v1/commerce/carts")
public class RestController {
    private final ManagedChannel grpcChannel;
    private final CartServiceGrpc.CartServiceBlockingStub blockingStub;

    public RestController(ManagedChannel grpcChannel, CartServiceGrpc.CartServiceBlockingStub blockingStub) {
        this.grpcChannel = grpcChannel;
        this.blockingStub = blockingStub;
    }

    @PostMapping("/items")
    public ResponseEntity<Response.CreateCartItemResponse> createCartItem(@RequestBody Request.CreateCartItemRequest request) {
        try {
            Response.CreateCartItemResponse response = this.blockingStub.createCartItem(request);
            return ResponseEntity.ok(response);
        } catch (Exception ex) {
            return ResponseEntity.internalServerError().build();
        }
    }

    @GetMapping("/items/{id}")
    public ResponseEntity<Model.CartItem> findCartItemById(@PathVariable String id) {
        try {
            Model.CartItem response = this.blockingStub.findCartItemById(Request.FindCartItemByIdRequest.newBuilder().setId(id).build());
            return ResponseEntity.ok(response);
        } catch (Exception ex) {
            return ResponseEntity.internalServerError().build();
        }
    }

    @GetMapping("/items")
    public ResponseEntity<Response.FindCartItemsWithPaginationResponse> findCartItemsWithPagination(@PathVariable("page") int page, @PathVariable("limit") int limit) {
        try {
            Response.FindCartItemsWithPaginationResponse response = this.blockingStub.findCartItemsWithPagination(Request.FindCartItemsWithPaginationRequest.newBuilder().setPage(page).setLimit(limit).build());
            return ResponseEntity.ok(response);
        } catch (Exception ex) {
            return ResponseEntity.internalServerError().build();
        }
    }

    @PutMapping("/items/{id}")
    public ResponseEntity<Response.UpdateCartItemByIdResponse> updateCartItemById(@PathVariable String id, @RequestBody Request.UpdateCartItemByIdRequest request) {
        try {
            Response.UpdateCartItemByIdResponse response = this.blockingStub.updateCartItemById(request.toBuilder().setId(id).build());
            return ResponseEntity.ok(response);
        } catch (Exception e) {
            throw new RuntimeException(e);
        }
    }

    @DeleteMapping("/items/{id}")
    public ResponseEntity<Response.DeleteCartItemByIdResponse> deleteCartItemById(@PathVariable String id) {
        try {
            Response.DeleteCartItemByIdResponse response = this.blockingStub.deleteCartItemById(Request.DeleteCartItemByIdRequest.newBuilder().setId(id).build());
            return ResponseEntity.ok(response);
        } catch (Exception ex) {
            return ResponseEntity.internalServerError().build();
        }
    }
}
