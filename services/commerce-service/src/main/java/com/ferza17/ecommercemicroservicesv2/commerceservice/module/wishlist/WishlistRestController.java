package com.ferza17.ecommercemicroservicesv2.commerceservice.module.wishlist;

import com.ferza17.ecommercemicroservicesv2.proto.v1.commerce.Request;
import com.ferza17.ecommercemicroservicesv2.proto.v1.commerce.Response;

import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import java.util.ArrayList;

@org.springframework.web.bind.annotation.RestController
@RequestMapping("/api/v1/commerce/wishlists")
public class WishlistRestController {
    private final WishlistUseCase WishlistUseCase;

    public WishlistRestController(WishlistUseCase wishlistUseCase) {
        WishlistUseCase = wishlistUseCase;
    }


    @PostMapping("/items")
    public ResponseEntity<Response.CreateCartItemResponse> createCartItem(@RequestBody Request.CreateCartItemRequest request) {
        // TODO: Implement Me
        return null;
    }

    @GetMapping("/items")
    public ResponseEntity<Response.CreateCartItemResponse> findWishlistItemWithPagination(
            @PathVariable("user_id") String userId,
            @PathVariable("product_ids") ArrayList<String> productIds,
            @PathVariable("page") int page,
            @PathVariable("limit") int limit) {
        // TODO: Implement Me
        return null;
    }

    @DeleteMapping("/items/{id}")
    public ResponseEntity<Response.DeleteCartItemByIdResponse> deleteCartItemById(@PathVariable("id") String id) {
        // TODO: Implement Me
        return null;
    }

}
