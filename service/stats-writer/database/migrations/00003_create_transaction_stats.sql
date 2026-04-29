-- +goose Up
-- +goose StatementBegin
CREATE TABLE transaction_stats (
    date Date,
    transaction_id UInt32,
    payment_method String,
    amount UInt64,
    status String,
    created_at DateTime DEFAULT now()
) ENGINE = MergeTree()
ORDER BY (date, transaction_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE transaction_stats;
-- +goose StatementEnd
