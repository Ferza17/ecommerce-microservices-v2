package com.ferza17.ecommercemicroservicesv2.commerceservice.module.wishlist;

import com.ferza17.ecommercemicroservicesv2.proto.v1.commerce.Request;
import com.ferza17.ecommercemicroservicesv2.proto.v1.commerce.Response;
import io.opentelemetry.api.trace.Span;
import io.opentelemetry.api.trace.Tracer;
import io.opentelemetry.context.Scope;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import java.util.ArrayList;

@org.springframework.web.bind.annotation.RestController
@RequestMapping("/api/v1/commerce/wishlists")
public class WishlistRestController {
    @Autowired
    private WishlistUseCase wishlistUseCase;
    @Autowired
    private Tracer tracer;

    @PostMapping("/items")
    public ResponseEntity<Response.AddToWishlistResponse> addToWishlist(@RequestBody Request.AddToWishlistRequest request) {
        Span span = this.tracer.spanBuilder("WishlistRestController.addToWishlist").startSpan();
        try (Scope scope = span.makeCurrent()) {
            // TODO: Implement Me
            return null;
        } catch (Exception e) {
            span.recordException(e);
            return null;
        } finally {
            span.end();
        }
    }

    @GetMapping("/items")
    public ResponseEntity<Response.FindWishlistItemWithPaginationResponse> findWishlistItemWithPagination(
            @PathVariable("user_id") String userId,
            @PathVariable("product_ids") ArrayList<String> productIds,
            @PathVariable("page") int page,
            @PathVariable("limit") int limit) {
        Span span = this.tracer.spanBuilder("WishlistRestController.findWishlistItemWithPagination").startSpan();
        try (Scope scope = span.makeCurrent()) {
            // TODO: Implement Me
            return null;
        } catch (Exception e) {
            span.recordException(e);
            return null;
        } finally {
            span.end();
        }
    }

    @DeleteMapping("/items/{id}")
    public ResponseEntity<Response.DeleteWishlistItemByIdResponse> deleteCartItemById(@PathVariable("id") String id) {
        Span span = this.tracer.spanBuilder("WishlistRestController.deleteCartItemById").startSpan();
        try (Scope scope = span.makeCurrent()) {
            // TODO: Implement Me
            return null;
        } catch (Exception e) {
            span.recordException(e);
            return null;
        } finally {
            span.end();
        }
    }

}
