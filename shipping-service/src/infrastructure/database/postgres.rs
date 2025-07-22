use crate::config::config::AppConfig;
use std::sync::Arc;
use tokio_postgres::{Client, Error, NoTls};

#[derive(Clone)]
pub struct PostgresInfrastructure {
    pub client: Arc<Client>,
}

impl PostgresInfrastructure {
    pub async fn new(config: AppConfig) -> Result<Self, Error> {
        // Create connection string
        let conn_str = format!(
            "host={} user={} password={} port={} dbname={}",
            config.database_postgres_host,
            config.database_postgres_username,
            config.database_postgres_password,
            config.database_postgres_port,
            config.database_postgres_database,
        );

        // Connect to PostgreSQL (async)
        let (client, connection) = tokio_postgres::connect(&conn_str, NoTls).await?;

        // Spawn the connection task (it drives the connection)
        tokio::spawn(async move {
            if let Err(e) = connection.await {
                eprintln!("PostgresSQL connection error: {}", e);
            }
        });

        println!("Successfully connected to PostgresSQL (async).");

        Ok(Self {
            client: Arc::new(client),
        })
    }
}
