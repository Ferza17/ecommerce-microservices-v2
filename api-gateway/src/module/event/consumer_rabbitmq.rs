use crate::config::config::AppConfig;
use crate::infrastructure::message_broker::rabbitmq::RabbitMQInfrastructure;
use crate::package::context::auth::AUTHORIZATION_HEADER;
use crate::package::context::request_id::X_REQUEST_ID_HEADER;
use futures::StreamExt;
use lapin::types::AMQPValue;
use lapin::{Error, ExchangeKind};
use prost::{DecodeError, Message};
use tokio::sync::mpsc::UnboundedReceiver;

#[derive(Debug, Clone)]
pub struct Consumer {
    app_config: AppConfig,
    infra_message_broker_rabbitmq: RabbitMQInfrastructure,
}

impl Consumer {
    pub fn new(
        app_config: AppConfig,
        infra_message_broker_rabbitmq: RabbitMQInfrastructure,
    ) -> Self {
        Self {
            app_config,
            infra_message_broker_rabbitmq,
        }
    }

    pub async fn consume_event_created(
        &self,
    ) -> UnboundedReceiver<Result<crate::model::rpc::event::AppendRequest, Error>> {
        let (tx, rx) = tokio::sync::mpsc::unbounded_channel();

        let mut messages = self
            .infra_message_broker_rabbitmq
            .binding(
                self.app_config
                    .service_event_rabbitmq
                    .queue_event_api_gateway_event_created
                    .as_str(),
                self.app_config
                    .service_event_rabbitmq
                    .exchange_fanout_event
                    .as_str(),
                ExchangeKind::Fanout,
            )
            .await
            .setup_consumer(
                self.app_config
                    .service_event_rabbitmq
                    .queue_event_api_gateway_event_created
                    .as_str(),
            )
            .await;

        tokio::spawn(async move {
            while let Some(delivery) = messages.next().await {
                match delivery {
                    Ok(delivery) => {
                        let mut token = String::new();
                        let mut request_id = String::new();
                        if let Some(headers) = delivery.properties.headers() {
                            for (key, value) in headers {
                                if key.to_string().to_lowercase().as_str() == AUTHORIZATION_HEADER {
                                    token = match value {
                                        AMQPValue::ShortString(str) => str.to_string(),
                                        AMQPValue::LongString(str) => str.to_string(),
                                        _ => "".to_string(),
                                    }
                                }

                                if key.to_string().to_lowercase().as_str() == X_REQUEST_ID_HEADER {
                                    request_id = match value {
                                        AMQPValue::ShortString(str) => str.to_string(),
                                        AMQPValue::LongString(str) => str.to_string(),
                                        _ => "".to_string(),
                                    }
                                }
                            }
                        }
                        

                        let data = match crate::model::rpc::event::AppendRequest::decode(
                            &*delivery.data,
                        ) {
                            Ok(data) => data,
                            Err(err) => {
                                eprintln!("{:?}", err);
                                return;
                            }
                        };

                        tx.send(Ok(data)).unwrap();
                    }
                    Err(err) => {
                        let _ = tx.send(Err(err));
                    }
                }
            }
        });

        rx
    }
}
