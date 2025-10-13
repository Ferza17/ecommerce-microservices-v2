package com.ferza17.ecommercemicroservicesv2.commerceservice.module.cart;

import com.ferza17.ecommercemicroservicesv2.proto.v1.commerce.Model;
import com.ferza17.ecommercemicroservicesv2.proto.v1.commerce.Request;
import com.ferza17.ecommercemicroservicesv2.proto.v1.commerce.Response;


@org.springframework.stereotype.Service
public class CartUseCase {

    public Response.CreateCartItemResponse createCartItem(Request.CreateCartItemRequest request) {
        try {
            // TODO: Fetch Repo Here
            return Response.CreateCartItemResponse.newBuilder().build();
        } catch (Exception ex) {
            return Response.CreateCartItemResponse.newBuilder().build();
        }
    }

    public Model.CartItem findCartItemById(Request.FindCartItemByIdRequest request) {
        try {
            // TODO: Fetch Repo Here
            return Model.CartItem.newBuilder().build();
        } catch (Exception ex) {
            return Model.CartItem.newBuilder().build();
        }
    }

    public Response.FindCartItemsWithPaginationResponse findCartItemsWithPagination(Request.FindCartItemsWithPaginationRequest request) {
        try {
            // TODO: Fetch Repo Here
            return Response.FindCartItemsWithPaginationResponse.newBuilder().build();
        } catch (Exception ex) {
            return Response.FindCartItemsWithPaginationResponse.newBuilder().build();
        }
    }

    public Response.UpdateCartItemByIdResponse updateCartItemById(Request.UpdateCartItemByIdRequest request) {
        try {
            // TODO: Fetch Repo Here
            return Response.UpdateCartItemByIdResponse.newBuilder().build();
        } catch (Exception ex) {
            return Response.UpdateCartItemByIdResponse.newBuilder().build();
        }
    }

    public Response.DeleteCartItemByIdResponse deleteCartItemById(Request.DeleteCartItemByIdRequest request) {
        try {
            // TODO: Fetch Repo Here
            return Response.DeleteCartItemByIdResponse.newBuilder().build();
        } catch (Exception ex) {
            return Response.DeleteCartItemByIdResponse.newBuilder().build();
        }
    }


}
