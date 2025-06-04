-- +goose Up

-- Table for PaymentItem
CREATE TABLE PaymentItem
(
    id         VARCHAR PRIMARY KEY,       -- Corresponds to string id
    product_id VARCHAR          NOT NULL, -- Corresponds to string productId
    amount     DOUBLE PRECISION NOT NULL, -- Corresponds to double amount
    qty        INTEGER          NOT NULL, -- Corresponds to int32 qty
    created_at TIMESTAMP        NOT NULL, -- Corresponds to google.protobuf.Timestamp crated_at
    updated_at TIMESTAMP        NOT NULL  -- Corresponds to google.protobuf.Timestamp updated_at
);
