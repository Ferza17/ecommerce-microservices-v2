-- +goose Up

-- Enum for ProviderMethod
CREATE TYPE ProviderMethod AS ENUM (
    'BANK', -- 0
    'CRYPTO_CURRENCY', -- 1
    'DEBIT', -- 2
    'CREDIT', -- 3
    'CASH_ON_DELIVERY' -- 4
    );

-- Table for Provider
CREATE TABLE Provider
(
    id           VARCHAR PRIMARY KEY,     -- Corresponds to string id
    name         VARCHAR        NOT NULL, -- Corresponds to string name
    method       ProviderMethod NOT NULL, -- Corresponds to enum ProviderMethod
    created_at   TIMESTAMP DEFAULT NOW(), -- Represents `google.protobuf.Timestamp created_at` with DEFAULT NOW()
    updated_at   TIMESTAMP,               -- Represents `google.protobuf.Timestamp updated_at`, auto-updated on change
    discarded_at TIMESTAMP
);