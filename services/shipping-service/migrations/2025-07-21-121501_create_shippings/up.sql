-- Your SQL goes here
CREATE TABLE shippings
(
    id                   VARCHAR(255) PRIMARY KEY,
    user_id              VARCHAR(255)             NOT NULL,
    payment_id           VARCHAR(255)             NOT NULL,
    shipping_provider_id VARCHAR                  NOT NULL REFERENCES shipping_providers (id),
    created_at           TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at           TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    discarded_at         TIMESTAMP WITH TIME ZONE
);