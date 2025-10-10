use crate::config::config::AppConfig;
use mongodb::Database;

pub async fn get_connection(config: &AppConfig) -> Database {
    let url = format!(
        "mongodb://{}:{}@{}:{}",
        config.database_mongodb.username,
        config.database_mongodb.password,
        config.database_mongodb.host,
        config.database_mongodb.port,
    );

    let client = mongodb::Client::with_uri_str(&url).await.unwrap();
    return client.database(&config.database_mongodb.database);
}
