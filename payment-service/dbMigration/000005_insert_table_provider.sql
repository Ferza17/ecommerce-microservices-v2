-- +goose Up

INSERT INTO Provider (id, name, method, created_at, updated_at, discarded_at)
VALUES ('1', 'Bank of America', 'BANK', NOW(), NULL, NULL),
       ('2', 'CoinBase', 'CRYPTO_CURRENCY', NOW(), NULL, NULL),
       ('3', 'Visa Debit Services', 'DEBIT', NOW(), NULL, NULL),
       ('4', 'MasterCard Credit', 'CREDIT', NOW(), NULL, NULL),
       ('5', 'Amazon Cash on Delivery', 'CASH_ON_DELIVERY', NOW(), NULL, NULL);
