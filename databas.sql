-- Active: 1692898566087@@127.0.0.1@5432@postgres
CREATE TABLE accounts (
    id SERIAL PRIMARY KEY,
    customer_id INT NOT NULL,
    account_limit NUMERIC,
    per_transaction_limit NUMERIC,
    last_account_limit NUMERIC,
    last_per_transaction_limit NUMERIC,
    account_limit_update_time TIMESTAMP,
    per_transaction_limit_update_time TIMESTAMP,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);
CREATE SEQUENCE customer_id_seq;

select * from accounts;

CREATE TABLE limit_offers (
    id SERIAL PRIMARY KEY,
    account_id INT NOT NULL,
    limit_type VARCHAR(20) NOT NULL,
    new_limit NUMERIC NOT NULL,
    offer_activation_time TIMESTAMP NOT NULL,
    offer_expiry_time TIMESTAMP NOT NULL,
    status VARCHAR(20) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);