package com.ferza17.ecommercemicroservicesv2.commerceservice.module.cart;

import com.ferza17.ecommercemicroservicesv2.proto.v1.commerce.Response;
import com.ferza17.ecommercemicroservicesv2.proto.v1.commerce.Request;
import io.opentelemetry.api.trace.Span;
import io.opentelemetry.api.trace.Tracer;
import io.opentelemetry.context.Scope;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

@org.springframework.web.bind.annotation.RestController("cartRestController")
@RequestMapping("/api/v1/commerce/carts")
public class CartRestController {
    @Autowired
    private CartUseCase cartUseCase;
    @Autowired
    private Tracer tracer;

    @PostMapping("/items")
    public ResponseEntity<Response.AddToCartResponse> addToCart(@RequestBody Request.AddToCartRequest request) {
        Span span = this.tracer.spanBuilder("CartRestController.addToCart").startSpan();
        try (Scope scope = span.makeCurrent()) {
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
        Span span = this.tracer.spanBuilder("CartRestController.FindCartItemsWithPaginationResponse").startSpan();
        try (Scope scope = span.makeCurrent()) {
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
        Span span = this.tracer.spanBuilder("CartRestController.deleteCartItemById").startSpan();
        try (Scope scope = span.makeCurrent()) {
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
