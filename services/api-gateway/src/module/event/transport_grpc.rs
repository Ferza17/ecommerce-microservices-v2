#[derive(Debug, Clone)]
pub struct Transport {
    event_store_client:
        crate::model::rpc::event::event_store_client::EventStoreClient<tonic::transport::Channel>,
}

impl Transport {
    pub async fn new(config: crate::config::config::AppConfig) -> Result<Self, anyhow::Error> {
        let channel = tonic::transport::Channel::from_shared(
            format!(
                "http://{}:{}",
                config.service_event.rpc_host, config.service_event.rpc_port
            )
            .to_string(),
        )
        .expect("Failed to connect to event store service")
        .connect()
        .await
        .map_err(|e| panic!("event store service not connected : {}", e))
        .unwrap();
        Ok(Self {
            event_store_client: crate::model::rpc::event::event_store_client::EventStoreClient::new(
                channel,
            ),
        })
    }

    pub async fn append(
        &mut self,
        request_id: String,
        mut request: tonic::Request<crate::model::rpc::event::AppendRequest>,
    ) -> Result<crate::model::rpc::event::AppendResponse, tonic::Status> {
        request.metadata_mut().insert(
            crate::package::context::request_id::X_REQUEST_ID_HEADER,
            request_id.parse().unwrap(),
        );
        match self.event_store_client.append(request).await {
            Ok(response) => {
                tracing::event!(
                    tracing::Level::INFO,
                    request_id = request_id,
                    data=?response
                );
                Ok(response.into_inner())
            }
            Err(err) => {
                tracing::event!(
                    tracing::Level::ERROR,
                    request_id = request_id,
                    error = %err,
                    "Failed to append event"
                );
                Err(err.into())
            }
        }
    }
}
