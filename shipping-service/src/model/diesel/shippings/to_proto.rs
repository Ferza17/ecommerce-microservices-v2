use crate::model::diesel::shippings::Shippings as ShippingsModel;
use crate::model::rpc::shipping::Shipping as ProtoShipping;
use prost_wkt_types::Timestamp;

pub fn shipping_to_proto(shipping: ShippingsModel) -> ProtoShipping {
    ProtoShipping {
        id: shipping.id,
        user_id: shipping.user_id,
        payment_id: shipping.payment_id,
        shipping_provider_id: shipping.shipping_provider_id,
        created_at: Option::from(Timestamp::from(shipping.created_at)),
        updated_at: Option::from(Timestamp::from(shipping.updated_at)),
        discarded_at: shipping.discarded_at.map(|dt| Timestamp::from(dt)),
    }
}

pub fn shippings_to_proto(shippings: Vec<ShippingsModel>) -> Vec<ProtoShipping> {
    shippings.into_iter().map(shipping_to_proto).collect()
}
