use crate::config::config::AppConfig;
use crate::infrastructure::message_broker::kafka::KafkaInfrastructure;
use crate::model::rpc::user::{AuthUserLoginByEmailAndPasswordRequest, AuthUserRegisterRequest};
use crate::package::context::request_id::X_REQUEST_ID_HEADER;
use crate::util::metadata::inject_trace_context_to_kafka_headers;
use prost::Message;
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

    #[instrument("user.transport_kafka.send_snapshot")]
    pub async fn send_snapshot<M: Message>(
        &self,
        request_id: String,
        request: tonic::Request<M>,
        topic: &str,
    ) -> Result<(), tonic::Status> {
        let mut headers =
            inject_trace_context_to_kafka_headers(OwnedHeaders::new(), &Span::current().context());

        if !request_id.is_empty() {
            headers = headers.insert(Header {
                key: X_REQUEST_ID_HEADER,
                value: Some(request_id.as_bytes()),
            });
        }

        let mut buf = Vec::new();
        request
            .into_inner()
            .encode(&mut buf)
            .map_err(|err| tonic::Status::new(tonic::Code::InvalidArgument, format!("{}", err)))?;

        self.kafka_infrastructure
            .publish(
                FutureRecord::to(topic)
                    .key(&request_id)
                    .headers(headers)
                    .payload(&buf),
            )
            .await
            .map_err(|e| {
                error!("error: {}", e);
                tonic::Status::internal("message not published")
            })
    }
}
