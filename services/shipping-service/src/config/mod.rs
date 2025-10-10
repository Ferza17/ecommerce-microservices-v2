pub mod config;
mod database_postgres;

mod message_broker_kafka;
mod message_broker_kafka_topic_sink_shipping;

mod service_payment;
mod service_shipping;
mod service_user;
mod telemetry_jaeger;
mod message_broker_kafka_topic_shipping;
mod database_mongodb;
