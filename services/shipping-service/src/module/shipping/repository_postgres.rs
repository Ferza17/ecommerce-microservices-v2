use crate::infrastructure::database::async_postgres::AsyncPgDeadPool;
use crate::model::diesel::schema::shippings::dsl::{discarded_at, id, shippings as shippingSchema};
use crate::model::diesel::shippings::{
    CreateShippings, Shippings as shippingModel, Shippings, UpdateShippings,
};
use anyhow::Context;
use anyhow::{Error, Result};
use diesel::QueryDsl;
use diesel::SelectableHelper;
use diesel::{ExpressionMethods, QueryResult};
use diesel_async::RunQueryDsl;
use std::fmt;
use tracing::{Level, event, instrument};

pub trait ShippingPostgresRepository {
    async fn create_shipping(
        &self,
        request_id: &str,
        shipping: &CreateShippings,
    ) -> Result<(), Error>;
    async fn get_shipping_by_id(
        &self,
        request_id: &str,
        shipping_id: &str,
    ) -> Result<shippingModel, Error>;
    async fn list_shipping(
        &self,
        request_id: &str,
        page: &u32,
        limit: &u32,
    ) -> Result<Vec<shippingModel>, Error>;
    async fn update_shipping(
        &self,
        request_id: &str,
        shipping_id: &str,
        shipping: &UpdateShippings,
    ) -> Result<(), Error>;

    async fn delete_shipping(&self, request_id: &str, shipping_id: &str) -> Result<(), Error>;
}

#[derive(Clone)]
pub struct ShippingPostgresRepositoryImpl {
    pg: AsyncPgDeadPool,
}

impl ShippingPostgresRepositoryImpl {
    pub fn new(pg: AsyncPgDeadPool) -> Self {
        Self { pg }
    }
}

impl fmt::Debug for ShippingPostgresRepositoryImpl {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        f.debug_struct("ShippingPostgresRepository")
            .field("pg", &"AsyncPgDeadPool")
            .finish()
    }
}

impl ShippingPostgresRepository for ShippingPostgresRepositoryImpl {
    async fn create_shipping(
        &self,
        request_id: &str,
        shipping: &CreateShippings,
    ) -> Result<(), Error> {
        event!(
            Level::INFO,
            name = "ShippingPostgresRepository.create_shipping",
            request_id = request_id
        );

        diesel::insert_into(shippingSchema)
            .values(shipping)
            .execute(&mut self.pg.get().await?)
            .await
            .map(|_| ())
            .map_err(Error::from)
    }

    async fn get_shipping_by_id(
        &self,
        request_id: &str,
        shipping_id: &str,
    ) -> Result<Shippings, Error> {
        event!(name: "ShippingPostgresRepository.get_shipping_by_id", Level::INFO, request_id = request_id, shipping_id = shipping_id);
        match shippingSchema
            .select(shippingModel::as_select())
            .filter(id.eq(shipping_id))
            .first(&mut self.pg.get().await?)
            .await
            .with_context(|| format!("Failed to find shipping with id: {}", shipping_id))
        {
            Ok(shipping) => Ok(shipping),
            Err(e) => Err(e),
        }
    }

    #[instrument("ShippingPostgresRepository.list_shipping")]
    async fn list_shipping(
        &self,
        request_id: &str,
        page: &u32,
        limit: &u32,
    ) -> Result<Vec<shippingModel>, Error> {
        event!(
            Level::INFO,
            name = "ShippingPostgresRepository.list_shipping",
            request_id = request_id
        );
        match shippingSchema
            .select(shippingModel::as_select())
            .filter(discarded_at.is_null())
            .limit(*limit as i64)
            .offset((*page - 1) as i64 * *limit as i64)
            .load::<shippingModel>(&mut self.pg.get().await?)
            .await
            .with_context(|| "Failed to list shipping")
        {
            Ok(result) => Ok(result),
            Err(e) => Err(e.into()),
        }
    }

    async fn update_shipping(
        &self,
        request_id: &str,
        shipping_id: &str,
        shipping: &UpdateShippings,
    ) -> Result<(), Error> {
        event!(
            Level::INFO,
            name = "ShippingPostgresRepository.update_shipping",
            request_id = request_id
        );

        match diesel::update(shippingSchema.filter(id.eq(shipping_id)))
            .set(shipping)
            .execute(&mut self.pg.get().await?)
            .await
        {
            Ok(_) => Ok(()),
            Err(err) => {
                return Err(err.into());
            }
        }
    }

    async fn delete_shipping(&self, request_id: &str, shipping_id: &str) -> Result<(), Error> {
        event!(
            Level::INFO,
            name = "ShippingPostgresRepository.delete_shipping",
            request_id = request_id
        );
        match diesel::delete(shippingSchema.filter(id.eq(shipping_id)))
            .execute(&mut self.pg.get().await?)
            .await
        {
            Ok(_) => Ok(()),
            Err(_) => {
                return Err(Error::msg("Failed to delete shipping"));
            }
        }
    }
}
