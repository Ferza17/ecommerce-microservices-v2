package com.ferza17.ecommercemicroservicesv2.commerceservice.model.mapper;

import com.ferza17.ecommercemicroservicesv2.commerceservice.model.mongodb.CartModelMongoDB;
import com.ferza17.ecommercemicroservicesv2.proto.v1.commerce.Model;
import com.ferza17.ecommercemicroservicesv2.commerceservice.util.ConvertTime;
import java.util.List;
import java.util.stream.Collectors;

public class CartMapper {
    public static CartModelMongoDB fromProto(Model.CartItem proto) {
        return CartModelMongoDB
                .builder()
                .id(proto.getId())
                .productId(proto.getProductId())
                .userId(proto.getUserId())
                .qty(proto.getQty())
                .price(proto.getPrice())
                .createdAt(ConvertTime.toInstant(proto.getCratedAt()))
                .updatedAt(ConvertTime.toInstant(proto.getUpdatedAt()))
                .build();
    }

    public static List<CartModelMongoDB> fromProtoList(List<Model.CartItem> protos) {
        return protos
                .stream()
                .map(CartMapper::fromProto)
                .collect(Collectors.toList());
    }

    public static Model.CartItem toProto(CartModelMongoDB entity) {
        return Model
                .CartItem
                .newBuilder()
                .setId(entity.getId())
                .setProductId(entity.getProductId())
                .setUserId(entity.getUserId())
                .setQty(entity.getQty())
                .setPrice(entity.getPrice())
                .setCratedAt(ConvertTime.toTimestamp(entity.getCreatedAt()))
                .setUpdatedAt(ConvertTime.toTimestamp(entity.getUpdatedAt()))
                .build();
    }

    public static List<Model.CartItem> toProtoList(List<CartModelMongoDB> entities) {
        return entities
                .stream()
                .map(CartMapper::toProto)
                .collect(Collectors.toList());
    }
}
