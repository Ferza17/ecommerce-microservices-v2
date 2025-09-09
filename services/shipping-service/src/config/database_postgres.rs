use consulrs::client::ConsulClient;

#[derive(Clone, Debug)]
pub struct DatabasePostgres {
    pub host: String,
    pub port: String,
    pub username: String,
    pub password: String,
    pub database: String,
}

impl Default for DatabasePostgres {
    fn default() -> Self {
        Self {
            host: "".to_string(),
            port: "".to_string(),
            username: "".to_string(),
            password: "".to_string(),
            database: "".to_string(),
        }
    }
}

impl DatabasePostgres {
    pub async fn with_consul_client(
        &mut self,
        env: String,
        client: &ConsulClient,
    ) -> Result<Self, Box<dyn std::error::Error>> {
        self.host = crate::config::config::get_kv(
            client,
            format!("{}/database/postgres/POSTGRES_HOST", env),
        )
        .await
        .parse()
        .unwrap_or_else(|_| "".to_string());

        self.port = crate::config::config::get_kv(
            client,
            format!("{}/database/postgres/POSTGRES_PORT", env),
        )
        .await
        .parse()
        .unwrap_or_else(|_| "".to_string());

        self.username = crate::config::config::get_kv(
            client,
            format!("{}/database/postgres/POSTGRES_USERNAME", env),
        )
        .await
        .parse()
        .unwrap_or_else(|_| "".to_string());

        self.password = crate::config::config::get_kv(
            client,
            format!("{}/database/postgres/POSTGRES_PASSWORD", env),
        )
        .await
        .parse()
        .unwrap_or_else(|_| "".to_string());

        self.database = crate::config::config::get_kv(
            client,
            format!("{}/database/postgres/POSTGRES_DATABASE_NAME/SHIPPINGS", env),
        )
        .await
        .parse()
        .unwrap_or_else(|_| "".to_string());

        Ok((self.clone()))
    }
}
