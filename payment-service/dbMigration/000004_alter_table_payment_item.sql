-- +goose Up

-- Add payment_id column and foreign key constraint to PaymentItem
ALTER TABLE payment_items
    ADD COLUMN payment_id VARCHAR NOT NULL, -- Foreign key to Payment table

    ADD CONSTRAINT fk_payment_item_payment
        FOREIGN KEY (payment_id)
            REFERENCES Payment (id)
            ON DELETE CASCADE -- Ensure cascading delete
            ON UPDATE CASCADE; -- Ensure updates cascade for references