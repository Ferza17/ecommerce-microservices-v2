CREATE STREAM IF NOT EXISTS STREAM_USER_USER_CREATED (
    email        VARCHAR,
    name         VARCHAR,
    password     VARCHAR,
    role_id      INT
) WITH (
    KAFKA_TOPIC = 'snapshot-users-user_created',
    VALUE_FORMAT = 'JSON'
);

CREATE STREAM IF NOT EXISTS STREAM_USER_USER_LOGIN (
    email        VARCHAR,
    password     VARCHAR
) WITH (
    KAFKA_TOPIC = 'snapshot-users-user_login',
    VALUE_FORMAT = 'JSON'
);

SHOW STREAMS;