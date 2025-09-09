-- Your SQL goes here
CREATE TABLE shipping_providers
(
    id           VARCHAR PRIMARY KEY DEFAULT gen_random_uuid(),
    name         VARCHAR   NOT NULL,
    created_at   TIMESTAMP NOT NULL  DEFAULT NOW(),
    updated_at   TIMESTAMP NOT NULL  DEFAULT NOW(),
    discarded_at TIMESTAMP           DEFAULT NULL
);