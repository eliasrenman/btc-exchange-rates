CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Create the exchange_rates table
CREATE TABLE exchange_rates (
    uuid UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    currency VARCHAR(255) NOT NULL,
    rate FLOAT8 NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Create an index on the currency column
DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_indexes WHERE tablename = 'exchange_rates' AND indexname = 'idx_currency') THEN
        CREATE INDEX idx_currency ON exchange_rates(currency);
    END IF;
END
$$;
