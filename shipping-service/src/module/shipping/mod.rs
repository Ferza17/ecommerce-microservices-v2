pub mod consumer_rabbitmq;
pub mod presenter;
pub mod repository_postgres;
pub mod repository_redis;
pub mod usecase;

// REPOSITORY
pub trait ShippingPostgresRepo: Send + Sync {}

pub trait ShippingRedisRepo: Send + Sync {}

pub trait ShippingUseCase: Send + Sync {}

pub trait ShippingPresenter: Send + Sync {}

pub trait ShippingRabbitMQConsumer: Send + Sync {}
