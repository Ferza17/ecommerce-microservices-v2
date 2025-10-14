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
    public ResponseEntity<Response.AddToCartResponse> addToCart(@RequestBody Request.AddToCartRequest request) {
        try {
            Response.AddToCartResponse response = this.cartUseCase.addToCart(request);
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
