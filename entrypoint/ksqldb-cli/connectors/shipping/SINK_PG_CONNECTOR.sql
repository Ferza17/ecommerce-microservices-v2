CREATE
SINK CONNECTOR SINK_CONNECTOR_PG_SHIPPINGS_SHIPPING_PROVIDERS
WITH (
    'connector.class'= 'io.confluent.connect.jdbc.JdbcSinkConnector',
    'tasks.max'= '1',
    'topics'= 'sink.pg.shippings.shipping_providers',
    'connection.url'= 'jdbc=postgresql=//postgres-local=5432/shippings',
    'connection.user'= 'postgres',
    'connection.password'= '1234',
    'insert.mode'= 'upsert',
    'pk.mode'= 'record_key',
    'pk.fields'= 'id',
    'auto.create'= 'true',
    'auto.evolve'= 'true',
    'table.name.format'= 'shipping_providers',
    'transforms'= 'flatten',
    'transforms.flatten.type'= 'org.apache.kafka.connect.transforms.Flatten$Value',
    'transforms.flatten.delimiter'= '_'
);


CREATE
SINK CONNECTOR SINK_CONNECTOR_PG_SHIPPINGS_SHIPPINGS
WITH (
    'connector.class'= 'io.confluent.connect.jdbc.JdbcSinkConnector',
    'tasks.max'= '1',
    'topics'= 'sink.pg.shippings.shippings',
    'connection.url'= 'jdbc=postgresql=//postgres-local=5432/shippings',
    'connection.user'= 'postgres',
    'connection.password'= '1234',
    'insert.mode'= 'upsert',
    'pk.mode'= 'record_key',
    'pk.fields'= 'id',
    'auto.create'= 'true',
    'auto.evolve'= 'true',
    'table.name.format'= 'shippings',
    'transforms'= 'flatten',
    'transforms.flatten.type'= 'org.apache.kafka.connect.transforms.Flatten$Value',
    'transforms.flatten.delimiter'= '_'
);


SHOW CONNECTORS;
