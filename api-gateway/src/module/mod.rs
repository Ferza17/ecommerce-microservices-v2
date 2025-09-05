pub mod auth {
    pub mod http_presenter;
    pub mod transport_grpc;
    pub mod transport_rabbitmq;
    pub mod usecase;
}

pub mod notification {
    pub mod http_presenter;
    pub mod usecase;
}

pub mod user {
    pub mod http_presenter;
    pub mod transport_grpc;
    pub mod transport_rabbitmq;
    pub mod usecase;
}

pub mod product {
    pub mod http_presenter;
    pub mod transport_grpc;
    pub mod transport_rabbitmq;
    pub mod usecase;
}

pub mod shipping {
    pub mod http_presenter;
    pub mod transport_grpc;
    pub mod usecase;
}

pub mod shipping_provider {
    pub mod http_presenter;
    pub mod transport_grpc;
    pub mod usecase;
}

pub mod payment_providers {
    pub mod http_presenter;
    pub mod transport_grpc;
    pub mod usecase;
}

pub mod payment {
    pub mod http_presenter;
    pub mod transport_grpc;
    pub mod transport_rabbitmq;
    pub mod usecase;
}

pub mod event {
    pub mod transport_grpc;
    pub mod transport_rabbitmq;
}
