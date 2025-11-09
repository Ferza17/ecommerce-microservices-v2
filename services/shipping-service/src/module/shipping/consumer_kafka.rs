use crate::model::rpc::shipping::CreateShippingRequest;
use crate::module::shipping::usecase::{ShippingUseCase, ShippingUseCaseImpl};
use crate::package::context::auth::AUTHORIZATION_HEADER;
use crate::package::context::request_id::X_REQUEST_ID_HEADER;
use crate::package::context::traceparent::TRACEPARENT_HEADER;
use prost::Message as prostMessage;
use rdkafka::message::{Headers, Message as rdKafkaMessage, OwnedMessage};
use tracing::{error, info};

pub struct Consumer {
    shipping_use_case: ShippingUseCaseImpl,
}

impl Consumer {
    pub fn new(shipping_use_case: ShippingUseCaseImpl) -> Self {
        Self { shipping_use_case }
    }

    pub async fn consume_snapshot_shippings_shipping_created(
        &self,
        message: OwnedMessage,
    ) -> Result<(), anyhow::Error> {
        let mut request_id = String::new();
        let mut token = String::new();
        match message.headers() {
            None => {
                info!("no headers found")
            }
            Some(h) => {
                for header in h.iter() {
                    if header.key == X_REQUEST_ID_HEADER {
                        request_id = header.key.to_string();
                        continue;
                    }

                    if header.key == AUTHORIZATION_HEADER {
                        token = header.key.to_string();
                        continue;
                    }

                    //TODO: Traceparent Header
                    if header.key == TRACEPARENT_HEADER {}
                }
            }
        }

        let mut request = CreateShippingRequest::default();
        match message.payload() {
            None => {
                info!("no payload found");
                Err(anyhow::Error::msg("no payload found"))?
            }
            Some(p) => {
                request = match CreateShippingRequest::decode(&*p) {
                    Ok(v) => v,
                    Err(err) => {
                        error!(
                            "[consume_snapshot_shippings_shipping_created] consume_snapshot_shippings_shipping_created : {}",
                            err
                        );
                        return Err(anyhow::Error::msg(err.to_string()));
                    }
                }
            }
        }

        match self
            .shipping_use_case
            .create_shipping(request_id, token, tonic::Request::new(request))
            .await
        {
            Ok(response) => {
                info!(
                    "[consume_snapshot_shippings_shipping_created] response: {:?}",
                    response
                );
            }
            Err(err) => {
                error!(
                    "[consume_snapshot_shippings_shipping_created] error: {:?}",
                    err
                );
                return Err(anyhow::Error::msg(err.to_string()));
            }
        }

        Ok(())
    }
}
