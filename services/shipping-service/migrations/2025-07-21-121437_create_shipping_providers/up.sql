-- Your SQL goes here
CREATE TABLE shipping_providers
(
    id           VARCHAR PRIMARY KEY               DEFAULT gen_random_uuid(),
    name         VARCHAR                  NOT NULL,
    created_at   TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at   TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    discarded_at TIMESTAMP WITH TIME ZONE          DEFAULT NULL
);