package com.ferza17.ecommercemicroservicesv2.commerceservice.module.wishlist;

import com.ferza17.ecommercemicroservicesv2.commerceservice.model.mapper.WishlistMapper;
import com.ferza17.ecommercemicroservicesv2.commerceservice.model.mongodb.WishlistModelMongoDB;
import com.ferza17.ecommercemicroservicesv2.proto.v1.commerce.Model.WishlistItem;
import com.ferza17.ecommercemicroservicesv2.proto.v1.commerce.Response.AddToWishlistResponse;
import com.ferza17.ecommercemicroservicesv2.proto.v1.commerce.Response.FindWishlistItemWithPaginationResponse;
import com.ferza17.ecommercemicroservicesv2.proto.v1.commerce.Response.DeleteWishlistItemByIdResponse;
import com.ferza17.ecommercemicroservicesv2.proto.v1.commerce.Request.AddToWishlistRequest;
import com.ferza17.ecommercemicroservicesv2.proto.v1.commerce.Request.FindWishlistItemWithPaginationRequest;
import com.ferza17.ecommercemicroservicesv2.proto.v1.commerce.Request.DeleteWishlistItemByIdRequest;
import com.ferza17.ecommercemicroservicesv2.proto.v1.product.ProductServiceGrpc;
import com.ferza17.ecommercemicroservicesv2.proto.v1.product.Model.Product;
import com.ferza17.ecommercemicroservicesv2.proto.v1.product.Request.FindProductByIdRequest;
import com.ferza17.ecommercemicroservicesv2.proto.v1.user.AuthServiceGrpc;
import com.ferza17.ecommercemicroservicesv2.proto.v1.user.Request;
import com.ferza17.ecommercemicroservicesv2.proto.v1.user.Response;
import io.opentelemetry.api.trace.Span;
import io.opentelemetry.api.trace.Tracer;
import io.opentelemetry.context.Scope;
import org.bson.types.ObjectId;
import org.slf4j.MDC;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.PageRequest;
import org.springframework.kafka.core.KafkaTemplate;

import java.time.Instant;

import static com.ferza17.ecommercemicroservicesv2.commerceservice.pkg.context.BaseContext.AUTHORIZATION_CONTEXT_KEY;
import static com.ferza17.ecommercemicroservicesv2.commerceservice.pkg.context.BaseContext.TRACEPARENT_CONTEXT_KEY;

@org.springframework.stereotype.Service
public class WishlistUseCase {
    @Autowired
    private WishlistMongoDBRepository wishlistMongoDBRepository;
    @Autowired
    private ProductServiceGrpc.ProductServiceBlockingStub productServiceBlockingStub;
    @Autowired
    private AuthServiceGrpc.AuthServiceBlockingStub authServiceBlockingStub;
    @Autowired
    private Tracer tracer;
    @Autowired
    private KafkaTemplate<String, WishlistItem> kafkaTemplateSinkMongoWishlist;

    public AddToWishlistResponse addToWishlist(AddToWishlistRequest request) {
        Span span = this.tracer.spanBuilder("WishlistUseCase.addToWishlist").startSpan();
        try (Scope scope = span.makeCurrent()) {
            String traceId = span.getSpanContext().getTraceId();
            MDC.put(TRACEPARENT_CONTEXT_KEY, traceId);
            Product product = this.productServiceBlockingStub.findProductById(FindProductByIdRequest.newBuilder().setId(request.getProductId()).build());
            Response.AuthUserFindUserByTokenResponse authResponse = this.authServiceBlockingStub.authUserFindUserByToken(Request.AuthUserFindUserByTokenRequest.newBuilder().setToken(MDC.get(AUTHORIZATION_CONTEXT_KEY).replaceAll("(?i)^Bearer\\s+", "")).build());
            WishlistModelMongoDB existingWishlist = this.wishlistMongoDBRepository.findByProductIdAndUserId(product.getId(), authResponse.getData().getUser().getId()).orElse(null);
            if (existingWishlist != null) {
                throw new Exception("Product already in wishlist");
            }

            Instant now = Instant.now();
            WishlistItem wishlist = WishlistMapper.toProto(WishlistModelMongoDB
                    .builder()
                    .id(ObjectId.get().toHexString())
                    .user_id(authResponse.getData().getUser().getId())
                    .product_id(request.getProductId())
                    .created_at(now)
                    .updated_at(now)
                    .build());

            this.kafkaTemplateSinkMongoWishlist.send("sink-mongo-commerce-wishlists", wishlist.getId(), wishlist);
            return AddToWishlistResponse
                    .newBuilder()
                    .setStatus("success")
                    .setMessage("addToWishlist")
                    .setData(AddToWishlistResponse
                            .AddToWishlistResponseData
                            .newBuilder()
                            .setId(wishlist.getId())
                            .build()
                    )
                    .build();
        } catch (Exception ex) {
            span.recordException(ex);
            return AddToWishlistResponse
                    .newBuilder()
                    .setStatus("failure")
                    .setMessage("addToWishlist")
                    .build();
        } finally {
            span.end();
        }
    }

    public FindWishlistItemWithPaginationResponse findWishlistItemWithPagination(FindWishlistItemWithPaginationRequest request) {
        Span span = this.tracer.spanBuilder("WishlistUseCase.findWishlistItemWithPagination").startSpan();
        try (Scope scope = span.makeCurrent()) {
            int page = Math.max(request.getPage() - 1, 0);
            PageRequest pageRequest = PageRequest
                    .of(page, request.getLimit());

            Page<WishlistModelMongoDB> wishlistPage = this
                    .wishlistMongoDBRepository
                    .findAllWithPagination(request.getUserId(), pageRequest);

            FindWishlistItemWithPaginationResponse.FindWishlistItemWithPaginationResponseData.Builder responseData = FindWishlistItemWithPaginationResponse
                    .FindWishlistItemWithPaginationResponseData
                    .newBuilder()
                    .setLimit(wishlistPage.getSize())
                    .setPage(wishlistPage.getNumber() + 1);

            wishlistPage
                    .stream()
                    .forEach(w -> responseData.addItems(WishlistMapper.toProto(w)));

            return FindWishlistItemWithPaginationResponse
                    .newBuilder()
                    .setStatus("success")
                    .setMessage("findWishlistItemWithPagination")
                    .setData(responseData.build())
                    .build();
        } catch (Exception ex) {
            span.recordException(ex);
            return FindWishlistItemWithPaginationResponse
                    .newBuilder()
                    .setStatus("failure")
                    .setMessage("findWishlistItemWithPagination")
                    .build();
        } finally {
            span.end();
        }
    }

    public DeleteWishlistItemByIdResponse deleteCartItemById(DeleteWishlistItemByIdRequest request) {
        Span span = this.tracer.spanBuilder("WishlistUseCase.deleteCartItemById").startSpan();
        try (Scope scope = span.makeCurrent()) {
            this.wishlistMongoDBRepository.deleteById(request.getId());
            return DeleteWishlistItemByIdResponse
                    .newBuilder()
                    .setStatus("success")
                    .setMessage("deleteWishlistItemById")
                    .build();
        } catch (Exception ex) {
            span.recordException(ex);
            return DeleteWishlistItemByIdResponse
                    .newBuilder()
                    .setStatus("failure")
                    .setMessage("deleteWishlistItemById")
                    .build();
        } finally {
            span.end();
        }
    }
}
