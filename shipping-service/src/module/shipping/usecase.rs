use crate::module::shipping::repository_postgres::{
    ShippingPostgresRepository, ShippingPostgresRepositoryImpl,
};

pub trait ShippingUseCase {}

#[derive(Clone, Debug)]
pub struct ShippingUseCaseImpl {
    shipping_repository: ShippingPostgresRepositoryImpl,
}

impl ShippingUseCaseImpl {
    pub fn new(shipping_repository: ShippingPostgresRepositoryImpl) -> Self {
        Self {
            shipping_repository,
        }
    }
}

impl ShippingUseCase for ShippingUseCaseImpl {}
