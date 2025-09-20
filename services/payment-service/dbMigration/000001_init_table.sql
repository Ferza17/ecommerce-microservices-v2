-- +goose Up
CREATE TABLE payment_providers
(
    id           VARCHAR PRIMARY KEY,                    -- Corresponds to string id
    name         VARCHAR     NOT NULL,                   -- Corresponds to string name
    method       VARCHAR(50) NOT NULL,                   -- Corresponds to enum ProviderMethod
    created_at   TIMESTAMP WITH TIME ZONE DEFAULT NOW(), -- Represents `google.protobuf.Timestamp created_at` with DEFAULT NOW()
    updated_at   TIMESTAMP WITH TIME ZONE DEFAULT NOW(), -- Represents `google.protobuf.Timestamp updated_at`, auto-updated on change
    discarded_at TIMESTAMP WITH TIME ZONE DEFAULT NULL
);


-- Table for Payment
CREATE TABLE payments
(
    id           VARCHAR PRIMARY KEY,                       -- Corresponds to string id
    code         VARCHAR          NOT NULL,                 -- Corresponds to string code
    total_price  DOUBLE PRECISION NOT NULL,                 -- Corresponds to double totalPrice
    status       VARCHAR(50)      NOT NULL,                 -- Corresponds to enum PaymentStatus
    provider_id  VARCHAR REFERENCES payment_providers (id), -- FK to Provider table
    user_id      VARCHAR          NOT NULL,                 -- Corresponds to string userId
    created_at   TIMESTAMP WITH TIME ZONE DEFAULT NOW(),    -- Represents `google.protobuf.Timestamp created_at` with DEFAULT NOW()
    updated_at   TIMESTAMP WITH TIME ZONE DEFAULT NOW(),    -- Represents `google.protobuf.Timestamp updated_at`, auto-updated on change
    discarded_at TIMESTAMP WITH TIME ZONE
);


-- Table for PaymentItem
CREATE TABLE payment_items
(
    id           VARCHAR PRIMARY KEY,                    -- Corresponds to string id
    product_id   VARCHAR          NOT NULL,              -- Corresponds to string productId
    amount       DOUBLE PRECISION NOT NULL,              -- Corresponds to double amount
    qty          INTEGER          NOT NULL,              -- Corresponds to int32 qty
    payment_id   VARCHAR REFERENCES payments (id),
    created_at   TIMESTAMP WITH TIME ZONE DEFAULT NOW(), -- Represents `google.protobuf.Timestamp created_at` with DEFAULT NOW()
    updated_at   TIMESTAMP WITH TIME ZONE DEFAULT NOW(), -- Represents `google.protobuf.Timestamp updated_at`, auto-updated on change
    discarded_at TIMESTAMP WITH TIME ZONE
);