use crate::config::config::AppConfig;
use deadpool::managed::Object;
use diesel_async::AsyncPgConnection;
use diesel_async::pooled_connection::AsyncDieselConnectionManager;
use diesel_async::pooled_connection::deadpool::Pool;

pub type AsyncPgDeadPool = deadpool::managed::Pool<
    AsyncDieselConnectionManager<AsyncPgConnection>,
    Object<AsyncDieselConnectionManager<AsyncPgConnection>>,
>;

pub async fn get_connection(config: &AppConfig) -> AsyncPgDeadPool {
    // create a new connection pool with the default config
    let config = AsyncDieselConnectionManager::<AsyncPgConnection>::new(
        format!(
            "postgres://{}:{}@{}:{}/{}",
            config.database_postgres_username,
            config.database_postgres_password,
            config.database_postgres_host,
            config.database_postgres_port,
            config.database_postgres_database,
        )
        .as_str(),
    );

    Pool::builder(config)
        .build()
        .unwrap_or_else(|_| panic!("Error creating pool"))
}
