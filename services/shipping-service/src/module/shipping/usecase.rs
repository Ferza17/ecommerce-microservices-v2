use crate::config::config::AppConfig;
use crate::infrastructure::message_broker::kafka::KafkaInfrastructure;
use crate::infrastructure::services::payment::PaymentServiceGrpcClient;
use crate::infrastructure::services::user::UserServiceGrpcClient;
use crate::model::diesel::shipping_providers::to_proto::shipping_provider_to_proto;
use crate::model::diesel::shippings::to_proto::shippings_to_proto;
use crate::model::diesel::shippings::{CreateShippings};
use crate::model::rpc::event::ReserveEvent;
use crate::model::rpc::payment::FindPaymentByIdRequest;
use crate::model::rpc::shipping::create_shipping_response::CreateShippingResponseData;
use crate::model::rpc::shipping::delete_shipping_response::DeleteShippingResponseData;
use crate::model::rpc::shipping::update_shipping_response::UpdateShippingResponseData;
use crate::model::rpc::shipping::{
    CreateShippingRequest, CreateShippingResponse, DeleteShippingRequest, DeleteShippingResponse,
    GetShippingByIdRequest, GetShippingByIdResponse, ListShippingRequest, ListShippingResponse,
    ShippingFullResponse, UpdateShippingRequest, UpdateShippingResponse,
};
use crate::model::rpc::user::{AuthUserFindUserByTokenRequest, FindUserByIdRequest};
use crate::module::shipping::repository_postgres::{
    ShippingPostgresRepository, ShippingPostgresRepositoryImpl,
};
use crate::module::shipping_provider::repository_postgres::{
    ShippingProviderPostgresRepository, ShippingProviderPostgresRepositoryImpl,
};
use crate::package::context::auth::AUTHORIZATION_HEADER;
use crate::package::context::request_id::X_REQUEST_ID_HEADER;
use crate::util::metadata::inject_trace_context_to_kafka_headers;
use prost_wkt_types::Timestamp;
use rdkafka::message::OwnedHeaders;
use tonic::{Request, Response, Status};
use tracing::{Level, Span, event, instrument};
use tracing_opentelemetry::OpenTelemetrySpanExt;

pub trait ShippingUseCase {
    async fn create_shipping(
        &self,
        request_id: String,
        token: String,
        request: Request<CreateShippingRequest>,
    ) -> Result<Response<CreateShippingResponse>, Status>;

    async fn confirm_create_shipping(
        &self,
        request_id: String,
        token: String,
        request: Request<ReserveEvent>,
    ) -> Result<(), Status>;
    async fn compensate_create_shipping(
        &self,
        request_id: String,
        token: String,
        request: Request<ReserveEvent>,
    ) -> Result<(), Status>;

    async fn get_shipping_by_id(
        &self,
        request_id: String,
        token: String,
        request: Request<GetShippingByIdRequest>,
    ) -> Result<Response<GetShippingByIdResponse>, Status>;
    async fn list_shipping(
        &self,
        request_id: String,
        request: Request<ListShippingRequest>,
    ) -> Result<Response<ListShippingResponse>, Status>;
    async fn update_shipping(
        &self,
        request_id: String,
        token: String,
        request: Request<UpdateShippingRequest>,
    ) -> Result<Response<UpdateShippingResponse>, Status>;
    async fn delete_shipping(
        &self,
        request_id: String,
        request: Request<DeleteShippingRequest>,
    ) -> Result<Response<DeleteShippingResponse>, Status>;
}

#[derive(Clone, Debug)]
pub struct ShippingUseCaseImpl {
    app_config: AppConfig,
    shipping_repository: ShippingPostgresRepositoryImpl,
    shipping_provider_repository: ShippingProviderPostgresRepositoryImpl,
    user_service: UserServiceGrpcClient,
    payment_service: PaymentServiceGrpcClient,
    kafka_infrastructure: KafkaInfrastructure,
}

impl ShippingUseCaseImpl {
    pub fn new(
        app_config: AppConfig,
        shipping_repository: ShippingPostgresRepositoryImpl,
        shipping_provider_repository: ShippingProviderPostgresRepositoryImpl,
        user_service: UserServiceGrpcClient,
        payment_service: PaymentServiceGrpcClient,
        kafka_infrastructure: KafkaInfrastructure,
    ) -> Self {
        Self {
            app_config,
            shipping_repository,
            shipping_provider_repository,
            user_service,
            payment_service,
            kafka_infrastructure,
        }
    }
}

