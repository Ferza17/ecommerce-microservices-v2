pub enum Registry {
    ShippingProvider,
    Shipping,
}

impl std::fmt::Display for Registry {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        match self {
            Registry::ShippingProvider => write!(
                f,
                "{}",
                crate::model::schema_registry::shipping_providers::SHIPPING_PROVIDERS_SCHEMA
            ),
            Registry::Shipping => {
                write!(f, "{}", crate::model::schema_registry::shipping::SHIPPING)
            }
        }
    }
}
