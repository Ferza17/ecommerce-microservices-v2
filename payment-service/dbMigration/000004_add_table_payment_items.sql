-- +goose Up

-- Table for PaymentItem
CREATE TABLE PaymentItem
(
    id           VARCHAR PRIMARY KEY,       -- Corresponds to string id
    product_id   VARCHAR          NOT NULL, -- Corresponds to string productId
    amount       DOUBLE PRECISION NOT NULL, -- Corresponds to double amount
    qty          INTEGER          NOT NULL, -- Corresponds to int32 qty
    created_at   TIMESTAMP DEFAULT NOW(),   -- Represents `google.protobuf.Timestamp created_at` with DEFAULT NOW()
    updated_at   TIMESTAMP,                 -- Represents `google.protobuf.Timestamp updated_at`, auto-updated on change
    discarded_at TIMESTAMP
);
