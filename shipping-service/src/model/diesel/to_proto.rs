use crate::model::diesel::shipping_providers::ShippingProviders;
use crate::model::rpc::shipping::ShippingProvider as ProtoShippingProvider;
use prost_wkt_types::Timestamp;

// Convert Model to Proto
pub fn shipping_provider_to_proto(shipping_provider: ShippingProviders) -> ProtoShippingProvider {
    ProtoShippingProvider {
        id: shipping_provider.id,
        name: shipping_provider.name,
        created_at: Option::from(Timestamp::from(shipping_provider.created_at)),
        updated_at: Option::from(Timestamp::from(shipping_provider.updated_at)),
        discarded_at: shipping_provider.discarded_at.map(|dt| Timestamp::from(dt)),
    }
}

pub fn shipping_providers_to_proto(providers: Vec<ShippingProviders>) -> Vec<ProtoShippingProvider> {
    providers
        .into_iter()
        .map(shipping_provider_to_proto)
        .collect()
}