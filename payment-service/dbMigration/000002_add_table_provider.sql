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
    id     VARCHAR PRIMARY KEY,     -- Corresponds to string id
    name   VARCHAR        NOT NULL, -- Corresponds to string name
    method ProviderMethod NOT NULL  -- Corresponds to enum ProviderMethod
);