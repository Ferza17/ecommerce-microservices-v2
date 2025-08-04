use crate::infrastructure::services::payment::PaymentServiceGrpcClient;
use crate::infrastructure::services::user::UserServiceGrpcClient;
use crate::model::diesel::shipping_providers::to_proto::shipping_provider_to_proto;
use crate::model::diesel::shippings::to_proto::shippings_to_proto;
use crate::model::diesel::shippings::{CreateShippings, UpdateShippings};
use crate::model::rpc::payment::FindPaymentByIdRequest;
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
use prost_wkt_types::Timestamp;
use tonic::{Request, Response, Status};
use tracing::{Level, event, instrument};

pub trait ShippingUseCase {
    async fn create_shipping(
        &self,
        request_id: String,
        token: String,
        request: Request<CreateShippingRequest>,
    ) -> Result<Response<CreateShippingResponse>, Status>;
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
    shipping_repository: ShippingPostgresRepositoryImpl,
    shipping_provider_repository: ShippingProviderPostgresRepositoryImpl,
    user_service: UserServiceGrpcClient,
    payment_service: PaymentServiceGrpcClient,
}

impl ShippingUseCaseImpl {
    pub fn new(
        shipping_repository: ShippingPostgresRepositoryImpl,
        shipping_provider_repository: ShippingProviderPostgresRepositoryImpl,
        user_service: UserServiceGrpcClient,
        payment_service: PaymentServiceGrpcClient,
    ) -> Self {
        Self {
            shipping_repository,
            shipping_provider_repository,
            user_service,
            payment_service,
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
                if e.to_string().contains("not found") {
                    return Status::not_found("payment id not found".to_string());
                }
                Status::internal("error".to_string())
            })?;

        // VALIDATE PAYMENT_ID
        self
            .payment_service
            .clone()
            .find_payment_by_id(
                request_id.clone(),
                token.clone(),
                FindPaymentByIdRequest {
                    id: request.get_ref().payment_id.clone(),
                },
            )
            .await
            .map_err(|e| {
                eprintln!("find_payment_by_id {:?}", e);
                event!(name: "ShippingUseCase.create_shipping.error", Level::ERROR, request_id = request_id, error = ?e);
                if e.to_string().contains("payment id not found") {
                    return Status::not_found("not found".to_string());
                }
                Status::internal("error".to_string())
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
                eprintln!("get_shipping_provider_by_id {:?}", e);
                event!(name: "ShippingUseCase.create_shipping.error", Level::ERROR, request_id = request_id, error = ?e);
                if e.to_string().contains("shipping provider id not found") {
                    return Status::not_found("not found".to_string());
                }
                Status::internal("error".to_string())
            })?;

        // CREATE SHIPPING
        let now = chrono::Utc::now().naive_utc();
        self.shipping_repository.clone().create_shipping(
            &*request_id.clone(),
            &CreateShippings {
                id: uuid::Uuid::new_v4().to_string(),
                user_id: request.get_ref().user_id.clone(),
                payment_id: request.get_ref().payment_id.clone(),
                shipping_provider_id: request.get_ref().shipping_provider_id.clone(),
                created_at: now.clone(),
                updated_at: now,
                discarded_at: None,
            },
        ).await.map_err(|e| {
            eprintln!("create_shipping {:?}", e);
            event!(name: "ShippingUseCase.create_shipping.error", Level::ERROR, request_id = request_id, error = ?e);
            if e.to_string().contains("shipping provider id not found") {
                return Status::not_found("not found".to_string());
            }
            Status::internal("error".to_string())
        })?;

        Ok(Response::new(CreateShippingResponse {
            message: "create_shipping".to_string(),
            status: "success".to_string(),
            data: None,
        }))
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
                if e.to_string().contains("not found") {
                    return Status::not_found("not found".to_string());
                }
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
                if e.to_string().contains("not found") {
                    return Status::not_found("shipping not found".to_string());
                }
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
        let fetch_shippings = self.shipping_repository
            .clone()
            .list_shipping(&*request_id, &request.get_ref().page, &request.get_ref().limit)
            .await
            .map_err(|e| {
                event!(name: "ShippingUseCase.fetch_shipping.error", Level::ERROR, request_id = request_id, error = ?e);
                if e.to_string().contains("not found") {
                    return Status::not_found("shippings not found".to_string());
                }
                Status::internal("error".to_string())
            })?;

        Ok(Response::new(ListShippingResponse {
            message: "update_shipping".to_string(),
            status: "success".to_string(),
            data: shippings_to_proto(fetch_shippings),
        }))
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

        let now = chrono::Utc::now().naive_utc();
        self.shipping_repository.clone().update_shipping(
            &*request_id.clone(),
            &*request.get_ref().id.clone(),
            &UpdateShippings {
                user_id: Option::from(shipping.user_id),
                payment_id: Option::from(shipping.payment_id),
                shipping_provider_id: Option::from(shipping.shipping_provider_id),
                created_at: Option::from(now),
                updated_at: Option::from(now),
                discarded_at: None,
            },
        ).await.map_err(|e| {
            event!(name: "ShippingUseCase.delete_shipping.error", Level::ERROR, request_id = request_id, error = ?e);
            Status::internal("error".to_string())
        })?;

        Ok(Response::new(UpdateShippingResponse {
            message: "update_shipping".to_string(),
            status: "success".to_string(),
            data: None,
        }))
    }

    #[instrument("ShippingUseCase.delete_shipping")]
    async fn delete_shipping(
        &self,
        request_id: String,
        request: Request<DeleteShippingRequest>,
    ) -> Result<Response<DeleteShippingResponse>, Status> {
        self.shipping_repository
            .clone()
            .get_shipping_by_id(&*request_id, &*request.get_ref().id.clone())
            .await
            .map_err(|e| {
                event!(name: "ShippingUseCase.delete_shipping.error", Level::ERROR, request_id = request_id, error = ?e);
                Status::internal("error".to_string())
            })?;

        self.shipping_repository
            .clone()
            .delete_shipping(&*request_id.clone(), &*request.get_ref().id.clone())
            .await
            .map_err(|e| {
                event!(name: "ShippingUseCase.delete_shipping.error", Level::ERROR, request_id = request_id, error = ?e);
                Status::internal("error".to_string())
            })?;

        Ok(Response::new(DeleteShippingResponse {
            message: "delete_shipping".to_string(),
            status: "success".to_string(),
            data: None,
        }))
    }
}
