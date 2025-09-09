// @generated automatically by Diesel CLI.

diesel::table! {
    shipping_providers (id) {
        id -> Varchar,
        name -> Varchar,
        created_at -> Timestamp,
        updated_at -> Timestamp,
        discarded_at -> Nullable<Timestamp>,
    }
}

diesel::table! {
    shippings (id) {
        #[max_length = 255]
        id -> Varchar,
        #[max_length = 255]
        user_id -> Varchar,
        #[max_length = 255]
        payment_id -> Varchar,
        shipping_provider_id -> Varchar,
        created_at -> Timestamp,
        updated_at -> Timestamp,
        discarded_at -> Nullable<Timestamp>,
    }
}

diesel::joinable!(shippings -> shipping_providers (shipping_provider_id));

diesel::allow_tables_to_appear_in_same_query!(
    shipping_providers,
    shippings,
);
