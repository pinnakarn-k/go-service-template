DROP TABLE IF EXISTS transactions;

CREATE TABLE transactions (
    id BIGSERIAL PRIMARY KEY,
    account_no VARCHAR(20) NOT NULL,
    symbol VARCHAR(20) NOT NULL,
    side CHAR(1) NOT NULL,
    quantity NUMERIC(18, 4) NOT NULL,
    price NUMERIC(18, 4) NOT NULL,
    amount NUMERIC(18, 2) NOT NULL,
    fee NUMERIC(18, 2) NOT NULL,
    status CHAR(1) NOT NULL,
    trade_date TIMESTAMP NOT NULL,
    settlement_date DATE,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT transactions_side_check
        CHECK (side IN ('B', 'S')),

    CONSTRAINT transactions_status_check
        CHECK (status IN ('C', 'P', 'F')),

    CONSTRAINT transactions_quantity_check
        CHECK (quantity > 0),

    CONSTRAINT transactions_price_check
        CHECK (price > 0),

    CONSTRAINT transactions_amount_check
        CHECK (amount >= 0),

    CONSTRAINT transactions_fee_check
        CHECK (fee >= 0)
);

WITH generated_transactions AS (
    SELECT
        gs,

        'ACC' || LPAD(
            (((gs - 1) % 20) + 1)::text,
            6,
            '0'
        ) AS account_no,

        (
            ARRAY[
                'AOT',
                'PTT',
                'KBANK',
                'SCB',
                'CPALL',
                'ADVANC',
                'TRUE',
                'BDMS',
                'KTC',
                'BEM'
            ]
        )[1 + ((gs - 1) % 10)] AS symbol,

        (
            ARRAY['B', 'S']
        )[1 + ((gs - 1) % 2)] AS side,

        (
            10 + ((gs * 7) % 491)
        )::NUMERIC(18, 4) AS quantity,

        ROUND(
            (
                20
                + ((gs * 13) % 380)
                + ((gs % 100)::NUMERIC / 100)
            )::NUMERIC,
            4
        ) AS price,

        (
            ARRAY['C', 'P', 'F']
        )[1 + ((gs - 1) % 3)] AS status,

        (
            TIMESTAMP '2026-04-01 09:00:00'
            + (((gs * 17) % 120) * INTERVAL '1 day')
            + (((gs * 11) % 480) * INTERVAL '1 minute')
        ) AS trade_date

    FROM generate_series(1, 1000) AS gs
)

INSERT INTO transactions (
    account_no,
    symbol,
    side,
    quantity,
    price,
    amount,
    fee,
    status,
    trade_date,
    settlement_date,
    created_at,
    updated_at
)
SELECT
    account_no,
    symbol,
    side,
    quantity,
    price,

    ROUND(
        quantity * price,
        2
    ) AS amount,

    ROUND(
        quantity * price * 0.0015,
        2
    ) AS fee,

    status,
    trade_date,

    CASE
        WHEN status = 'C'
            THEN trade_date::DATE + 2
        ELSE NULL
    END AS settlement_date,

    trade_date + INTERVAL '5 minutes' AS created_at,
    trade_date + INTERVAL '10 minutes' AS updated_at

FROM generated_transactions;

CREATE INDEX idx_transactions_trade_date
    ON transactions (trade_date);

CREATE INDEX idx_transactions_symbol
    ON transactions (symbol);

CREATE INDEX idx_transactions_account_no
    ON transactions (account_no);

CREATE INDEX idx_transactions_side
    ON transactions (side);

CREATE INDEX idx_transactions_status
    ON transactions (status);

CREATE INDEX idx_transactions_search
    ON transactions (
        trade_date,
        symbol,
        side,
        status
    );