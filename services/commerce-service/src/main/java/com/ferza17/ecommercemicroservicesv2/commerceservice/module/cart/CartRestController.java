package com.ferza17.ecommercemicroservicesv2.commerceservice.module.cart;

import com.ferza17.ecommercemicroservicesv2.proto.v1.commerce.Response;
import com.ferza17.ecommercemicroservicesv2.proto.v1.commerce.Request;
import com.ferza17.ecommercemicroservicesv2.proto.v1.commerce.Model;


import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

@org.springframework.web.bind.annotation.RestController("cartRestController")
@RequestMapping("/api/v1/commerce/carts")
public class CartRestController {
    private final CartUseCase cartUseCase;

    public CartRestController(CartUseCase cartUseCase) {
        this.cartUseCase = cartUseCase;
    }


    @PostMapping("/items")
    public ResponseEntity<Response.CreateCartItemResponse> createCartItem(@RequestBody Request.CreateCartItemRequest request) {
        try {
            Response.CreateCartItemResponse response = this.cartUseCase.createCartItem(request);
            return ResponseEntity.ok(response);
        } catch (Exception ex) {
            return ResponseEntity.internalServerError().build();
        }
    }

    @GetMapping("/items/{id}")
    public ResponseEntity<Model.CartItem> findCartItemById(@PathVariable String id) {
        try {
            Model.CartItem response = this.cartUseCase.findCartItemById(Request.FindCartItemByIdRequest.newBuilder().setId(id).build());
            return ResponseEntity.ok(response);
        } catch (Exception ex) {
            return ResponseEntity.internalServerError().build();
        }
    }

    @GetMapping("/items")
    public ResponseEntity<Response.FindCartItemsWithPaginationResponse> findCartItemsWithPagination(@PathVariable("page") int page, @PathVariable("limit") int limit) {
        try {
            Response.FindCartItemsWithPaginationResponse response = this.cartUseCase.findCartItemsWithPagination(Request.FindCartItemsWithPaginationRequest.newBuilder().setPage(page).setLimit(limit).build());
            return ResponseEntity.ok(response);
        } catch (Exception ex) {
            return ResponseEntity.internalServerError().build();
        }
    }

    @PutMapping("/items/{id}")
    public ResponseEntity<Response.UpdateCartItemByIdResponse> updateCartItemById(@PathVariable String id, @RequestBody Request.UpdateCartItemByIdRequest request) {
        try {
            Response.UpdateCartItemByIdResponse response = this.cartUseCase.updateCartItemById(request.toBuilder().setId(id).build());
            return ResponseEntity.ok(response);
        } catch (Exception e) {
            throw new RuntimeException(e);
        }
    }

    @DeleteMapping("/items/{id}")
    public ResponseEntity<Response.DeleteCartItemByIdResponse> deleteCartItemById(@PathVariable String id) {
        try {
            Response.DeleteCartItemByIdResponse response = this.cartUseCase.deleteCartItemById(Request.DeleteCartItemByIdRequest.newBuilder().setId(id).build());
            return ResponseEntity.ok(response);
        } catch (Exception ex) {
            return ResponseEntity.internalServerError().build();
        }
    }
}
