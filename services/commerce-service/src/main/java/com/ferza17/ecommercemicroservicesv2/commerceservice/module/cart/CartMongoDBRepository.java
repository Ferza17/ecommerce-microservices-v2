package com.ferza17.ecommercemicroservicesv2.commerceservice.module.cart;

import com.ferza17.ecommercemicroservicesv2.commerceservice.model.mongodb.CartModelMongoDB;
import org.springframework.data.mongodb.repository.MongoRepository;
import org.springframework.data.mongodb.repository.Query;
import org.springframework.stereotype.Repository;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.Pageable;


import java.util.Optional;

@Repository
public interface CartMongoDBRepository extends MongoRepository<CartModelMongoDB, String> {
    @Query("{ 'product_id': ?0, 'user_id': ?1 }")
    Optional<CartModelMongoDB> findByProductIdAndUserId(String productId, String userId);

    @Query("{ 'user_id': ?0 }")
    Page<CartModelMongoDB> findAllWithPagination(String user_id, Pageable pageable);
}
