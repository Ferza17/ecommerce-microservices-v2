use crate::model::diesel::schema::shipping_providers;
use crate::model::rpc::shipping::ShippingProvider as ProtoShippingProvider;
use chrono::NaiveDateTime;
use diesel::prelude::*;
use prost_wkt_types::Timestamp;

#[derive(Queryable, Selectable)]
#[diesel(table_name = shipping_providers)]
#[diesel(check_for_backend(diesel::pg::Pg))]
pub struct ShippingProvider {
    pub id: String,
    pub name: String,
    pub created_at: NaiveDateTime,
    pub updated_at: NaiveDateTime,
    pub discarded_at: Option<NaiveDateTime>,
}

// Convert Model to Proto
pub fn shipping_provider_to_proto(shipping_provider: ShippingProvider) -> ProtoShippingProvider {
    ProtoShippingProvider {
        id: shipping_provider.id,
        name: shipping_provider.name,
        created_at: Option::from(Timestamp::from(shipping_provider.created_at)),
        updated_at: Option::from(Timestamp::from(shipping_provider.updated_at)),
        discarded_at: shipping_provider.discarded_at.map(|dt| Timestamp::from(dt)),
    }
}

pub fn shipping_providers_to_proto(providers: Vec<ShippingProvider>) -> Vec<ProtoShippingProvider> {
    providers
        .into_iter()
        .map(shipping_provider_to_proto)
        .collect()
}
