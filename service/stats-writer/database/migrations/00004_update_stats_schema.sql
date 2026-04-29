-- +goose Up
-- +goose StatementBegin
ALTER TABLE category_stats ADD COLUMN IF NOT EXISTS merchant_id UInt32 AFTER category_id;
ALTER TABLE category_stats ADD COLUMN IF NOT EXISTS total_price UInt64 AFTER total_transactions;

ALTER TABLE order_stats ADD COLUMN IF NOT EXISTS merchant_id UInt32 AFTER order_id;

ALTER TABLE transaction_stats ADD COLUMN IF NOT EXISTS merchant_id UInt32 AFTER transaction_id;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE category_stats DROP COLUMN IF EXISTS merchant_id;
ALTER TABLE category_stats DROP COLUMN IF EXISTS total_price;

ALTER TABLE order_stats DROP COLUMN IF EXISTS merchant_id;

ALTER TABLE transaction_stats DROP COLUMN IF EXISTS merchant_id;
-- +goose StatementEnd
