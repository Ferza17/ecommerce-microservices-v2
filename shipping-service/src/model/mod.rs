pub mod diesel {
    pub mod schema;
    pub mod orm {
        pub mod shipping;
        pub mod shipping_provider;
    }
}
pub mod rpc {
    pub mod payment;
    pub mod shipping;
    pub mod user;
    
    pub mod response;
}
