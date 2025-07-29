/* @generated and managed by dsync */

#[allow(unused)]
use crate::model::diesel::schema::*;

/// Struct representing a row in table `shipping_providers`
#[derive(Debug, Clone, serde::Serialize, serde::Deserialize, diesel::Queryable, diesel::Selectable, diesel::QueryableByName, diesel::Identifiable)]
#[diesel(table_name=shipping_providers, primary_key(id))]
pub struct ShippingProviders {
    /// Field representing column `id`
    pub id: String,
    /// Field representing column `name`
    pub name: String,
    /// Field representing column `created_at`
    pub created_at: chrono::NaiveDateTime,
    /// Field representing column `updated_at`
    pub updated_at: chrono::NaiveDateTime,
    /// Field representing column `discarded_at`
    pub discarded_at: Option<chrono::NaiveDateTime>,
}

/// Create Struct for a row in table `shipping_providers` for [`ShippingProviders`]
#[derive(Debug, Clone, serde::Serialize, serde::Deserialize, diesel::Insertable)]
#[diesel(table_name=shipping_providers)]
pub struct CreateShippingProviders {
    /// Field representing column `id`
    pub id: String,
    /// Field representing column `name`
    pub name: String,
    /// Field representing column `created_at`
    pub created_at: chrono::NaiveDateTime,
    /// Field representing column `updated_at`
    pub updated_at: chrono::NaiveDateTime,
    /// Field representing column `discarded_at`
    pub discarded_at: Option<chrono::NaiveDateTime>,
}

/// Update Struct for a row in table `shipping_providers` for [`ShippingProviders`]
#[derive(Debug, Clone, serde::Serialize, serde::Deserialize, diesel::AsChangeset, PartialEq, Default)]
#[diesel(table_name=shipping_providers)]
pub struct UpdateShippingProviders {
    /// Field representing column `name`
    pub name: Option<String>,
    /// Field representing column `created_at`
    pub created_at: Option<chrono::NaiveDateTime>,
    /// Field representing column `updated_at`
    pub updated_at: Option<chrono::NaiveDateTime>,
    /// Field representing column `discarded_at`
    pub discarded_at: Option<Option<chrono::NaiveDateTime>>,
}
