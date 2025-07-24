pub mod database {
    pub mod async_postgres;
    pub mod redis;
}

pub mod services {
    pub mod user;
}

pub mod message_broker {
    pub mod rabbitmq;
}

pub mod telemetry {
    pub mod jaeger;
}
