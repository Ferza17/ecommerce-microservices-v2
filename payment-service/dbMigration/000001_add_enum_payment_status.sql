-- +goose Up

CREATE TYPE PaymentStatus AS ENUM (
    'PENDING', -- 0
    'PARTIAL', -- 1
    'SUCCESS', -- 2
    'FAILED' -- 3
    );

