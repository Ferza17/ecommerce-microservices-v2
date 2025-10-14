package com.ferza17.ecommercemicroservicesv2.commerceservice.module.wishlist;

import com.ferza17.ecommercemicroservicesv2.commerceservice.model.mapper.WishlistMapper;
import com.ferza17.ecommercemicroservicesv2.commerceservice.model.mongodb.WishlistModelMongoDB;
import com.ferza17.ecommercemicroservicesv2.proto.v1.commerce.Request;
import com.ferza17.ecommercemicroservicesv2.proto.v1.commerce.Response;
import org.bson.types.ObjectId;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.PageRequest;

import java.time.Instant;

@org.springframework.stereotype.Service
public class WishlistUseCase {
    private final WishlistMongoDBRepository wishlistMongoDBRepository;

    public WishlistUseCase(WishlistMongoDBRepository wishlistMongoDBRepository) {
        this.wishlistMongoDBRepository = wishlistMongoDBRepository;
    }

    public Response.AddToWishlistResponse addToWishlist(Request.AddToWishlistRequest request) {
        try {
            // TODO:
            // 1. Validate in DB
            // 2. Insert Via Sink Connector Event
            // 3. Insert Via Sink Connector Commerce Carts

            Instant now = Instant.now();
            WishlistModelMongoDB wishlist = WishlistModelMongoDB
                    .builder()
                    .id(ObjectId.get().toHexString())
                    .user_id(request.getUserId())
                    .product_id(request.getProductId())
                    .created_at(now)
                    .updated_at(now)
                    .build();

            this.wishlistMongoDBRepository.save(wishlist);
            return Response
                    .AddToWishlistResponse
                    .newBuilder()
                    .setStatus("success")
                    .setMessage("addToWishlist")
                    .setData(Response
                            .AddToWishlistResponse
                            .AddToWishlistResponseData
                            .newBuilder()
                            .setId(wishlist.getId())
                            .build()
                    )
                    .build();
        } catch (Exception ex) {
            return Response
                    .AddToWishlistResponse
                    .newBuilder()
                    .setStatus("failure")
                    .setMessage("addToWishlist")
                    .build();
        }
    }

    public Response.FindWishlistItemWithPaginationResponse findWishlistItemWithPagination(Request.FindWishlistItemWithPaginationRequest request) {
        try {
            int page = Math.max(request.getPage() - 1, 0);
            PageRequest pageRequest = PageRequest
                    .of(page, request.getLimit());

            Page<WishlistModelMongoDB> wishlistPage = this
                    .wishlistMongoDBRepository
                    .findAllWithPagination(request.getUserId(), pageRequest);

            Response.FindWishlistItemWithPaginationResponse.FindWishlistItemWithPaginationResponseData.Builder responseData = Response
                    .FindWishlistItemWithPaginationResponse
                    .FindWishlistItemWithPaginationResponseData
                    .newBuilder()
                    .setLimit(wishlistPage.getSize())
                    .setPage(wishlistPage.getNumber() + 1);

            wishlistPage
                    .stream()
                    .forEach(w -> responseData.addItems(WishlistMapper.toProto(w)));

            return Response
                    .FindWishlistItemWithPaginationResponse
                    .newBuilder()
                    .setStatus("success")
                    .setMessage("findWishlistItemWithPagination")
                    .setData(responseData.build())
                    .build();
        } catch (Exception ex) {
            return Response
                    .FindWishlistItemWithPaginationResponse
                    .newBuilder()
                    .setStatus("failure")
                    .setMessage("findWishlistItemWithPagination")
                    .build();
        }
    }

    public Response.DeleteWishlistItemByIdResponse deleteCartItemById(Request.DeleteWishlistItemByIdRequest request) {
        try {
            // TODO:
            // 1. Validate in DB
            // 2. Delete In Collection Event Stores
            // 3. Delete In Collection Cart Item
            this.wishlistMongoDBRepository.deleteById(request.getId());
            return Response
                    .DeleteWishlistItemByIdResponse
                    .newBuilder()
                    .setStatus("success")
                    .setMessage("deleteWishlistItemById")
                    .build();
        } catch (Exception ex) {
            return Response
                    .DeleteWishlistItemByIdResponse
                    .newBuilder()
                    .setStatus("failure")
                    .setMessage("deleteWishlistItemById")
                    .build();
        }
    }
}
