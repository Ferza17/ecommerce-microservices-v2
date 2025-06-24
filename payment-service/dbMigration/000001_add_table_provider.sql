-- +goose Up
-- Table for Provider
CREATE TABLE providers
(
    id           VARCHAR PRIMARY KEY,     -- Corresponds to string id
    name         VARCHAR NOT NULL,        -- Corresponds to string name
    method       INTEGER NOT NULL,        -- Corresponds to enum ProviderMethod
    created_at   TIMESTAMP DEFAULT NOW(), -- Represents `google.protobuf.Timestamp created_at` with DEFAULT NOW()
    updated_at   TIMESTAMP,               -- Represents `google.protobuf.Timestamp updated_at`, auto-updated on change
    discarded_at TIMESTAMP
);