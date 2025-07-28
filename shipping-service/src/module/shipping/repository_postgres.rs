use crate::infrastructure::database::async_postgres::AsyncPgDeadPool;
use std::fmt;

pub trait ShippingPostgresRepository {}

#[derive(Clone)]
pub struct ShippingPostgresRepositoryImpl {
    pg: AsyncPgDeadPool,
}

impl ShippingPostgresRepositoryImpl {
    pub fn new(pg: AsyncPgDeadPool) -> Self {
        Self { pg }
    }
}

impl fmt::Debug for ShippingPostgresRepositoryImpl {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        f.debug_struct("ShippingPostgresRepository")
            .field("pg", &"AsyncPgDeadPool")
            .finish()
    }
}

impl ShippingPostgresRepository for ShippingPostgresRepositoryImpl {
    
}