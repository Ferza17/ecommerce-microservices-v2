use crate::config::config::AppConfig;
use schema_registry_converter::async_impl::schema_registry::post_schema;
use schema_registry_converter::schema_registry_common::{
    RegisteredSchema, SchemaType, SubjectNameStrategy, SuppliedSchema,
};

pub const SHIPPING_PROVIDERS_SCHEMA: &str = r#"
    {
      "$schema": "http://json-schema.org/draft-07/schema#",
      "title": "ShippingProvider",
      "type": "object",
      "properties": {
        "id": { "type": "string" },
        "name": { "type": "string" },
        "created_at": { "type": "string", "format": "date-time" },
        "updated_at": { "type": "string", "format": "date-time" }
      },
      "required": ["id", "name", "created_at", "updated_at"]
    }
    "#;

pub async fn publish_shipping_providers_schema(
    app_config: AppConfig,
) -> Result<RegisteredSchema, anyhow::Error> {
    let schema = SuppliedSchema {
        name: Option::from(String::from(
            app_config
                .message_broker_kafka_topic_sink_shipping
                .pg_shippings_shipping_providers
                .clone(),
        )),
        schema_type: SchemaType::Json,
        schema: SHIPPING_PROVIDERS_SCHEMA.to_string(),
        references: vec![],
    };

    let sr_settings = schema_registry_converter::async_impl::schema_registry::SrSettings::new(
        app_config.message_broker_kafka.schema_registry_url.clone(),
    );

    let subject_name_strategy = match SubjectNameStrategy::TopicNameStrategy(
        app_config
            .message_broker_kafka_topic_sink_shipping
            .pg_shippings_shipping_providers,
        false,
    )
    .get_subject()
    {
        Ok(v) => v,
        Err(e) => Err(anyhow::Error::msg(e.error))?,
    };

    match post_schema(&sr_settings, subject_name_strategy, schema).await {
        Ok(d) => Ok(d),
        Err(e) => Err(anyhow::Error::msg(e.error))?,
    }
}
