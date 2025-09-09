use crate::infrastructure::database::async_postgres::AsyncPgDeadPool;
use crate::model::diesel::schema::shipping_providers::dsl::{
    discarded_at, id, shipping_providers as ShippingProviderSchema,
};
use crate::model::diesel::shipping_providers::ShippingProviders as ShippingProviderModel;
use anyhow::{Context, Result};
use deadpool::managed::Object;
use diesel::QueryDsl;
use diesel::prelude::*;
use diesel_async::pooled_connection::AsyncDieselConnectionManager;
use diesel_async::{AsyncPgConnection, RunQueryDsl};
use std::fmt;
use tracing::{Level, event, instrument};

pub trait ShippingProviderPostgresRepository {
    async fn get_shipping_provider_by_id(
        &self,
        request_id: &str,
        provider_id: &str,
    ) -> Result<ShippingProviderModel>;

    async fn list_shipping_providers(
        &self,
        request_id: &str,
        page: &u32,
        limit: &u32,
    ) -> Result<Vec<ShippingProviderModel>>;
}

#[derive(Clone)]
pub struct ShippingProviderPostgresRepositoryImpl {
    pg: AsyncPgDeadPool,
}

impl fmt::Debug for ShippingProviderPostgresRepositoryImpl {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        f.debug_struct("ShippingProviderPostgresRepository")
            .field("pg", &"AsyncPgDeadPool")
            .finish()
    }
}

impl ShippingProviderPostgresRepositoryImpl {
    pub fn new(
        pg: deadpool::managed::Pool<
            AsyncDieselConnectionManager<AsyncPgConnection>,
            Object<AsyncDieselConnectionManager<AsyncPgConnection>>,
        >,
    ) -> Self {
        Self { pg }
    }
}

impl ShippingProviderPostgresRepository for ShippingProviderPostgresRepositoryImpl {
    #[instrument("ShippingProviderPostgresRepository.get_shipping_provider_by_id")]
    async fn get_shipping_provider_by_id(
        &self,
        request_id: &str,
        provider_id: &str,
    ) -> Result<ShippingProviderModel> {
        event!(name: "ShippingProviderPostgresRepository.get_shipping_provider_by_id", Level::INFO, request_id = request_id, provider_id = provider_id);
        match ShippingProviderSchema
            .select(ShippingProviderModel::as_select())
            .filter(id.eq(provider_id))
            .first(&mut self.pg.get().await?)
            .await
            .with_context(|| format!("Failed to find shipping provider with id: {}", provider_id))
        {
            Ok(shipping_provider) => Ok(shipping_provider),
            Err(e) => Err(e),
        }
    }

    #[instrument(name = "ShippingProviderPostgresRepository.list_shipping_providers")]
    async fn list_shipping_providers(
        &self,
        request_id: &str,
        page: &u32,
        limit: &u32,
    ) -> Result<Vec<crate::model::diesel::shipping_providers::ShippingProviders>> {
        event!(
            Level::INFO,
            name = "ShippingProviderPostgresRepository.list_shipping_providers",
            request_id = request_id
        );
        match ShippingProviderSchema
            .select(ShippingProviderModel::as_select())
            .filter(discarded_at.is_null())
            .limit(*limit as i64)
            .offset((*page - 1) as i64 * *limit as i64)
            .load::<ShippingProviderModel>(&mut self.pg.get().await?)
            .await
            .with_context(|| "Failed to find list_shipping_providers")
        {
            Ok(result) => Ok(result),
            Err(e) => Err(e.into()),
        }
    }
}
