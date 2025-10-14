package com.ferza17.ecommercemicroservicesv2.commerceservice.model.mongodb;

import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;
import lombok.NoArgsConstructor;
import org.springframework.data.annotation.Id;
import org.springframework.data.mongodb.core.mapping.Document;
import org.springframework.data.mongodb.core.mapping.Field;

import java.time.Instant;

@Document(collection = "carts")
@Data
@NoArgsConstructor
@AllArgsConstructor
@Builder
public class CartModelMongoDB {
    @Id
    private String id;

    @Field("product_id")
    private String productId;

    @Field("user_id")
    private String userId;

    @Field("qty")
    private Integer qty;

    @Field("price")
    private Double price;

    @Field("created_at")
    private Instant createdAt;

    @Field("updated_at")
    private Instant updatedAt;
}
