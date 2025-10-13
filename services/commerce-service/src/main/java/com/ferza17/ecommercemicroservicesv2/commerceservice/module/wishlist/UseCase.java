package com.ferza17.ecommercemicroservicesv2.commerceservice.module.wishlist;

import com.ferza17.ecommercemicroservicesv2.proto.v1.commerce.Request;
import com.ferza17.ecommercemicroservicesv2.proto.v1.commerce.Response;

@org.springframework.stereotype.Service
public class UseCase {

    public Response.FindWishlistItemWithPaginationResponse findWishlistItemWithPagination(Request.FindWishlistItemWithPaginationRequest request) {
        try {
            // TODO: Fetch Repo Here
            return Response.FindWishlistItemWithPaginationResponse.newBuilder().build();
        } catch (Exception ex) {
            return Response.FindWishlistItemWithPaginationResponse.newBuilder().build();
        }
    }

    public Response.CreateWishlistItemResponse createWishlistItem(Request.CreateWishlistItemRequest request) {
        try {
            // TODO: Fetch Repo Here
            return Response.CreateWishlistItemResponse.newBuilder().build();
        } catch (Exception ex) {
            return Response.CreateWishlistItemResponse.newBuilder().build();
        }
    }

    public Response.DeleteWishlistItemByIdResponse deleteCartItemById(Request.DeleteWishlistItemByIdRequest request) {
        try {
            // TODO: Fetch Repo Here
            return Response.DeleteWishlistItemByIdResponse.newBuilder().build();
        } catch (Exception ex) {
            return Response.DeleteWishlistItemByIdResponse.newBuilder().build();
        }
    }
}
