use consulrs::client::ConsulClient;

#[derive(Clone, Debug)]
pub struct DatabaseMongoDB {
    pub host: String,
    pub port: String,
    pub username: String,
    pub password: String,
    pub database: String,
}

impl Default for DatabaseMongoDB {
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

impl DatabaseMongoDB {
    pub async fn with_consul_client(
        &mut self,
        env: String,
        client: &ConsulClient,
    ) -> Result<Self, Box<dyn std::error::Error>> {
        self.host = crate::config::config::get_kv(
            client,
            format!("{}/database/mongodb/MONGO_HOST", env),
        )
            .await
            .parse()
            .unwrap_or_else(|_| "".to_string());

        self.port = crate::config::config::get_kv(
            client,
            format!("{}/database/mongodb/MONGO_PORT", env),
        )
            .await
            .parse()
            .unwrap_or_else(|_| "".to_string());

        self.username = crate::config::config::get_kv(
            client,
            format!("{}/database/mongodb/MONGO_USERNAME", env),
        )
            .await
            .parse()
            .unwrap_or_else(|_| "".to_string());

        self.password = crate::config::config::get_kv(
            client,
            format!("{}/database/mongodb/MONGO_PASSWORD", env),
        )
            .await
            .parse()
            .unwrap_or_else(|_| "".to_string());

        self.database = crate::config::config::get_kv(
            client,
            format!("{}/database/mongodb/MONGO_DATABASE_NAME/EVENT_STORE", env),
        )
            .await
            .parse()
            .unwrap_or_else(|_| "".to_string());

        Ok((self.clone()))
    }
}
