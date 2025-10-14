package com.ferza17.ecommercemicroservicesv2.commerceservice.model.mongodb;

import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;
import lombok.NoArgsConstructor;
import org.springframework.data.annotation.Id;
import org.springframework.data.mongodb.core.mapping.Document;
import org.springframework.data.mongodb.core.mapping.Field;

import java.time.Instant;

@Document(collection = "wishlists")
@Data
@NoArgsConstructor
@AllArgsConstructor
@Builder
public class WishlistModelMongoDB {
    @Id
    private String id;

    @Field("product_id")
    private String product_id;

    @Field("user_id")
    private String user_id;

    @Field("created_at")
    private Instant created_at;

    @Field("updated_at")
    private Instant updated_at;
}
