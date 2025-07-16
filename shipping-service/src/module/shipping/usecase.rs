use std::sync::Arc;

use crate::module::shipping::{ShippingPostgresRepo, ShippingRedisRepo};

pub struct ShippingUseCase {
    pub shipping_postgres_repo: Arc<dyn ShippingPostgresRepo>,
    pub shipping_redis_repo: Arc<dyn ShippingRedisRepo>,
}

impl ShippingUseCase {
    pub fn new(
        shipping_postgres_repo: Arc<dyn ShippingPostgresRepo>,
        shipping_redis_repo: Arc<dyn ShippingRedisRepo>,
    ) -> Self {
        Self {
            shipping_postgres_repo,
            shipping_redis_repo,
        }
    }
}
