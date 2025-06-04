-- +goose Up

-- Table for Payment
CREATE TABLE Payment
(
    id          VARCHAR PRIMARY KEY,              -- Corresponds to string id
    code        VARCHAR          NOT NULL,        -- Corresponds to string code
    total_price DOUBLE PRECISION NOT NULL,        -- Corresponds to double totalPrice
    status      PaymentStatus    NOT NULL,        -- Corresponds to enum PaymentStatus
    provider_id VARCHAR REFERENCES Provider (id), -- FK to Provider table
    user_id     VARCHAR          NOT NULL,        -- Corresponds to string userId
    created_at  TIMESTAMP        NOT NULL,        -- Corresponds to google.protobuf.Timestamp created_at
    updated_at  TIMESTAMP        NOT NULL         -- Corresponds to google.protobuf.Timestamp updated_at
);

