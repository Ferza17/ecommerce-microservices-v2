-- +goose Up

-- Table for Payment
CREATE TABLE payments
(
    id           VARCHAR PRIMARY KEY,               -- Corresponds to string id
    code         VARCHAR          NOT NULL,         -- Corresponds to string code
    total_price  DOUBLE PRECISION NOT NULL,         -- Corresponds to double totalPrice
    status       INTEGER          NOT NULL,         -- Corresponds to enum PaymentStatus
    provider_id  VARCHAR REFERENCES providers (id), -- FK to Provider table
    user_id      VARCHAR          NOT NULL,         -- Corresponds to string userId
    created_at   TIMESTAMP DEFAULT NOW(),           -- Represents `google.protobuf.Timestamp created_at` with DEFAULT NOW()
    updated_at   TIMESTAMP,                         -- Represents `google.protobuf.Timestamp updated_at`, auto-updated on change
    discarded_at TIMESTAMP
);