impl ShippingUseCase for ShippingUseCaseImpl {
    #[instrument("ShippingUseCase.create_shipping")]
    async fn create_shipping(
        &self,
        request_id: String,
        token: String,
        request: Request<CreateShippingRequest>,
    ) -> Result<Response<CreateShippingResponse>, Status> {
        // VALIDATE USER_ID
        self.user_service
            .clone()
            .find_user_by_id(
                request_id.clone(),
                token.clone(),
                tonic::Request::new(FindUserByIdRequest {
                    id: request.get_ref().user_id.clone(),
                }),
            ).await
            .map_err(|e| {
                eprintln!("find_user_by_id {:?}", e);
                event!(name: "ShippingUseCase.create_shipping.error", Level::ERROR, request_id = request_id, error = ?e);
                Status::from(e)
            })?;

        // VALIDATE SHIPPING_PROVIDER_ID
        self
            .shipping_provider_repository
            .clone()
            .get_shipping_provider_by_id(
                &*request_id.clone(),
                &*request.get_ref().shipping_provider_id.clone(),
            )
            .await
            .map_err(|e| {
                event!(name: "ShippingUseCase.create_shipping.error", Level::ERROR, request_id = request_id.clone(), error = ?e);
                Status::internal("error".to_string())
            })?;

        // CREATE SHIPPING
        let now = chrono::Utc::now();
        let id = uuid::Uuid::new_v4().to_string();

        // TODO:
        // 1. Move to Kafaka Sink Mongodb Shipping Event Stores
        // 2. Add Confirm function
        // 3. Add Compensate Function



        match self
            .kafka_infrastructure
            .publish_with_json_schema(
                self.app_config
                    .message_broker_kafka_topic_sink_shipping
                    .pg_shippings_shippings
                    .clone(),
                crate::model::schema_registry::registry::Registry::Shipping,
                serde_json::to_value(&CreateShippings {
                    id: id.clone(),
                    user_id: request.get_ref().user_id.clone(),
                    payment_id: request.get_ref().payment_id.clone(),
                    shipping_provider_id: request.get_ref().shipping_provider_id.clone(),
                    created_at: now.clone(),
                    updated_at: now,
                    discarded_at: None,
                })
                .unwrap(),
                id.clone(),
                Some(
                    inject_trace_context_to_kafka_headers(
                        OwnedHeaders::new(),
                        &Span::current().context(),
                    )
                    .insert(rdkafka::message::Header {
                        key: X_REQUEST_ID_HEADER,
                        value: Some(request_id.clone().as_bytes()),
                    })
                    .insert(rdkafka::message::Header {
                        key: AUTHORIZATION_HEADER,
                        value: Some(format!("Bearer {}", token).as_bytes()),
                    }),
                ),
            )
            .await
        {
            Ok(_) => {
                event!(name : "ShippingUseCase.create_shipping.success", Level::INFO, request_id = request_id);
                Ok(Response::new(CreateShippingResponse {
                    message: "create_shipping".to_string(),
                    status: "success".to_string(),
                    data: Some(CreateShippingResponseData { id }),
                }))
            }
            Err(e) => {
                event!(name: "ShippingUseCase.create_shipping.error", Level::ERROR, request_id = request_id, error = ?e);
                Err(Status::internal(format!(
                    "invalid to create shipping : {}",
                    e.to_string()
                )))
            }
        }
    }

    #[instrument("ShippingUseCase.confirm_create_shipping")]
    async fn confirm_create_shipping(
        &self,
        request_id: String,
        token: String,
        request: Request<ReserveEvent>,
    ) -> Result<(), Status>{
        Ok(())
    }

    #[instrument("ShippingUseCase.compensate_create_shipping")]
    async fn compensate_create_shipping(
        &self,
        request_id: String,
        token: String,
        request: Request<ReserveEvent>,
    ) -> Result<(), Status>{
        Ok(())
    }

