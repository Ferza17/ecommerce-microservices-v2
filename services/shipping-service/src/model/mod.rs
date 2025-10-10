pub mod diesel;
pub mod rpc {
    pub mod event;
    pub mod payment;
    pub mod shipping;
    pub mod user;

    pub mod response;
}

pub mod schema_registry {
    pub mod registry;
    mod shipping;
    mod shipping_providers;
}
