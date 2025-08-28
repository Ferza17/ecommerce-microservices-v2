use crate::model::rpc::response::ResponseCommand;
use crate::model::rpc::response::response_command::ResponseCommandData;
use futures::{SinkExt, StreamExt};

#[derive(Debug, Clone)]
pub struct Presenter {
    notification_use_case: crate::module::notification::usecase::UseCase,
}

pub const ROUTE_PREFIX: &str = "/api/v1/notification";
pub const TAG: &str = "Notification";

impl Presenter {
    pub fn new(notification_use_case: crate::module::notification::usecase::UseCase) -> Self {
        Self {
            notification_use_case,
        }
    }
    pub fn router(&self) -> axum::Router {
        axum::Router::new()
            .route(
                "/{request_id}",
                axum::routing::get(get_notification_with_request_id),
            )
            .with_state(self.clone())
    }
}

// FOR TESTING
#[derive(Debug, Clone, serde::Serialize, serde::Deserialize)]
pub struct ClientMessage {
    pub message_type: String,
    pub correlation_id: Option<String>,
    pub payload: serde_json::Value,
}

#[tracing::instrument("notification.http_presenter.get_notification_with_request_id")]
pub async fn get_notification_with_request_id(
    ws: axum::extract::ws::WebSocketUpgrade,
    axum::extract::Path(request_id): axum::extract::Path<String>,
    axum::extract::State(state): axum::extract::State<Presenter>,
) -> impl axum::response::IntoResponse {
    async fn send_notification_to_client(
        sender: &mut futures::stream::SplitSink<
            axum::extract::ws::WebSocket,
            axum::extract::ws::Message,
        >,
        req: ResponseCommand,
    ) {
        if sender
            .send(axum::extract::ws::Message::Text(
                axum::extract::ws::Utf8Bytes::from(serde_json::to_string(&req).unwrap()),
            ))
            .await
            .is_err()
        {
            return;
        }
    }

    ws.on_upgrade(|socket| async move {
        let (mut sender, mut receiver) = socket.split();
        // Send a welcome message
        send_notification_to_client(
            &mut sender,
            ResponseCommand {
                status: "success".to_string(),
                message: "waiting response to be ready".to_string(),
                data: None,
            },
        )
        .await;

        // Handle incoming messages from client
        tokio::spawn(async move {
            while let Some(msg) = receiver.next().await {
                match msg {
                    Ok(axum::extract::ws::Message::Text(text)) => {
                        // TODO: Get from redis or rabbitmq
                        // Require match request_id
                        // send to client

                        // Handle client messages (e.g., subscription changes, heartbeat)
                        if let Ok(client_msg) = serde_json::from_str::<ClientMessage>(&text) {
                            match client_msg.message_type.as_str() {
                                "heartbeat" => {
                                    send_notification_to_client(
                                        &mut sender,
                                        ResponseCommand {
                                            status: "pending".to_string(),
                                            message: "waiting response to be ready".to_string(),
                                            data: Option::from(ResponseCommandData {
                                                request_id: request_id.clone(),
                                                websocket_notification_url: "".to_string(),
                                            }),
                                        },
                                    )
                                    .await;
                                }
                                _ => {
                                    println!("Received client message: {:?}", client_msg);
                                }
                            }
                        }
                    }
                    Ok(_) => {}
                    Err(_) => {}
                }
            }
        });
    })
}
