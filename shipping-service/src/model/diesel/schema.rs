// @generated automatically by Diesel CLI.

diesel::table! {
    shipping_providers (id) {
        id -> Varchar,
        name -> Nullable<Varchar>,
        created_at -> Nullable<Timestamp>,
        updated_at -> Nullable<Timestamp>,
        discarded_at -> Nullable<Timestamp>,
    }
}

diesel::table! {
    shippings (id) {
        id -> Varchar,
        created_by_id -> Varchar,
        order_id -> Varchar,
        #[max_length = 15]
        status -> Varchar,
        shipping_provider_id -> Nullable<Varchar>,
        created_at -> Nullable<Timestamp>,
        updated_at -> Nullable<Timestamp>,
        discarded_at -> Nullable<Timestamp>,
    }
}

diesel::joinable!(shippings -> shipping_providers (shipping_provider_id));

diesel::allow_tables_to_appear_in_same_query!(
    shipping_providers,
    shippings,
);
