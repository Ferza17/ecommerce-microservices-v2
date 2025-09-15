CREATE
SINK CONNECTOR SINK_CONNECTOR_PG_USERS_ROLES
WITH (
    'connector.class'= 'io.confluent.connect.jdbc.JdbcSinkConnector',
    'tasks.max'= '1',
    'topics'= 'sink.pg.users.roles',
    'connection.url'= 'jdbc=postgresql=//postgres-local=5432/users',
    'connection.user'= 'postgres',
    'connection.password'= '1234',
    'insert.mode'= 'upsert',
    'pk.mode'= 'record_key',
    'pk.fields'= 'id',
    'auto.create'= 'true',
    'auto.evolve'= 'true',
    'table.name.format'= 'roles',
    'table.whitelist' = 'roles',
    'transforms'= 'flatten',
    'transforms.flatten.type'= 'org.apache.kafka.connect.transforms.Flatten$Value',
    'transforms.flatten.delimiter'= '_'
);

CREATE
SINK CONNECTOR SINK_CONNECTOR_PG_USERS_USERS
WITH (
    'connector.class'= 'io.confluent.connect.jdbc.JdbcSinkConnector',
    'tasks.max'= '1',
    'topics'= 'sink.pg.users.users',
    'connection.url'= 'jdbc=postgresql=//postgres-local=5432/users',
    'connection.user'= 'postgres',
    'connection.password'= '1234',
    'insert.mode'= 'upsert',
    'pk.mode'= 'record_key',
    'pk.fields'= 'id',
    'auto.create'= 'true',
    'auto.evolve'= 'true',
    'table.name.format'= 'users',
    'table.whitelist' = 'users',
    'key.converter'= 'org.apache.kafka.connect.json.JsonConverter',
    'value.converter'= 'org.apache.kafka.connect.json.JsonConverter',
    'key.converter.schemas.enable'= 'false',
    'value.converter.schemas.enable'= 'false',
    'transforms'= 'flatten',
    'transforms.flatten.delimiter'= '_',
    'transforms.flatten.type'= 'org.apache.kafka.connect.transforms.Flatten$Value'
);


SHOW CONNECTORS;
