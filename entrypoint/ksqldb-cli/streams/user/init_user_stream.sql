CREATE STREAM IF NOT EXISTS STREAM_USER_USER_CREATED (
    email        VARCHAR,
    name         VARCHAR,
    password     VARCHAR,
    role_id      INT
) WITH (
    KAFKA_TOPIC = 'users.user_created.snapshot',
    VALUE_FORMAT = 'JSON'
);

CREATE STREAM IF NOT EXISTS STREAM_USER_USER_LOGIN (
    email        VARCHAR,
    password     VARCHAR
) WITH (
    KAFKA_TOPIC = 'users.user_login.snapshot',
    VALUE_FORMAT = 'JSON'
);

CREATE STREAM IF NOT EXISTS STREAM_SINK_CONNECTOR_PG_USERS_ROLES (
    id VARCHAR KEY,
    role VARCHAR,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
) WITH (
    KAFKA_TOPIC='sink.pg.users.roles',
    VALUE_FORMAT='JSON'
);

CREATE STREAM STREAM_SINK_CONNECTOR_PG_USERS_USERS (
    id VARCHAR KEY,
    name VARCHAR,
    email VARCHAR,
    password VARCHAR,
    is_verified BOOLEAN,
    role_id VARCHAR,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    discarded_at TIMESTAMP
) WITH (
    KAFKA_TOPIC='sink.pg.users.users',
    VALUE_FORMAT='JSON'
);


SHOW STREAMS;