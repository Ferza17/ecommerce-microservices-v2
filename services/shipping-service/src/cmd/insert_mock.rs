use crate::config::config::AppConfig;
use crate::model::schema_registry::shipping_providers::publish_shipping_providers_schema;
use anyhow::Error;
use clap::Args;
use consulrs::client::{ConsulClient, ConsulClientSettingsBuilder};
use schema_registry_converter::schema_registry_common::SubjectNameStrategy;
use std::fs::File;
use std::io::BufReader;
use std::time::Duration;

#[derive(Args, Debug)]
pub struct InsertMockArgs {
    #[arg(short, long, help = "run direction: 'local' or 'production'")]
    pub direction: String,
}

pub async fn handle_insert_mock_command(args: InsertMockArgs) {
    let mut cfg = AppConfig::new(&*args.direction)
        .await
        .map_err(|e| {
            eprintln!("Failed to load configuration: {}", e);
            std::process::exit(1);
        })
        .unwrap();

    let client = ConsulClient::new(
        ConsulClientSettingsBuilder::default()
            .address(format!(
                "http://{}:{}",
                cfg.config_env.consul_host, cfg.config_env.consul_port
            ))
            .build()
            .map_err(|e| eprintln!(" Error Consul :  {:?}", e))
            .unwrap(),
    )
    .unwrap();

    cfg = cfg
        .with_message_broker_kafka_topic_sink_shipping_from_consul(&client)
        .with_message_broker_kafka_from_consul(&client);

    let topic_name = cfg
        .message_broker_kafka_topic_sink_shipping
        .pg_shippings_shipping_providers
        .clone();
    let kafka = crate::infrastructure::message_broker::kafka::KafkaInfrastructure::new(cfg.clone());

    let file = File::open("mock_shipping_providers.json").unwrap();
    let reader = BufReader::new(file);

    // Register Schema
    match publish_shipping_providers_schema(cfg.clone()).await {
        Ok(_) => {}
        Err(_) => {
            eprintln!("Schema already registered");
            return;
        }
    }

    let sr_settings = schema_registry_converter::async_impl::schema_registry::SrSettings::new(
        cfg.message_broker_kafka.schema_registry_url.clone(),
    );
    let mut encoder = schema_registry_converter::async_impl::json::JsonEncoder::new(sr_settings);

    let shipping_providers: Vec<crate::model::diesel::shipping_providers::ShippingProviders> =
        serde_json::from_reader(reader).unwrap();

    for provider in shipping_providers {
        let subject_name_strategy =
            SubjectNameStrategy::TopicNameStrategy(topic_name.clone(), false);

        let payload = match encoder
            .encode(
                &serde_json::to_value(&provider).unwrap(),
                subject_name_strategy,
            )
            .await
        {
            Ok(v) => v,
            Err(e) => {
                eprintln!("Error: {}", e);
                continue;
            }
        };

        match kafka
            .publish(
                rdkafka::producer::FutureRecord::to(topic_name.as_str())
                    .key(provider.id.to_string().as_bytes())
                    .payload(&payload),
            )
            .await
        {
            Ok(_) => eprintln!("sent : {:?}", provider),
            Err(e) => eprintln!("Error: {}", e),
        }

        tokio::time::sleep(Duration::from_millis(500)).await;
    }
}
