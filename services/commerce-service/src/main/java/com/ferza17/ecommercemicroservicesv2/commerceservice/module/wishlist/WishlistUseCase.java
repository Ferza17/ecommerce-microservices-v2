package com.ferza17.ecommercemicroservicesv2.commerceservice.module.wishlist;

import com.ferza17.ecommercemicroservicesv2.proto.v1.commerce.Request;
import com.ferza17.ecommercemicroservicesv2.proto.v1.commerce.Response;

@org.springframework.stereotype.Service
public class WishlistUseCase {

    public Response.FindWishlistItemWithPaginationResponse findWishlistItemWithPagination(Request.FindWishlistItemWithPaginationRequest request) {
        try {
            // TODO:
            // 1. Find in DB
            return Response.FindWishlistItemWithPaginationResponse.newBuilder().build();
        } catch (Exception ex) {
            return Response.FindWishlistItemWithPaginationResponse.newBuilder().build();
        }
    }

    public Response.AddToWishlistResponse addToWishlist(Request.AddToWishlistRequest request) {
        try {
            // TODO:
            // 1. Validate in DB
            // 2. Insert Via Sink Connector Event
            // 3. Insert Via Sink Connector Commerce Carts
            return Response.AddToWishlistResponse.newBuilder().build();
        } catch (Exception ex) {
            return Response.AddToWishlistResponse.newBuilder().build();
        }
    }


    public Response.DeleteWishlistItemByIdResponse deleteCartItemById(Request.DeleteWishlistItemByIdRequest request) {
        try {
            // TODO:
            // 1. Validate in DB
            // 2. Delete In Collection Event Stores
            // 3. Delete In Collection Cart Item
            return Response.DeleteWishlistItemByIdResponse.newBuilder().build();
        } catch (Exception ex) {
            return Response.DeleteWishlistItemByIdResponse.newBuilder().build();
        }
    }
}
