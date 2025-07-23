-- Your SQL goes here
CREATE TABLE shippings
(
    id                   VARCHAR PRIMARY KEY  DEFAULT gen_random_uuid(),
    created_by_id        VARCHAR     NOT NULL,
    order_id             VARCHAR     NOT NULL,
    _status               VARCHAR(15) NOT NULL DEFAULT 'PENDING',
    shipping_provider_id VARCHAR REFERENCES shipping_providers (id),

    created_at           TIMESTAMP            DEFAULT NOW(),
    updated_at           TIMESTAMP            DEFAULT NOW(),
    discarded_at         TIMESTAMP            DEFAULT NULL
);