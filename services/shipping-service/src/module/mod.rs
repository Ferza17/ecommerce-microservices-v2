pub mod shipping_provider {
    pub mod presenter_grpc;
    pub mod presenter_http;
    pub mod repository_postgres;
    pub mod usecase;
}

pub mod shipping {

    pub mod consumer_rabbitmq;
    pub mod consumer_kafka;
    pub mod presenter_grpc;
    pub mod presenter_http;
    pub mod repository_postgres;
    pub mod usecase;
}
