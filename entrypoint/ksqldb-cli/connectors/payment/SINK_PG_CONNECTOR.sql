CREATE
SINK CONNECTOR SINK_CONNECTOR_PG_PAYMENTS_PAYMENT_PROVIDERS
WITH (
    'connector.class'= 'io.confluent.connect.jdbc.JdbcSinkConnector',
    'tasks.max'= '1',
    'topics'= 'sink.pg.payments.payment_providers',
    'connection.url'= 'jdbc=postgresql=//postgres-local=5432/payments',
    'connection.user'= 'postgres',
    'connection.password'= '1234',
    'insert.mode'= 'upsert',
    'pk.mode'= 'record_key',
    'pk.fields'= 'id',
    'auto.create'= 'true',
    'auto.evolve'= 'true',
    'table.name.format'= 'payment_providers',
    'transforms'= 'flatten',
    'transforms.flatten.type'= 'org.apache.kafka.connect.transforms.Flatten$Value',
    'transforms.flatten.delimiter'= '_'
);


CREATE
SINK CONNECTOR SINK_CONNECTOR_PG_PAYMENTS_PAYMENTS
WITH (
    'connector.class'= 'io.confluent.connect.jdbc.JdbcSinkConnector',
    'tasks.max'= '1',
    'topics'= 'sink.pg.payments.payments',
    'connection.url'= 'jdbc=postgresql=//postgres-local=5432/payments',
    'connection.user'= 'postgres',
    'connection.password'= '1234',
    'insert.mode'= 'upsert',
    'pk.mode'= 'record_key',
    'pk.fields'= 'id',
    'auto.create'= 'true',
    'auto.evolve'= 'true',
    'table.name.format'= 'payments',
    'transforms'= 'flatten',
    'transforms.flatten.type'= 'org.apache.kafka.connect.transforms.Flatten$Value',
    'transforms.flatten.delimiter'= '_'
);

SHOW CONNECTORS;
