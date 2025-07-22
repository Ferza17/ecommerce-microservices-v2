-- Your SQL goes here
CREATE TABLE shipping_providers
(
    id           VARCHAR PRIMARY KEY DEFAULT gen_random_uuid(),
    name         VARCHAR NULL,
    created_at   TIMESTAMP           DEFAULT NOW(),
    updated_at   TIMESTAMP           DEFAULT NOW(),
    discarded_at TIMESTAMP           DEFAULT NULL
);