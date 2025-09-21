pub mod diesel;
pub mod rpc {
    pub mod payment;
    pub mod shipping;
    pub mod user;

    pub mod response;
}

pub mod schema_registry{
    pub mod shipping_providers;
}