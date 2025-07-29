/* @generated and managed by dsync */

#[allow(unused)]
use crate::model::diesel::shipping_providers::ShippingProviders;
use crate::model::diesel::schema::*;

/// Struct representing a row in table `shippings`
#[derive(Debug, Clone, serde::Serialize, serde::Deserialize, diesel::Queryable, diesel::Selectable, diesel::QueryableByName, diesel::Associations, diesel::Identifiable)]
#[diesel(table_name=shippings, primary_key(id), belongs_to(ShippingProviders, foreign_key=shipping_provider_id))]
pub struct Shippings {
    /// Field representing column `id`
    pub id: String,
    /// Field representing column `user_id`
    pub user_id: String,
    /// Field representing column `payment_id`
    pub payment_id: String,
    /// Field representing column `shipping_provider_id`
    pub shipping_provider_id: String,
    /// Field representing column `created_at`
    pub created_at: chrono::NaiveDateTime,
    /// Field representing column `updated_at`
    pub updated_at: chrono::NaiveDateTime,
    /// Field representing column `discarded_at`
    pub discarded_at: Option<chrono::NaiveDateTime>,
}

/// Create Struct for a row in table `shippings` for [`Shippings`]
#[derive(Debug, Clone, serde::Serialize, serde::Deserialize, diesel::Insertable)]
#[diesel(table_name=shippings)]
pub struct CreateShippings {
    /// Field representing column `id`
    pub id: String,
    /// Field representing column `user_id`
    pub user_id: String,
    /// Field representing column `payment_id`
    pub payment_id: String,
    /// Field representing column `shipping_provider_id`
    pub shipping_provider_id: String,
    /// Field representing column `created_at`
    pub created_at: chrono::NaiveDateTime,
    /// Field representing column `updated_at`
    pub updated_at: chrono::NaiveDateTime,
    /// Field representing column `discarded_at`
    pub discarded_at: Option<chrono::NaiveDateTime>,
}

/// Update Struct for a row in table `shippings` for [`Shippings`]
#[derive(Debug, Clone, serde::Serialize, serde::Deserialize, diesel::AsChangeset, PartialEq, Default)]
#[diesel(table_name=shippings)]
pub struct UpdateShippings {
    /// Field representing column `user_id`
    pub user_id: Option<String>,
    /// Field representing column `payment_id`
    pub payment_id: Option<String>,
    /// Field representing column `shipping_provider_id`
    pub shipping_provider_id: Option<String>,
    /// Field representing column `created_at`
    pub created_at: Option<chrono::NaiveDateTime>,
    /// Field representing column `updated_at`
    pub updated_at: Option<chrono::NaiveDateTime>,
    /// Field representing column `discarded_at`
    pub discarded_at: Option<Option<chrono::NaiveDateTime>>,
}
