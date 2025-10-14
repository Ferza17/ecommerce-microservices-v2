package com.ferza17.ecommercemicroservicesv2.commerceservice.module.cart;

import com.ferza17.ecommercemicroservicesv2.commerceservice.model.mapper.CartMapper;
import com.ferza17.ecommercemicroservicesv2.commerceservice.model.mongodb.CartModelMongoDB;
import com.ferza17.ecommercemicroservicesv2.proto.v1.commerce.Request;
import com.ferza17.ecommercemicroservicesv2.proto.v1.commerce.Response;
import org.bson.types.ObjectId;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.PageRequest;

import java.time.Instant;


@org.springframework.stereotype.Service
public class CartUseCase {
    private final CartMongoDBRepository cartMongoDBRepository;

    public CartUseCase(CartMongoDBRepository cartMongoDBRepository) {
        this.cartMongoDBRepository = cartMongoDBRepository;
    }

    public Response.AddToCartResponse addToCart(Request.AddToCartRequest request) {
        try {
            // TODO:
            // 1. Validate in DB
            // 2. Insert Via Sink Connector Event
            // 3. Insert Via Sink Connector Commerce

            Instant now = Instant
                    .now();
            CartModelMongoDB cart = CartModelMongoDB
                    .builder()
                    .id(ObjectId.get().toHexString())
                    .userId(request.getUserId())
                    .productId(request.getProductId())
                    .userId(request.getUserId())
                    .qty(request.getQty())
                    .price(request.getPrice())
                    .createdAt(now)
                    .updatedAt(now)
                    .build();

            this.cartMongoDBRepository.save(cart);
            return Response
                    .AddToCartResponse
                    .newBuilder()
                    .setStatus("success")
                    .setMessage("addToCart")
                    .setData(Response.AddToCartResponse.AddToCartResponseData.newBuilder().setId(cart.getId()).build())
                    .build();
        } catch (Exception ex) {
            return Response
                    .AddToCartResponse
                    .newBuilder()
                    .setStatus("failure")
                    .setMessage("addToCart")
                    .build();
        }
    }

    public Response.FindCartItemsWithPaginationResponse findCartItemsWithPagination(Request.FindCartItemsWithPaginationRequest request) {
        try {
            int page = Math.max(request.getPage() - 1, 0);
            PageRequest pageRequest = PageRequest
                    .of(page, request.getLimit());

            Page<CartModelMongoDB> cartPage = this
                    .cartMongoDBRepository
                    .findAllWithPagination(request.getUserId(), pageRequest);

            Response.FindCartItemsWithPaginationResponse.FindCartItemsWithPaginationResponseData.Builder responseData = Response
                    .FindCartItemsWithPaginationResponse
                    .FindCartItemsWithPaginationResponseData
                    .newBuilder()
                    .setLimit(cartPage.getSize())
                    .setPage(cartPage.getNumber() + 1)
                    .setTotal(Math.toIntExact(cartPage.getTotalElements()));

            cartPage
                    .stream()
                    .forEach(c -> responseData.addItems(CartMapper.toProto(c)));

            return Response
                    .FindCartItemsWithPaginationResponse
                    .newBuilder()
                    .setStatus("success")
                    .setMessage("FindCartItemsWithPagination")
                    .setData(responseData.build())
                    .build();

        } catch (Exception ex) {

            return Response
                    .FindCartItemsWithPaginationResponse
                    .newBuilder()
                    .setStatus("failure")
                    .setMessage("FindCartItemsWithPagination")
                    .build();
        }
    }

    public Response.DeleteCartItemByIdResponse deleteCartItemById(Request.DeleteCartItemByIdRequest request) {
        try {
            // TODO:
            // 1. Validate in DB
            // 2. Delete in DB Event
            // 3. Delete in DB Commerce

            this.cartMongoDBRepository.deleteById(request.getId());
            return Response
                    .DeleteCartItemByIdResponse
                    .newBuilder()
                    .setStatus("success")
                    .setMessage("deleteCartItemById")
                    .build();
        } catch (Exception ex) {

            return Response
                    .DeleteCartItemByIdResponse
                    .newBuilder()
                    .setStatus("failure")
                    .setMessage("deleteCartItemById")
                    .build();
        }
    }


}
