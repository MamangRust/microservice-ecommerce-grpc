-- +goose Up
-- +goose StatementBegin
-- Category Stats expansion
ALTER TABLE category_stats ADD COLUMN IF NOT EXISTS category_name String AFTER category_id;
ALTER TABLE category_stats ADD COLUMN IF NOT EXISTS items_sold UInt32 AFTER total_transactions;
ALTER TABLE category_stats ADD COLUMN IF NOT EXISTS order_count UInt32 AFTER items_sold;

-- Order Stats expansion
ALTER TABLE order_stats ADD COLUMN IF NOT EXISTS total_revenue UInt64 AFTER order_id;
ALTER TABLE order_stats ADD COLUMN IF NOT EXISTS total_items_sold UInt32 AFTER total_revenue;
ALTER TABLE order_stats ADD COLUMN IF NOT EXISTS active_cashiers UInt32 AFTER total_items_sold;
ALTER TABLE order_stats ADD COLUMN IF NOT EXISTS unique_products_sold UInt32 AFTER active_cashiers;

-- Transaction Stats expansion
ALTER TABLE transaction_stats ADD COLUMN IF NOT EXISTS total_amount UInt64 AFTER amount;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE category_stats DROP COLUMN IF EXISTS category_name;
ALTER TABLE category_stats DROP COLUMN IF EXISTS items_sold;
ALTER TABLE category_stats DROP COLUMN IF EXISTS order_count;

ALTER TABLE order_stats DROP COLUMN IF EXISTS total_revenue;
ALTER TABLE order_stats DROP COLUMN IF EXISTS total_items_sold;
ALTER TABLE order_stats DROP COLUMN IF EXISTS active_cashiers;
ALTER TABLE order_stats DROP COLUMN IF EXISTS unique_products_sold;

ALTER TABLE transaction_stats DROP COLUMN IF EXISTS total_amount;
-- +goose StatementEnd
