-- +goose Up
CREATE TABLE products
(
    id           VARCHAR PRIMARY KEY               DEFAULT gen_random_uuid(),
    name         VARCHAR                  NOT NULL,
    description  VARCHAR                  NOT NULL,
    uom          VARCHAR                  NOT NULL,
    image        VARCHAR                  NOT NULL,
    price        NUMERIC                  NOT NULL DEFAULT 0,
    stock        INT                      NOT NULL,
    created_at   TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at   TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    discarded_at TIMESTAMP WITH TIME ZONE          DEFAULT NULL
);