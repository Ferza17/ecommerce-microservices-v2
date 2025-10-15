package com.ferza17.ecommercemicroservicesv2.commerceservice.module.cart;

import com.ferza17.ecommercemicroservicesv2.proto.v1.commerce.Response;
import com.ferza17.ecommercemicroservicesv2.proto.v1.commerce.Request;
import com.ferza17.ecommercemicroservicesv2.proto.v1.commerce.Model;


import io.opentelemetry.api.trace.Span;
import io.opentelemetry.sdk.OpenTelemetrySdk;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

@org.springframework.web.bind.annotation.RestController("cartRestController")
@RequestMapping("/api/v1/commerce/carts")
public class CartRestController {
    private final CartUseCase cartUseCase;
    private final OpenTelemetrySdk openTelemetrySdk;
    public CartRestController(CartUseCase cartUseCase, OpenTelemetrySdk openTelemetrySdk) {
        this.cartUseCase = cartUseCase;
        this.openTelemetrySdk = openTelemetrySdk;
    }


    @PostMapping("/items")
    public ResponseEntity<Response.AddToCartResponse> addToCart(@RequestBody Request.AddToCartRequest request) {
        Span span = this.openTelemetrySdk.getTracer(CartGrpcService.class.getSimpleName()).spanBuilder("addToCart").startSpan();
        try {
            Response.AddToCartResponse response = this.cartUseCase.addToCart(request);
            return ResponseEntity.ok(response);
        } catch (Exception ex) {
            span.recordException(ex);
            return ResponseEntity.internalServerError().build();
        } finally {
            span.end();
        }
    }

    @GetMapping("/items")
    public ResponseEntity<Response.FindCartItemsWithPaginationResponse> findCartItemsWithPagination(@PathVariable("page") int page, @PathVariable("limit") int limit) {
        Span span = this.openTelemetrySdk.getTracer(CartGrpcService.class.getSimpleName()).spanBuilder("findCartItemsWithPagination").startSpan();
        try {
            Response.FindCartItemsWithPaginationResponse response = this.cartUseCase.findCartItemsWithPagination(Request.FindCartItemsWithPaginationRequest.newBuilder().setPage(page).setLimit(limit).build());
            return ResponseEntity.ok(response);
        } catch (Exception ex) {
            span.recordException(ex);
            return ResponseEntity.internalServerError().build();
        } finally {
            span.end();
        }
    }

    @DeleteMapping("/items/{id}")
    public ResponseEntity<Response.DeleteCartItemByIdResponse> deleteCartItemById(@PathVariable String id) {
        Span span = this.openTelemetrySdk.getTracer(CartGrpcService.class.getSimpleName()).spanBuilder("deleteCartItemById").startSpan();
        try {
            Response.DeleteCartItemByIdResponse response = this.cartUseCase.deleteCartItemById(Request.DeleteCartItemByIdRequest.newBuilder().setId(id).build());
            return ResponseEntity.ok(response);
        } catch (Exception ex) {
            span.recordException(ex);
            return ResponseEntity.internalServerError().build();
        } finally {
            span.end();
        }
    }
}
