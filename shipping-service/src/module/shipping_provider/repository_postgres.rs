use crate::infrastructure::database::postgres::PostgresPool;

pub struct ShippingProviderPostgresRepository {
    pg: PostgresPool,
}

impl ShippingProviderPostgresRepository {
    pub fn new(pg: PostgresPool) -> Self {
        Self { pg }
    }
}
