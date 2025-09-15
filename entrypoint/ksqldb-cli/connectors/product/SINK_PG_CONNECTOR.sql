CREATE
SINK CONNECTOR SINK_CONNECTOR_ES_PRODUCTS_PRODUCTS
WITH (
    'connector.class'= 'io.confluent.connect.elasticsearch.ElasticsearchSinkConnector',
    'topics'= 'sink.es.products.products',
    'connection.url'= 'http=//elasticsearch-local=9200/products',
    'connection.user'= '',
    'connection.password'= '',
    'type.name'= '_doc',
    'key.ignore'= 'true',
    'schema.ignore'= 'false',
    'behavior.on.null.values'= 'delete',
    'transforms'= 'flatten',
    'transforms.flatten.type'= 'org.apache.kafka.connect.transforms.Flatten$Value',
    'transforms.flatten.delimiter'= '_',
    'batch.size'= '1000',
    'batch.size.bytes'= '10000000',
    'linger.ms'= '1000',
    'flush.timeout.ms'= '10000',
    'max.in.flight.requests'= '1',
    'input.data.format'= 'JSON'
);

CREATE
SINK CONNECTOR SINK_CONNECTOR_PG_PRODUCTS_PRODUCTS
WITH (
    'connector.class'= 'io.confluent.connect.jdbc.JdbcSinkConnector',
    'tasks.max'= '1',
    'topics'= 'sink.pg.products.products',
    'connection.url'= 'jdbc=postgresql=//postgres-local=5432/products',
    'connection.user'= 'postgres',
    'connection.password'= '1234',
    'insert.mode'= 'upsert',
    'pk.mode'= 'record_key',
    'pk.fields'= 'id',
    'auto.create'= 'true',
    'auto.evolve'= 'true',
    'table.name.format'= 'products',
    'transforms'= 'flatten',
    'transforms.flatten.type'= 'org.apache.kafka.connect.transforms.Flatten$Value',
    'transforms.flatten.delimiter'= '_'
);

SHOW CONNECTORS;
