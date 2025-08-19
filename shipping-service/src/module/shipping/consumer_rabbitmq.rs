use crate::model::rpc::shipping::{CreateShippingRequest, UpdateShippingRequest};
use crate::module::shipping::usecase::{ShippingUseCase, ShippingUseCaseImpl};
use crate::package::context::auth::AUTHORIZATION_HEADER;
use crate::package::context::request_id::X_REQUEST_ID_HEADER;
use anyhow::{Error, Result};
use lapin::message::Delivery;
use lapin::types::AMQPValue;
use prost::Message;
use tracing::{error};
use uuid::Uuid;

pub struct ShippingRabbitMQConsumer {
    shipping_use_case: ShippingUseCaseImpl,
}

impl ShippingRabbitMQConsumer {
    pub fn new(shipping_use_case: ShippingUseCaseImpl) -> Self {
        Self { shipping_use_case }
    }

    pub async fn consumer_shipping_created(&self, delivery: Delivery) -> Result<(), Error> {
        let mut token = String::new();
        let mut request_id = String::new();
        if let Some(headers) = delivery.properties.headers() {
            for (key, value) in headers {
                if key.to_string().to_lowercase().as_str() == AUTHORIZATION_HEADER {
                    token = match value {
                        AMQPValue::ShortString(str) => str.to_string(),
                        AMQPValue::LongString(str) => str.to_string(),
                        _ => {
                            return Err::<(), Error>(Error::msg(format!(
                                "missing {} header",
                                AUTHORIZATION_HEADER.to_lowercase()
                            )));
                        }
                    }
                }

                if key.to_string().to_lowercase().as_str() == X_REQUEST_ID_HEADER {
                    request_id = match value {
                        AMQPValue::ShortString(str) => str.to_string(),
                        AMQPValue::LongString(str) => str.to_string(),
                        _ => Uuid::new_v4().to_string(),
                    }
                }
            }
        } else {
            error!("No headers found");
        }

        let data = match CreateShippingRequest::decode(&*delivery.data) {
            Ok(data) => data,
            Err(err) => {
                error!("[shipping] consumer_shipping_created : {}", err);
                return Err(Error::msg(format!(
                    "Failed to decode shipping request: {}",
                    err
                )));
            }
        };

        match self
            .shipping_use_case
            .create_shipping(request_id, token, tonic::Request::new(data))
            .await
        {
            Ok(_) => Ok(delivery
                .acker
                .ack(lapin::options::BasicAckOptions::default())
                .await?),
            Err(err) => {
                error!("[shipping] consumer_shipping_created : {}", err);
                Err(Error::msg(format!(
                    "[shipping] consumer_shipping_created : {}",
                    err
                )))
            }
        }
    }

    pub async fn consumer_shipping_updated(&self, delivery: Delivery) -> Result<(), Error> {
        let mut request_id = String::new();
        let mut token = String::new();
        if let Some(headers) = delivery.properties.headers() {
            for (key, value) in headers {
                if key.to_string().to_lowercase().as_str() == AUTHORIZATION_HEADER {
                    token = match value {
                        AMQPValue::ShortString(str) => str.to_string(),
                        AMQPValue::LongString(str) => str.to_string(),
                        _ => {
                            return Err::<(), Error>(Error::msg(format!(
                                "missing {} header",
                                AUTHORIZATION_HEADER.to_lowercase()
                            )));
                        }
                    }
                }

                if key.to_string().to_lowercase().as_str() == X_REQUEST_ID_HEADER {
                    request_id = match value {
                        AMQPValue::ShortString(str) => str.to_string(),
                        AMQPValue::LongString(str) => str.to_string(),
                        _ => Uuid::new_v4().to_string(),
                    }
                }
            }
        } else {
            error!("No headers found");
        }

        let data = match UpdateShippingRequest::decode(&*delivery.data) {
            Ok(data) => data,
            Err(err) => {
                error!("[shipping] consumer_shipping_updated : {}", err);
                return Err(Error::msg(format!(
                    "Failed to decode shipping update: {}",
                    err
                )));
            }
        };

        match self
            .shipping_use_case
            .update_shipping(request_id, token,tonic::Request::new(data))
            .await
        {
            Ok(_) => Ok(delivery
                .acker
                .ack(lapin::options::BasicAckOptions::default())
                .await?),
            Err(err) => {
                error!("[shipping] consumer_shipping_updated : {}", err);
                Err(Error::msg(format!(
                    "[shipping] consumer_shipping_updated : {}",
                    err
                )))
            }
        }
    }
}
