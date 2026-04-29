-- +goose Up
-- +goose StatementBegin
CREATE TABLE category_stats (
    date Date,
    category_id UInt32,
    total_views UInt32,
    total_transactions UInt32,
    created_at DateTime DEFAULT now()
) ENGINE = MergeTree()
ORDER BY (date, category_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE category_stats;
-- +goose StatementEnd
