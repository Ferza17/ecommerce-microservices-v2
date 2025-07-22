use crate::infrastructure::database::postgres::PostgresInfrastructure;

#[derive(Clone)]
pub struct ShippingProviderPostgresRepository {
    pg: PostgresInfrastructure,
}

impl ShippingProviderPostgresRepository {
    pub fn new(pg: PostgresInfrastructure) -> Self {
        Self { pg }
    }
}
