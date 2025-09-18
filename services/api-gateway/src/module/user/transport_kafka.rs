use crate::config::config::AppConfig;
use crate::infrastructure::message_broker::kafka::KafkaInfrastructure;
use crate::model::rpc::user::AuthUserLoginByEmailAndPasswordRequest;
use crate::package::context::request_id::X_REQUEST_ID_HEADER;
use crate::util::metadata::inject_trace_context_to_kafka_headers;
use rdkafka::message::{Header, OwnedHeaders};
use rdkafka::producer::FutureRecord;
use tracing::{Span, error, instrument};
use tracing_opentelemetry::OpenTelemetrySpanExt;

#[derive(Debug, Clone)]
pub struct Transport {
    app_config: AppConfig,
    kafka_infrastructure: KafkaInfrastructure,
}

impl Transport {
    pub fn new(app_config: AppConfig, kafka_infrastructure: KafkaInfrastructure) -> Self {
        Self {
            app_config,
            kafka_infrastructure,
        }
    }

    #[instrument("user.transport_kafka.send_snapshot_users_user_login")]
    pub async fn send_snapshot_users_user_login(
        &self,
        request_id: String,
        request: tonic::Request<AuthUserLoginByEmailAndPasswordRequest>,
    ) -> Result<(), tonic::Status> {
        let headers =
            inject_trace_context_to_kafka_headers(OwnedHeaders::new(), &Span::current().context())
                .insert(Header {
                    key: X_REQUEST_ID_HEADER,
                    value: Some(request_id.clone().as_bytes()),
                });

        match self
            .kafka_infrastructure
            .publish(
                FutureRecord::to(
                    self.app_config
                        .service_user_kafka
                        .topic_snapshot_users_user_login
                        .as_str(),
                )
                .key(&request_id)
                .headers(headers)
                .payload(serde_json::to_string(request.get_ref()).unwrap().as_bytes()),
            )
            .await
        {
            Ok(_) => Ok(()),
            Err(e) => {
                error!("error: {}", e);
                Err(tonic::Status::internal("message not published"))
            }
        }
    }
}
