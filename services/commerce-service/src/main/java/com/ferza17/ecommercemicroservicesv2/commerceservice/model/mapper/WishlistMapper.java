package com.ferza17.ecommercemicroservicesv2.commerceservice.model.mapper;

import com.ferza17.ecommercemicroservicesv2.commerceservice.model.mongodb.WishlistModelMongoDB;
import com.ferza17.ecommercemicroservicesv2.commerceservice.util.ConvertTime;
import com.ferza17.ecommercemicroservicesv2.proto.v1.commerce.Model;

import java.util.List;
import java.util.stream.Collectors;

public class WishlistMapper {
    public static WishlistModelMongoDB fromProto(Model.WishlistItem proto) {
        return WishlistModelMongoDB
                .builder()
                .id(proto.getId())
                .product_id(proto.getProductId())
                .user_id(proto.getUserId())
                .created_at(ConvertTime.toInstant(proto.getCratedAt()))
                .updated_at(ConvertTime.toInstant(proto.getUpdatedAt()))
                .build();
    }

    public static List<WishlistModelMongoDB> fromProtoList(List<Model.WishlistItem> protos) {
        return protos
                .stream()
                .map(WishlistMapper::fromProto)
                .collect(Collectors.toList());
    }


    public static Model.WishlistItem toProto(WishlistModelMongoDB entity) {
        return Model
                .WishlistItem
                .newBuilder()
                .setId(entity.getId())
                .setProductId(entity.getProduct_id())
                .setUserId(entity.getUser_id())
                .setCratedAt(ConvertTime.toTimestamp(entity.getCreated_at()))
                .setUpdatedAt(ConvertTime.toTimestamp(entity.getUpdated_at()))
                .build();
    }

    public static List<Model.WishlistItem> toProtoList(List<WishlistModelMongoDB> entities) {
        return entities
                .stream()
                .map(WishlistMapper::toProto)
                .collect(Collectors.toList());
    }

}
