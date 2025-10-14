package com.ferza17.ecommercemicroservicesv2.commerceservice.module.wishlist;

import com.ferza17.ecommercemicroservicesv2.commerceservice.model.mongodb.WishlistModelMongoDB;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.Pageable;
import org.springframework.data.mongodb.repository.MongoRepository;
import org.springframework.data.mongodb.repository.Query;

public interface WishlistMongoDBRepository extends MongoRepository<WishlistModelMongoDB, String> {
    @Query("{'user_id': ?0}")
    Page<WishlistModelMongoDB> findAllWithPagination(String query, Pageable pageable);
}
