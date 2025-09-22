use crate::config::config::AppConfig;
use clap::Args;
use consulrs::client::{ConsulClient, ConsulClientSettingsBuilder};
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

    let kafka = crate::infrastructure::message_broker::kafka::KafkaInfrastructure::new(cfg.clone());

    let file = File::open("mock_shipping_providers.json").unwrap();
    let reader = BufReader::new(file);

    let shipping_providers: Vec<crate::model::diesel::shipping_providers::ShippingProviders> =
        serde_json::from_reader(reader).unwrap();

    for provider in shipping_providers {
        match kafka
            .publish_with_json_schema(
                cfg.message_broker_kafka_topic_sink_shipping
                    .pg_shippings_shipping_providers
                    .clone(),
                crate::model::schema_registry::registry::Registry::ShippingProvider,
                serde_json::to_value(&provider.clone()).unwrap(),
                provider.clone().id,
                None,
            )
            .await
        {
            Ok(_) => eprintln!("sent : {:?}", provider),
            Err(e) => eprintln!("Error: {}", e),
        }
        tokio::time::sleep(Duration::from_millis(500)).await;
    }
}
