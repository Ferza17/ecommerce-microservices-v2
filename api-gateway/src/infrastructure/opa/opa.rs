use crate::config::config::AppConfig;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug)]
pub struct OPA {
    path: String,
    client: reqwest::Client,
}

#[derive(Debug, Serialize)]
pub struct OpaInput {
    pub method: String,
    pub path: String,
    pub user_id: String,
    pub user_role: String,
}

impl OPA {
    pub fn new(app_config: AppConfig) -> Self {
        Self {
            path: app_config.opa_path,
            client: reqwest::Client::new(),
        }
    }

    #[tracing::instrument]
    pub async fn validate_http_access(&self, request: OpaInput) -> Result<bool, anyhow::Error> {
        #[derive(Debug, Serialize)]
        struct OpaRequest {
            input: OpaInput,
        }

        #[derive(Debug, Deserialize)]
        struct OpaResponse {
            result: bool,
        }

        let response = self
            .client
            .post(&self.path)
            .json(&OpaRequest { input: request })
            .send()
            .await?;
        let opa_response: OpaResponse = response.json().await?;
        Ok(opa_response.result)
    }
}
