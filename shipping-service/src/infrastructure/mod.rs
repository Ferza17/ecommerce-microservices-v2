pub mod database {
    pub mod postgres;
    pub mod redis;
}

pub mod services {
    pub mod user;
}

pub mod message_broker {
    pub mod rabbitmq;
}