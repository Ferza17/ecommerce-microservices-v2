-- +goose Up
CREATE TABLE users
(
    id           VARCHAR PRIMARY KEY DEFAULT gen_random_uuid(),
    name         VARCHAR NOT NULL,
    email        VARCHAR NOT NULL UNIQUE,
    password     VARCHAR NOT NULL,
    created_at   TIMESTAMP           DEFAULT NOW(),
    updated_at   TIMESTAMP           DEFAULT NOW(),
    discarded_at TIMESTAMP           DEFAULT NULL
);