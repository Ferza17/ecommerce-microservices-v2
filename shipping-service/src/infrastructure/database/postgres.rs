use crate::config::config::AppConfig;
use anyhow::Result;
use diesel::prelude::*;
use diesel::r2d2::{ConnectionManager, Pool as R2D2Pool};
use diesel_migrations::{EmbeddedMigrations, MigrationHarness, embed_migrations};

pub type PostgresPool = R2D2Pool<ConnectionManager<PgConnection>>;

pub async fn create_postgres_pool(config: &AppConfig) -> Result<PostgresPool> {
    let manager = ConnectionManager::<PgConnection>::new(
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
    let pool = R2D2Pool::builder()
        .max_size(5)
        .connection_timeout(std::time::Duration::from_secs(5))
        .test_on_check_out(true)
        .build(manager)?;

    Ok(pool)
}

pub fn run_migrations(connection: &mut PgConnection) -> Result<()> {
    const MIGRATIONS: EmbeddedMigrations = embed_migrations!();
    connection
        .run_pending_migrations(MIGRATIONS)
        .expect("Error running migrations");
    Ok(())
}