    #[instrument("ShippingUseCase.get_shipping_by_id")]
    async fn get_shipping_by_id(
        &self,
        request_id: String,
        token: String,
        request: Request<GetShippingByIdRequest>,
    ) -> Result<Response<GetShippingByIdResponse>, Status> {
        event!(name: "ShippingUseCase.get_shipping_by_id", Level::INFO, request_id = request_id, request = ?request.get_ref());
        let fetch_shipping = self
            .clone()
            .shipping_repository
            .get_shipping_by_id(&*request_id.clone(), &*request.get_ref().id.clone())
            .await
            .map_err(|e| {
                event!(name: "ShippingUseCase.get_shipping_by_id.error", Level::ERROR, request_id = request_id, error = ?e);
                Status::internal("error".to_string())
            })?;

        let fetch_user = self
            .user_service
            .clone()
            .auth_user_find_user_by_token(request_id.clone(), tonic::Request::new(AuthUserFindUserByTokenRequest {
                token: token.clone(),
            }))
            .await
            .map_err(|e| {
                event!(name: "ShippingUseCase.get_shipping_by_id.error", Level::ERROR, request_id = request_id, error = ?e);
                e
            })?;

        let fetch_payment = self
            .payment_service
            .clone()
            .find_payment_by_id(
                request_id.clone(),
                token.clone(),
                FindPaymentByIdRequest {
                    id: fetch_shipping.payment_id.clone(),
                },
            )
            .await
            .map_err(|e| {
                event!(name: "ShippingUseCase.find_payment_by_id.error", Level::ERROR, request_id = request_id, error = ?e);
                e
            })?;

        let fetch_shipping_provider = self
            .shipping_provider_repository
            .clone()
            .get_shipping_provider_by_id(
                &*request_id.clone(),
                &*fetch_shipping.shipping_provider_id.clone(),
            )
            .await
            .map_err(|e| {
                event!(name: "ShippingUseCase.fetch_shipping_provider.error", Level::ERROR, request_id = request_id, error = ?e);
                Status::internal("error".to_string())
            })?;

        Ok(Response::new(GetShippingByIdResponse {
            message: "get_shipping_by_id".to_string(),
            status: "success".to_string(),
            data: Option::from(ShippingFullResponse {
                id: fetch_shipping.id,
                user: fetch_user.data.unwrap().user,
                payment: Option::from(fetch_payment.data.unwrap().payment),
                shipping_provider: Option::from(shipping_provider_to_proto(
                    fetch_shipping_provider,
                )),
                created_at: Option::from(Timestamp::from(fetch_shipping.created_at)),
                updated_at: Option::from(Timestamp::from(fetch_shipping.updated_at)),
                discarded_at: fetch_shipping.discarded_at.map(|dt| Timestamp::from(dt)),
            }),
        }))
    }

    #[instrument("ShippingUseCase.list_shipping")]
    async fn list_shipping(
        &self,
        request_id: String,
        request: Request<ListShippingRequest>,
    ) -> Result<Response<ListShippingResponse>, Status> {
        match self
            .shipping_repository
            .clone()
            .list_shipping(
                &*request_id,
                &request.get_ref().page,
                &request.get_ref().limit,
            )
            .await
        {
            Ok(v) => {
                event!(name: "ShippingUseCase.list_shipping.success", Level::INFO, request_id = request_id, shippings = ?v);
                Ok(Response::new(ListShippingResponse {
                    message: "list_shipping".to_string(),
                    status: "success".to_string(),
                    data: shippings_to_proto(v),
                }))
            }
            Err(err) => {
                event!(name: "ShippingUseCase.list_shipping.error", Level::ERROR, request_id = request_id, error = ?err);
                Err(Status::internal("error".to_string()))
            }
        }
    }

