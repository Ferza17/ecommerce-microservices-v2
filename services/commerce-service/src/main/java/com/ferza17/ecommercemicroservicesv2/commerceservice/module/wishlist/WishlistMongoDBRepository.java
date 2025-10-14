package com.ferza17.ecommercemicroservicesv2.commerceservice.module.wishlist;

import com.ferza17.ecommercemicroservicesv2.commerceservice.model.mongodb.WishlistModelMongoDB;
import org.springframework.data.mongodb.repository.MongoRepository;

public interface WishlistMongoDBRepository extends MongoRepository<WishlistModelMongoDB, String> {
}
