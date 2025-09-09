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
    let config = AsyncDieselConnectionManager::<AsyncPgConnection>::new(
        format!(
            "postgres://{}:{}@{}:{}/{}",
            config.database_postgres.username,
            config.database_postgres.password,
            config.database_postgres.host,
            config.database_postgres.port,
            config.database_postgres.database,
        )
        .as_str(),
    );

    Pool::builder(config)
        .build()
        .unwrap_or_else(|_| panic!("Error creating pool"))
}
