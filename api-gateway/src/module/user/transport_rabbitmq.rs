use crate::config::config::AppConfig;
use crate::infrastructure::message_broker::rabbitmq::RabbitMQInfrastructure;
use crate::model::rpc::user::{AuthUserLoginByEmailAndPasswordRequest, AuthUserRegisterRequest};

use crate::package::context::content_type::{APPLICATION_JSON, CONTENT_TYPE};
use crate::package::context::request_id::X_REQUEST_ID_HEADER;
use crate::util::metadata::inject_trace_context_to_lapin_table;
use lapin::BasicProperties;
use lapin::types::AMQPValue;
use tracing::{Span, error, info, instrument, span};
use tracing_opentelemetry::OpenTelemetrySpanExt;

#[derive(Debug, Clone)]
pub struct Transport {
    app_config: AppConfig,
    rabbitmq_infrastructure: RabbitMQInfrastructure,
}

impl Transport {
    pub fn new(app_config: AppConfig, rabbitmq_infrastructure: RabbitMQInfrastructure) -> Self {
        Self {
            app_config,
            rabbitmq_infrastructure,
        }
    }

    #[instrument("auth.transport_rabbitmq.publish_user_login")]
    pub async fn publish_user_login(
        &self,
        request_id: String,
        request: tonic::Request<AuthUserLoginByEmailAndPasswordRequest>,
    ) -> Result<(), tonic::Status> {
        let mut table = lapin::types::FieldTable::default();
        table.insert(
            X_REQUEST_ID_HEADER.into(),
            AMQPValue::LongString(request_id.clone().into()),
        );

        match self
            .rabbitmq_infrastructure
            .publish(
                &*self.app_config.service_user_rabbitmq.exchange_user,
                &*self.app_config.service_user_rabbitmq.queue_user_login,
                BasicProperties::default()
                    .with_headers(inject_trace_context_to_lapin_table(
                        table,
                        Span::current().context(),
                    ))
                    .with_delivery_mode(1) // TRANSIENT
                    .with_correlation_id(request_id.into())
                    .with_content_type(APPLICATION_JSON.into()),
                serde_json::to_string(request.get_ref()).unwrap().as_bytes(),
            )
            .await
        {
            Ok(_) => Ok(()),
            Err(e) => {
                error!(
                    "exchange:{} queue:{} Message not published: {}",
                    self.app_config.service_user_rabbitmq.exchange_user,
                    self.app_config.service_user_rabbitmq.queue_user_login,
                    e
                );
                Err(tonic::Status::internal("Message not published"))
            }
        }
    }

    pub async fn publish_user_created(
        &self,
        request_id: String,
        request: tonic::Request<AuthUserRegisterRequest>,
    ) -> Result<(), tonic::Status> {
        let mut table = lapin::types::FieldTable::default();
        table.insert(
            X_REQUEST_ID_HEADER.into(),
            AMQPValue::LongString(request_id.clone().into()),
        );
        match self
            .rabbitmq_infrastructure
            .publish(
                &*self.app_config.service_user_rabbitmq.exchange_user,
                &*self.app_config.service_user_rabbitmq.queue_user_created,
                BasicProperties::default()
                    .with_headers(inject_trace_context_to_lapin_table(
                        table,
                        Span::current().context(),
                    ))
                    .with_delivery_mode(1) // TRANSIENT
                    .with_correlation_id(request_id.into())
                    .with_content_type(APPLICATION_JSON.into()),
                serde_json::to_string(request.get_ref()).unwrap().as_bytes(),
            )
            .await
        {
            Ok(_) => Ok(()),
            Err(e) => {
                error!(
                    "exchange:{} queue:{} Message not published: {}",
                    self.app_config.service_user_rabbitmq.exchange_user,
                    self.app_config.service_user_rabbitmq.queue_user_created,
                    e
                );
                Err(tonic::Status::internal("Message not published"))
            }
        }
        
    }
}
