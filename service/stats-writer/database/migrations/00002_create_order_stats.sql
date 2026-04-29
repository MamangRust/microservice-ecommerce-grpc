-- +goose Up
-- +goose StatementBegin
CREATE TABLE order_stats (
    date Date,
    order_id UInt32,
    total_revenue UInt64,
    items_count UInt32,
    created_at DateTime DEFAULT now()
) ENGINE = MergeTree()
ORDER BY (date, order_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE order_stats;
-- +goose StatementEnd
