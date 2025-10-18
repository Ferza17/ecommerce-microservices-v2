package com.ferza17.ecommercemicroservicesv2.commerceservice.module.cart;

import com.ferza17.ecommercemicroservicesv2.commerceservice.model.mapper.CartMapper;
import com.ferza17.ecommercemicroservicesv2.commerceservice.model.mongodb.CartModelMongoDB;

import com.ferza17.ecommercemicroservicesv2.proto.v1.commerce.Model;
import com.ferza17.ecommercemicroservicesv2.proto.v1.commerce.Response.*;
import com.ferza17.ecommercemicroservicesv2.proto.v1.commerce.Request.*;

import com.ferza17.ecommercemicroservicesv2.proto.v1.product.ProductServiceGrpc;
import com.ferza17.ecommercemicroservicesv2.proto.v1.product.Request.*;
import com.ferza17.ecommercemicroservicesv2.proto.v1.product.Model.*;

import com.ferza17.ecommercemicroservicesv2.proto.v1.user.AuthServiceGrpc;
import com.ferza17.ecommercemicroservicesv2.proto.v1.user.Request.*;
import com.ferza17.ecommercemicroservicesv2.proto.v1.user.Response.*;
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
public class CartUseCase {
    @Autowired
    private CartMongoDBRepository cartMongoDBRepository;
    @Autowired
    private ProductServiceGrpc.ProductServiceBlockingStub productServiceBlockingStub;
    @Autowired
    private AuthServiceGrpc.AuthServiceBlockingStub authServiceBlockingStub;
    @Autowired
    private KafkaTemplate<String, Model.CartItem> kafkaTemplateSinkMongoCart;
    @Autowired
    private Tracer tracer;

    public AddToCartResponse addToCart(AddToCartRequest request) throws Exception {
        Span span = this.tracer.spanBuilder("CartUseCase.addToCart").startSpan();
        try (Scope scope = span.makeCurrent()) {
            // TODO:
            // 1. Validate in DB
            // 2. Insert Via Sink Connector Event
            // 3. Insert Via Sink Connector Commerce
            String traceId = span.getSpanContext().getTraceId();
            MDC.put(TRACEPARENT_CONTEXT_KEY, traceId);
            Product product = this.productServiceBlockingStub.findProductById(FindProductByIdRequest.newBuilder().setId(request.getProductId()).build());
            AuthUserFindUserByTokenResponse authResponse = this.authServiceBlockingStub.authUserFindUserByToken(AuthUserFindUserByTokenRequest.newBuilder().setToken(MDC.get(AUTHORIZATION_CONTEXT_KEY).replaceAll("(?i)^Bearer\\s+", "")).build());
            CartModelMongoDB existingCart = this.cartMongoDBRepository.findByProductIdAndUserId(product.getId(), authResponse.getData().getUser().getId()).orElse(null);

            Instant now = Instant.now();
            if (existingCart == null) {
                existingCart = new CartModelMongoDB()
                        .builder()
                        .id(ObjectId.get().toHexString())
                        .userId(authResponse.getData().getUser().getId())
                        .productId(request.getProductId())
                        .qty(request.getQty())
                        .price(request.getQty() * product.getPrice())
                        .createdAt(now)
                        .build();
            } else {
                Integer qty = existingCart.getQty() + request.getQty();
                existingCart.setQty(qty);
                existingCart.setPrice(qty * product.getPrice());
                existingCart.setUpdatedAt(now);
            }

            existingCart.setUpdatedAt(now);
            // TODO: Move This to sink connector
//            this.cartMongoDBRepository.save(existingCart);
            Model.CartItem ci = CartMapper.toProto(existingCart);

            this.kafkaTemplateSinkMongoCart.send("sink-mongo-commerce-carts", ci.getId(), ci);
            return AddToCartResponse
                    .newBuilder()
                    .setStatus("success")
                    .setMessage("addToCart")
                    .setData(AddToCartResponse.AddToCartResponseData.newBuilder().setId(existingCart.getId()).build())
                    .build();
        } catch (Exception ex) {
            span.recordException(ex);
            throw new Exception(ex);
        } finally {
            span.end();
        }
    }

    public FindCartItemsWithPaginationResponse findCartItemsWithPagination(FindCartItemsWithPaginationRequest request) {
        Span span = this.tracer.spanBuilder("CartGrpcService.findCartItemsWithPagination").startSpan();
        try (Scope scope = span.makeCurrent()) {
            int page = Math.max(request.getPage() - 1, 0);
            PageRequest pageRequest = PageRequest
                    .of(page, request.getLimit());

            Page<CartModelMongoDB> cartPage = this
                    .cartMongoDBRepository
                    .findAllWithPagination(request.getUserId(), pageRequest);

            FindCartItemsWithPaginationResponse.FindCartItemsWithPaginationResponseData.Builder responseData = FindCartItemsWithPaginationResponse
                    .FindCartItemsWithPaginationResponseData
                    .newBuilder()
                    .setLimit(cartPage.getSize())
                    .setPage(cartPage.getNumber() + 1)
                    .setTotal(Math.toIntExact(cartPage.getTotalElements()));

            cartPage
                    .stream()
                    .forEach(c -> responseData.addItems(CartMapper.toProto(c)));

            return FindCartItemsWithPaginationResponse
                    .newBuilder()
                    .setStatus("success")
                    .setMessage("FindCartItemsWithPagination")
                    .setData(responseData.build())
                    .build();

        } catch (Exception ex) {
            span.recordException(ex);
            return FindCartItemsWithPaginationResponse
                    .newBuilder()
                    .setStatus("failure")
                    .setMessage("FindCartItemsWithPagination")
                    .build();
        } finally {
            span.end();
        }
    }

    public DeleteCartItemByIdResponse deleteCartItemById(DeleteCartItemByIdRequest request) {
        Span span = this.tracer.spanBuilder("CartGrpcService.deleteCartItemById").startSpan();
        try (Scope scope = span.makeCurrent()) {
            // TODO:
            // 1. Validate in DB
            // 2. Delete in DB Event
            // 3. Delete in DB Commerce

            this.cartMongoDBRepository.deleteById(request.getId());
            return DeleteCartItemByIdResponse
                    .newBuilder()
                    .setStatus("success")
                    .setMessage("deleteCartItemById")
                    .build();
        } catch (Exception ex) {
            span.recordException(ex);
            return DeleteCartItemByIdResponse
                    .newBuilder()
                    .setStatus("failure")
                    .setMessage("deleteCartItemById")
                    .build();
        } finally {
            span.end();
        }
    }


}