    #[instrument("ShippingUseCase.update_shipping")]
    async fn update_shipping(
        &self,
        request_id: String,
        token: String,
        request: Request<UpdateShippingRequest>,
    ) -> Result<Response<UpdateShippingResponse>, Status> {
        let mut shipping = self.shipping_repository
            .clone()
            .get_shipping_by_id(&*request_id, &*request.get_ref().id.clone())
            .await
            .map_err(|e| {
                event!(name: "ShippingUseCase.delete_shipping.error", Level::ERROR, request_id = request_id, error = ?e);
                Status::internal("error".to_string())
            })?;

        // Partial Update Changes Field
        if request.get_ref().shipping_provider_id != shipping.shipping_provider_id {
            shipping.shipping_provider_id = request.get_ref().shipping_provider_id.clone();
        }
        if request.get_ref().payment_id != shipping.payment_id {
            shipping.payment_id = request.get_ref().payment_id.clone();
        }
        if request.get_ref().user_id != shipping.user_id {
            shipping.user_id = request.get_ref().user_id.clone();
        }

        let now = chrono::Utc::now();
        shipping.updated_at = now;

        match self
            .kafka_infrastructure
            .publish_with_json_schema(
                self.app_config
                    .message_broker_kafka_topic_sink_shipping
                    .pg_shippings_shippings
                    .clone(),
                crate::model::schema_registry::registry::Registry::Shipping,
                serde_json::to_value(&shipping).unwrap(),
                shipping.id.clone(),
                Some(
                    inject_trace_context_to_kafka_headers(
                        OwnedHeaders::new(),
                        &Span::current().context(),
                    )
                    .insert(rdkafka::message::Header {
                        key: X_REQUEST_ID_HEADER,
                        value: Some(request_id.clone().as_bytes()),
                    })
                    .insert(rdkafka::message::Header {
                        key: AUTHORIZATION_HEADER,
                        value: Some(format!("Bearer {}", token).as_bytes()),
                    }),
                ),
            )
            .await
        {
            Ok(_) => {
                event!(name : "ShippingUseCase.update_shipping.success", Level::INFO, request_id = request_id, shipping = ?shipping);
                Ok(Response::new(UpdateShippingResponse {
                    message: "update_shipping".to_string(),
                    status: "success".to_string(),
                    data: Some(UpdateShippingResponseData { id: shipping.id }),
                }))
            }
            Err(e) => {
                event!(name: "ShippingUseCase.update_shipping.error", Level::ERROR, request_id = request_id, error = ?e);
                Err(Status::internal(format!(
                    "invalid to update shipping : {}",
                    e.to_string()
                )))
            }
        }
    }

    #[instrument("ShippingUseCase.delete_shipping")]
    async fn delete_shipping(
        &self,
        request_id: String,
        request: Request<DeleteShippingRequest>,
    ) -> Result<Response<DeleteShippingResponse>, Status> {
        let shipping = self.shipping_repository
            .clone()
            .get_shipping_by_id(&*request_id, &*request.get_ref().id.clone())
            .await
            .map_err(|e| {
                event!(name: "ShippingUseCase.delete_shipping.error", Level::ERROR, request_id = request_id, error = ?e);
                Status::internal("error".to_string())
            })?;

        let now = chrono::Utc::now();
        match self
            .kafka_infrastructure
            .publish_with_json_schema(
                self.app_config
                    .message_broker_kafka_topic_sink_shipping
                    .pg_shippings_shippings
                    .clone(),
                crate::model::schema_registry::registry::Registry::Shipping,
                serde_json::to_value(&CreateShippings {
                    id: shipping.id.clone(),
                    user_id: shipping.user_id,
                    payment_id: shipping.payment_id,
                    shipping_provider_id: shipping.shipping_provider_id,
                    created_at: shipping.created_at,
                    updated_at: shipping.created_at,
                    discarded_at: Some(now),
                })
                .unwrap(),
                shipping.id.clone(),
                Some(
                    inject_trace_context_to_kafka_headers(
                        OwnedHeaders::new(),
                        &Span::current().context(),
                    )
                    .insert(rdkafka::message::Header {
                        key: X_REQUEST_ID_HEADER,
                        value: Some(request_id.clone().as_bytes()),
                    }),
                ),
            )
            .await
        {
            Ok(_) => {
                event!(name : "ShippingUseCase.delete_shipping.success", Level::INFO, request_id = request_id);
                Ok(Response::new(DeleteShippingResponse {
                    message: "delete_shipping".to_string(),
                    status: "success".to_string(),
                    data: Some(DeleteShippingResponseData {
                        id: shipping.id.to_string(),
                    }),
                }))
            }
            Err(e) => {
                event!(name: "ShippingUseCase.delete_shipping.error", Level::ERROR, request_id = request_id, error = ?e);
                Err(Status::internal(format!(
                    "invalid to delete shipping : {}",
                    e.to_string()
                )))
            }
        }
    }
}
