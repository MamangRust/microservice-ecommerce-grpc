package repository

import (
	"context"
	chDriver "github.com/ClickHouse/clickhouse-go/v2"
	"github.com/MamangRust/monolith-ecommerce-shared/pb"
)

type clickhouseReaderRepository struct {
	conn chDriver.Conn
}

func NewClickHouseReaderRepository(conn chDriver.Conn) Repository {
	return &clickhouseReaderRepository{
		conn: conn,
	}
}

// Category Stats Implementation

func (r *clickhouseReaderRepository) FindMonthlyTotalPrices(ctx context.Context, req *pb.FindYearMonthTotalPrices) ([]*pb.CategoriesMonthlyTotalPriceResponse, error) {
	query := `SELECT toYear(date) as y, formatDateTime(date, '%M') as m, sum(total_price) FROM category_stats WHERE toYear(date) = ? AND toMonth(date) = ? GROUP BY y, m`
	rows, err := r.conn.Query(ctx, query, req.Year, req.Month)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []*pb.CategoriesMonthlyTotalPriceResponse
	for rows.Next() {
		var item pb.CategoriesMonthlyTotalPriceResponse
		if err := rows.Scan(&item.Year, &item.Month, &item.TotalRevenue); err != nil {
			return nil, err
		}
		result = append(result, &item)
	}
	return result, nil
}

func (r *clickhouseReaderRepository) FindYearlyTotalPrices(ctx context.Context, req *pb.FindYearTotalPrices) ([]*pb.CategoriesYearlyTotalPriceResponse, error) {
	query := `SELECT toYear(date) as y, sum(total_price) FROM category_stats WHERE toYear(date) = ? GROUP BY y`
	rows, err := r.conn.Query(ctx, query, req.Year)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []*pb.CategoriesYearlyTotalPriceResponse
	for rows.Next() {
		var item pb.CategoriesYearlyTotalPriceResponse
		if err := rows.Scan(&item.Year, &item.TotalRevenue); err != nil {
			return nil, err
		}
		result = append(result, &item)
	}
	return result, nil
}

func (r *clickhouseReaderRepository) FindMonthPrice(ctx context.Context, req *pb.FindYearCategory) ([]*pb.CategoryMonthPriceResponse, error) {
	query := `SELECT formatDateTime(date, '%b') as m, category_id, any(category_name), sum(order_count), sum(items_sold), sum(total_price) 
			  FROM category_stats WHERE toYear(date) = ? GROUP BY m, category_id ORDER BY m, category_id`
	rows, err := r.conn.Query(ctx, query, req.Year)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []*pb.CategoryMonthPriceResponse
	for rows.Next() {
		var item pb.CategoryMonthPriceResponse
		if err := rows.Scan(&item.Month, &item.CategoryId, &item.CategoryName, &item.OrderCount, &item.ItemsSold, &item.TotalRevenue); err != nil {
			return nil, err
		}
		result = append(result, &item)
	}
	return result, nil
}

func (r *clickhouseReaderRepository) FindYearPrice(ctx context.Context, req *pb.FindYearCategory) ([]*pb.CategoryYearPriceResponse, error) {
	query := `SELECT toYear(date) as y, category_id, any(category_name), sum(order_count), sum(items_sold), sum(total_price), count(DISTINCT product_id)
			  FROM category_stats WHERE toYear(date) = ? GROUP BY y, category_id ORDER BY category_id`
	// Note: using count(DISTINCT product_id) as placeholder for unique_products_sold if column doesn't exist yet, but I'll update schema accordingly.
	rows, err := r.conn.Query(ctx, query, req.Year)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []*pb.CategoryYearPriceResponse
	for rows.Next() {
		var item pb.CategoryYearPriceResponse
		if err := rows.Scan(&item.Year, &item.CategoryId, &item.CategoryName, &item.OrderCount, &item.ItemsSold, &item.TotalRevenue, &item.UniqueProductsSold); err != nil {
			return nil, err
		}
		result = append(result, &item)
	}
	return result, nil
}

// Category Stats By ID Implementation (Similar but filtered by category_id)

func (r *clickhouseReaderRepository) FindMonthlyTotalPricesById(ctx context.Context, req *pb.FindYearMonthTotalPriceById) ([]*pb.CategoriesMonthlyTotalPriceResponse, error) {
	query := `SELECT toYear(date) as y, formatDateTime(date, '%M') as m, sum(total_price) FROM category_stats WHERE toYear(date) = ? AND toMonth(date) = ? AND category_id = ? GROUP BY y, m`
	rows, err := r.conn.Query(ctx, query, req.Year, req.Month, req.CategoryId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []*pb.CategoriesMonthlyTotalPriceResponse
	for rows.Next() {
		var item pb.CategoriesMonthlyTotalPriceResponse
		if err := rows.Scan(&item.Year, &item.Month, &item.TotalRevenue); err != nil {
			return nil, err
		}
		result = append(result, &item)
	}
	return result, nil
}

func (r *clickhouseReaderRepository) FindYearlyTotalPricesById(ctx context.Context, req *pb.FindYearTotalPriceById) ([]*pb.CategoriesYearlyTotalPriceResponse, error) {
	query := `SELECT toYear(date) as y, sum(total_price) FROM category_stats WHERE toYear(date) = ? AND category_id = ? GROUP BY y`
	rows, err := r.conn.Query(ctx, query, req.Year, req.CategoryId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []*pb.CategoriesYearlyTotalPriceResponse
	for rows.Next() {
		var item pb.CategoriesYearlyTotalPriceResponse
		if err := rows.Scan(&item.Year, &item.TotalRevenue); err != nil {
			return nil, err
		}
		result = append(result, &item)
	}
	return result, nil
}

func (r *clickhouseReaderRepository) FindMonthPriceById(ctx context.Context, req *pb.FindYearCategoryById) ([]*pb.CategoryMonthPriceResponse, error) {
	query := `SELECT formatDateTime(date, '%b') as m, category_id, any(category_name), sum(order_count), sum(items_sold), sum(total_price) 
			  FROM category_stats WHERE toYear(date) = ? AND category_id = ? GROUP BY m, category_id ORDER BY m`
	rows, err := r.conn.Query(ctx, query, req.Year, req.CategoryId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []*pb.CategoryMonthPriceResponse
	for rows.Next() {
		var item pb.CategoryMonthPriceResponse
		if err := rows.Scan(&item.Month, &item.CategoryId, &item.CategoryName, &item.OrderCount, &item.ItemsSold, &item.TotalRevenue); err != nil {
			return nil, err
		}
		result = append(result, &item)
	}
	return result, nil
}

func (r *clickhouseReaderRepository) FindYearPriceById(ctx context.Context, req *pb.FindYearCategoryById) ([]*pb.CategoryYearPriceResponse, error) {
	query := `SELECT toYear(date) as y, category_id, any(category_name), sum(order_count), sum(items_sold), sum(total_price), count(DISTINCT product_id)
			  FROM category_stats WHERE toYear(date) = ? AND category_id = ? GROUP BY y, category_id`
	rows, err := r.conn.Query(ctx, query, req.Year, req.CategoryId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []*pb.CategoryYearPriceResponse
	for rows.Next() {
		var item pb.CategoryYearPriceResponse
		if err := rows.Scan(&item.Year, &item.CategoryId, &item.CategoryName, &item.OrderCount, &item.ItemsSold, &item.TotalRevenue, &item.UniqueProductsSold); err != nil {
			return nil, err
		}
		result = append(result, &item)
	}
	return result, nil
}

// Category Stats By Merchant Implementation

func (r *clickhouseReaderRepository) FindMonthlyTotalPricesByMerchant(ctx context.Context, req *pb.FindYearMonthTotalPriceByMerchant) ([]*pb.CategoriesMonthlyTotalPriceResponse, error) {
	query := `SELECT toYear(date) as y, formatDateTime(date, '%M') as m, sum(total_price) FROM category_stats WHERE toYear(date) = ? AND toMonth(date) = ? AND merchant_id = ? GROUP BY y, m`
	rows, err := r.conn.Query(ctx, query, req.Year, req.Month, req.MerchantId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []*pb.CategoriesMonthlyTotalPriceResponse
	for rows.Next() {
		var item pb.CategoriesMonthlyTotalPriceResponse
		if err := rows.Scan(&item.Year, &item.Month, &item.TotalRevenue); err != nil {
			return nil, err
		}
		result = append(result, &item)
	}
	return result, nil
}

func (r *clickhouseReaderRepository) FindYearlyTotalPricesByMerchant(ctx context.Context, req *pb.FindYearTotalPriceByMerchant) ([]*pb.CategoriesYearlyTotalPriceResponse, error) {
	query := `SELECT toYear(date) as y, sum(total_price) FROM category_stats WHERE toYear(date) = ? AND merchant_id = ? GROUP BY y`
	rows, err := r.conn.Query(ctx, query, req.Year, req.MerchantId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []*pb.CategoriesYearlyTotalPriceResponse
	for rows.Next() {
		var item pb.CategoriesYearlyTotalPriceResponse
		if err := rows.Scan(&item.Year, &item.TotalRevenue); err != nil {
			return nil, err
		}
		result = append(result, &item)
	}
	return result, nil
}

func (r *clickhouseReaderRepository) FindMonthPriceByMerchant(ctx context.Context, req *pb.FindYearCategoryByMerchant) ([]*pb.CategoryMonthPriceResponse, error) {
	query := `SELECT formatDateTime(date, '%b') as m, category_id, any(category_name), sum(order_count), sum(items_sold), sum(total_price) 
			  FROM category_stats WHERE toYear(date) = ? AND merchant_id = ? GROUP BY m, category_id ORDER BY m, category_id`
	rows, err := r.conn.Query(ctx, query, req.Year, req.MerchantId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []*pb.CategoryMonthPriceResponse
	for rows.Next() {
		var item pb.CategoryMonthPriceResponse
		if err := rows.Scan(&item.Month, &item.CategoryId, &item.CategoryName, &item.OrderCount, &item.ItemsSold, &item.TotalRevenue); err != nil {
			return nil, err
		}
		result = append(result, &item)
	}
	return result, nil
}

func (r *clickhouseReaderRepository) FindYearPriceByMerchant(ctx context.Context, req *pb.FindYearCategoryByMerchant) ([]*pb.CategoryYearPriceResponse, error) {
	query := `SELECT toYear(date) as y, category_id, any(category_name), sum(order_count), sum(items_sold), sum(total_price), count(DISTINCT product_id)
			  FROM category_stats WHERE toYear(date) = ? AND merchant_id = ? GROUP BY y, category_id ORDER BY category_id`
	rows, err := r.conn.Query(ctx, query, req.Year, req.MerchantId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []*pb.CategoryYearPriceResponse
	for rows.Next() {
		var item pb.CategoryYearPriceResponse
		if err := rows.Scan(&item.Year, &item.CategoryId, &item.CategoryName, &item.OrderCount, &item.ItemsSold, &item.TotalRevenue, &item.UniqueProductsSold); err != nil {
			return nil, err
		}
		result = append(result, &item)
	}
	return result, nil
}

// Order Stats Implementation

func (r *clickhouseReaderRepository) FindMonthlyTotalRevenue(ctx context.Context, req *pb.FindYearMonthTotalRevenue) ([]*pb.OrderMonthlyTotalRevenueResponse, error) {
	query := `SELECT toYear(date) as y, formatDateTime(date, '%M') as m, sum(order_count), sum(total_revenue), sum(total_items_sold) 
			  FROM order_stats WHERE toYear(date) = ? AND toMonth(date) = ? GROUP BY y, m`
	rows, err := r.conn.Query(ctx, query, req.Year, req.Month)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []*pb.OrderMonthlyTotalRevenueResponse
	for rows.Next() {
		var item pb.OrderMonthlyTotalRevenueResponse
		if err := rows.Scan(&item.Year, &item.Month, &item.OrderCount, &item.TotalRevenue, &item.TotalItemsSold); err != nil {
			return nil, err
		}
		result = append(result, &item)
	}
	return result, nil
}

func (r *clickhouseReaderRepository) FindYearlyTotalRevenue(ctx context.Context, req *pb.FindYearTotalRevenue) ([]*pb.OrderYearlyTotalRevenueResponse, error) {
	query := `SELECT toYear(date) as y, sum(order_count), sum(total_revenue), sum(total_items_sold), sum(active_cashiers), sum(unique_products_sold) 
			  FROM order_stats WHERE toYear(date) = ? GROUP BY y`
	rows, err := r.conn.Query(ctx, query, req.Year)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []*pb.OrderYearlyTotalRevenueResponse
	for rows.Next() {
		var item pb.OrderYearlyTotalRevenueResponse
		if err := rows.Scan(&item.Year, &item.OrderCount, &item.TotalRevenue, &item.TotalItemsSold, &item.ActiveCashiers, &item.UniqueProductsSold); err != nil {
			return nil, err
		}
		result = append(result, &item)
	}
	return result, nil
}

func (r *clickhouseReaderRepository) FindMonthlyRevenue(ctx context.Context, req *pb.FindYearOrder) ([]*pb.OrderMonthlyResponse, error) {
	query := `SELECT formatDateTime(date, '%b') as m, sum(order_count), sum(total_revenue), sum(total_items_sold) 
			  FROM order_stats WHERE toYear(date) = ? GROUP BY m ORDER BY m`
	rows, err := r.conn.Query(ctx, query, req.Year)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []*pb.OrderMonthlyResponse
	for rows.Next() {
		var item pb.OrderMonthlyResponse
		if err := rows.Scan(&item.Month, &item.OrderCount, &item.TotalRevenue, &item.TotalItemsSold); err != nil {
			return nil, err
		}
		result = append(result, &item)
	}
	return result, nil
}

func (r *clickhouseReaderRepository) FindYearlyRevenue(ctx context.Context, req *pb.FindYearOrder) ([]*pb.OrderYearlyResponse, error) {
	query := `SELECT toYear(date) as y, sum(order_count), sum(total_revenue), sum(total_items_sold), sum(active_cashiers), sum(unique_products_sold) 
			  FROM order_stats WHERE toYear(date) = ? GROUP BY y`
	rows, err := r.conn.Query(ctx, query, req.Year)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []*pb.OrderYearlyResponse
	for rows.Next() {
		var item pb.OrderYearlyResponse
		if err := rows.Scan(&item.Year, &item.OrderCount, &item.TotalRevenue, &item.TotalItemsSold, &item.ActiveCashiers, &item.UniqueProductsSold); err != nil {
			return nil, err
		}
		result = append(result, &item)
	}
	return result, nil
}

// Order Stats By Merchant Implementation

func (r *clickhouseReaderRepository) FindMonthlyTotalRevenueByMerchant(ctx context.Context, req *pb.FindYearMonthTotalRevenueByMerchant) ([]*pb.OrderMonthlyTotalRevenueResponse, error) {
	query := `SELECT toYear(date) as y, formatDateTime(date, '%M') as m, sum(order_count), sum(total_revenue), sum(total_items_sold) 
			  FROM order_stats WHERE toYear(date) = ? AND toMonth(date) = ? AND merchant_id = ? GROUP BY y, m`
	rows, err := r.conn.Query(ctx, query, req.Year, req.Month, req.MerchantId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []*pb.OrderMonthlyTotalRevenueResponse
	for rows.Next() {
		var item pb.OrderMonthlyTotalRevenueResponse
		if err := rows.Scan(&item.Year, &item.Month, &item.OrderCount, &item.TotalRevenue, &item.TotalItemsSold); err != nil {
			return nil, err
		}
		result = append(result, &item)
	}
	return result, nil
}

func (r *clickhouseReaderRepository) FindYearlyTotalRevenueByMerchant(ctx context.Context, req *pb.FindYearTotalRevenueByMerchant) ([]*pb.OrderYearlyTotalRevenueResponse, error) {
	query := `SELECT toYear(date) as y, sum(order_count), sum(total_revenue), sum(total_items_sold), sum(active_cashiers), sum(unique_products_sold) 
			  FROM order_stats WHERE toYear(date) = ? AND merchant_id = ? GROUP BY y`
	rows, err := r.conn.Query(ctx, query, req.Year, req.MerchantId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []*pb.OrderYearlyTotalRevenueResponse
	for rows.Next() {
		var item pb.OrderYearlyTotalRevenueResponse
		if err := rows.Scan(&item.Year, &item.OrderCount, &item.TotalRevenue, &item.TotalItemsSold, &item.ActiveCashiers, &item.UniqueProductsSold); err != nil {
			return nil, err
		}
		result = append(result, &item)
	}
	return result, nil
}

func (r *clickhouseReaderRepository) FindMonthlyRevenueByMerchant(ctx context.Context, req *pb.FindYearOrderByMerchant) ([]*pb.OrderMonthlyResponse, error) {
	query := `SELECT formatDateTime(date, '%b') as m, sum(order_count), sum(total_revenue), sum(total_items_sold) 
			  FROM order_stats WHERE toYear(date) = ? AND merchant_id = ? GROUP BY m ORDER BY m`
	rows, err := r.conn.Query(ctx, query, req.Year, req.MerchantId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []*pb.OrderMonthlyResponse
	for rows.Next() {
		var item pb.OrderMonthlyResponse
		if err := rows.Scan(&item.Month, &item.OrderCount, &item.TotalRevenue, &item.TotalItemsSold); err != nil {
			return nil, err
		}
		result = append(result, &item)
	}
	return result, nil
}

func (r *clickhouseReaderRepository) FindYearlyRevenueByMerchant(ctx context.Context, req *pb.FindYearOrderByMerchant) ([]*pb.OrderYearlyResponse, error) {
	query := `SELECT toYear(date) as y, sum(order_count), sum(total_revenue), sum(total_items_sold), sum(active_cashiers), sum(unique_products_sold) 
			  FROM order_stats WHERE toYear(date) = ? AND merchant_id = ? GROUP BY y`
	rows, err := r.conn.Query(ctx, query, req.Year, req.MerchantId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []*pb.OrderYearlyResponse
	for rows.Next() {
		var item pb.OrderYearlyResponse
		if err := rows.Scan(&item.Year, &item.OrderCount, &item.TotalRevenue, &item.TotalItemsSold, &item.ActiveCashiers, &item.UniqueProductsSold); err != nil {
			return nil, err
		}
		result = append(result, &item)
	}
	return result, nil
}

// Transaction Stats Implementation

func (r *clickhouseReaderRepository) GetMonthlyAmountSuccess(ctx context.Context, req *pb.MonthAmountTransactionRequest) ([]*pb.TransactionMonthlyAmountSuccess, error) {
	query := `SELECT toYear(date) as y, formatDateTime(date, '%M') as m, sum(total_success), sum(total_amount) 
			  FROM transaction_stats WHERE toYear(date) = ? AND toMonth(date) = ? AND status = 'success' GROUP BY y, m`
	rows, err := r.conn.Query(ctx, query, req.Year, req.Month)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []*pb.TransactionMonthlyAmountSuccess
	for rows.Next() {
		var item pb.TransactionMonthlyAmountSuccess
		if err := rows.Scan(&item.Year, &item.Month, &item.TotalSuccess, &item.TotalAmount); err != nil {
			return nil, err
		}
		result = append(result, &item)
	}
	return result, nil
}

func (r *clickhouseReaderRepository) GetYearlyAmountSuccess(ctx context.Context, req *pb.YearAmountTransactionRequest) ([]*pb.TransactionYearlyAmountSuccess, error) {
	query := `SELECT toYear(date) as y, sum(total_success), sum(total_amount) FROM transaction_stats WHERE toYear(date) = ? AND status = 'success' GROUP BY y`
	rows, err := r.conn.Query(ctx, query, req.Year)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []*pb.TransactionYearlyAmountSuccess
	for rows.Next() {
		var item pb.TransactionYearlyAmountSuccess
		if err := rows.Scan(&item.Year, &item.TotalSuccess, &item.TotalAmount); err != nil {
			return nil, err
		}
		result = append(result, &item)
	}
	return result, nil
}

func (r *clickhouseReaderRepository) GetMonthlyAmountFailed(ctx context.Context, req *pb.MonthAmountTransactionRequest) ([]*pb.TransactionMonthlyAmountFailed, error) {
	query := `SELECT toYear(date) as y, formatDateTime(date, '%M') as m, sum(total_failed), sum(total_amount) 
			  FROM transaction_stats WHERE toYear(date) = ? AND toMonth(date) = ? AND status = 'failed' GROUP BY y, m`
	rows, err := r.conn.Query(ctx, query, req.Year, req.Month)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []*pb.TransactionMonthlyAmountFailed
	for rows.Next() {
		var item pb.TransactionMonthlyAmountFailed
		if err := rows.Scan(&item.Year, &item.Month, &item.TotalFailed, &item.TotalAmount); err != nil {
			return nil, err
		}
		result = append(result, &item)
	}
	return result, nil
}

func (r *clickhouseReaderRepository) GetYearlyAmountFailed(ctx context.Context, req *pb.YearAmountTransactionRequest) ([]*pb.TransactionYearlyAmountFailed, error) {
	query := `SELECT toYear(date) as y, sum(total_failed), sum(total_amount) FROM transaction_stats WHERE toYear(date) = ? AND status = 'failed' GROUP BY y`
	rows, err := r.conn.Query(ctx, query, req.Year)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []*pb.TransactionYearlyAmountFailed
	for rows.Next() {
		var item pb.TransactionYearlyAmountFailed
		if err := rows.Scan(&item.Year, &item.TotalFailed, &item.TotalAmount); err != nil {
			return nil, err
		}
		result = append(result, &item)
	}
	return result, nil
}

func (r *clickhouseReaderRepository) GetMonthlyTransactionMethodSuccess(ctx context.Context, req *pb.MonthMethodTransactionRequest) ([]*pb.TransactionMonthlyMethod, error) {
	query := `SELECT formatDateTime(date, '%b') as m, payment_method, count(DISTINCT transaction_id), sum(total_amount) 
			  FROM transaction_stats WHERE toYear(date) = ? AND toMonth(date) = ? AND status = 'success' GROUP BY m, payment_method`
	rows, err := r.conn.Query(ctx, query, req.Year, req.Month)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []*pb.TransactionMonthlyMethod
	for rows.Next() {
		var item pb.TransactionMonthlyMethod
		if err := rows.Scan(&item.Month, &item.PaymentMethod, &item.TotalTransactions, &item.TotalAmount); err != nil {
			return nil, err
		}
		result = append(result, &item)
	}
	return result, nil
}

func (r *clickhouseReaderRepository) GetYearlyTransactionMethodSuccess(ctx context.Context, req *pb.YearMethodTransactionRequest) ([]*pb.TransactionYearlyMethod, error) {
	query := `SELECT toYear(date) as y, payment_method, count(DISTINCT transaction_id), sum(total_amount) 
			  FROM transaction_stats WHERE toYear(date) = ? AND status = 'success' GROUP BY y, payment_method`
	rows, err := r.conn.Query(ctx, query, req.Year)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []*pb.TransactionYearlyMethod
	for rows.Next() {
		var item pb.TransactionYearlyMethod
		if err := rows.Scan(&item.Year, &item.PaymentMethod, &item.TotalTransactions, &item.TotalAmount); err != nil {
			return nil, err
		}
		result = append(result, &item)
	}
	return result, nil
}

func (r *clickhouseReaderRepository) GetMonthlyTransactionMethodFailed(ctx context.Context, req *pb.MonthMethodTransactionRequest) ([]*pb.TransactionMonthlyMethod, error) {
	query := `SELECT formatDateTime(date, '%b') as m, payment_method, count(DISTINCT transaction_id), sum(total_amount) 
			  FROM transaction_stats WHERE toYear(date) = ? AND toMonth(date) = ? AND status = 'failed' GROUP BY m, payment_method`
	rows, err := r.conn.Query(ctx, query, req.Year, req.Month)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []*pb.TransactionMonthlyMethod
	for rows.Next() {
		var item pb.TransactionMonthlyMethod
		if err := rows.Scan(&item.Month, &item.PaymentMethod, &item.TotalTransactions, &item.TotalAmount); err != nil {
			return nil, err
		}
		result = append(result, &item)
	}
	return result, nil
}

func (r *clickhouseReaderRepository) GetYearlyTransactionMethodFailed(ctx context.Context, req *pb.YearMethodTransactionRequest) ([]*pb.TransactionYearlyMethod, error) {
	query := `SELECT toYear(date) as y, payment_method, count(DISTINCT transaction_id), sum(total_amount) 
			  FROM transaction_stats WHERE toYear(date) = ? AND status = 'failed' GROUP BY y, payment_method`
	rows, err := r.conn.Query(ctx, query, req.Year)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []*pb.TransactionYearlyMethod
	for rows.Next() {
		var item pb.TransactionYearlyMethod
		if err := rows.Scan(&item.Year, &item.PaymentMethod, &item.TotalTransactions, &item.TotalAmount); err != nil {
			return nil, err
		}
		result = append(result, &item)
	}
	return result, nil
}

// Transaction Stats By Merchant Implementation

func (r *clickhouseReaderRepository) GetMonthlyAmountSuccessByMerchant(ctx context.Context, req *pb.MonthAmountTransactionMerchantRequest) ([]*pb.TransactionMonthlyAmountSuccess, error) {
	query := `SELECT toYear(date) as y, formatDateTime(date, '%M') as m, sum(total_success), sum(total_amount) 
			  FROM transaction_stats WHERE toYear(date) = ? AND toMonth(date) = ? AND merchant_id = ? AND status = 'success' GROUP BY y, m`
	rows, err := r.conn.Query(ctx, query, req.Year, req.Month, req.MerchantId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []*pb.TransactionMonthlyAmountSuccess
	for rows.Next() {
		var item pb.TransactionMonthlyAmountSuccess
		if err := rows.Scan(&item.Year, &item.Month, &item.TotalSuccess, &item.TotalAmount); err != nil {
			return nil, err
		}
		result = append(result, &item)
	}
	return result, nil
}

func (r *clickhouseReaderRepository) GetYearlyAmountSuccessByMerchant(ctx context.Context, req *pb.YearAmountTransactionMerchantRequest) ([]*pb.TransactionYearlyAmountSuccess, error) {
	query := `SELECT toYear(date) as y, sum(total_success), sum(total_amount) FROM transaction_stats WHERE toYear(date) = ? AND merchant_id = ? AND status = 'success' GROUP BY y`
	rows, err := r.conn.Query(ctx, query, req.Year, req.MerchantId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []*pb.TransactionYearlyAmountSuccess
	for rows.Next() {
		var item pb.TransactionYearlyAmountSuccess
		if err := rows.Scan(&item.Year, &item.TotalSuccess, &item.TotalAmount); err != nil {
			return nil, err
		}
		result = append(result, &item)
	}
	return result, nil
}

func (r *clickhouseReaderRepository) GetMonthlyAmountFailedByMerchant(ctx context.Context, req *pb.MonthAmountTransactionMerchantRequest) ([]*pb.TransactionMonthlyAmountFailed, error) {
	query := `SELECT toYear(date) as y, formatDateTime(date, '%M') as m, sum(total_failed), sum(total_amount) 
			  FROM transaction_stats WHERE toYear(date) = ? AND toMonth(date) = ? AND merchant_id = ? AND status = 'failed' GROUP BY y, m`
	rows, err := r.conn.Query(ctx, query, req.Year, req.Month, req.MerchantId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []*pb.TransactionMonthlyAmountFailed
	for rows.Next() {
		var item pb.TransactionMonthlyAmountFailed
		if err := rows.Scan(&item.Year, &item.Month, &item.TotalFailed, &item.TotalAmount); err != nil {
			return nil, err
		}
		result = append(result, &item)
	}
	return result, nil
}

func (r *clickhouseReaderRepository) GetYearlyAmountFailedByMerchant(ctx context.Context, req *pb.YearAmountTransactionMerchantRequest) ([]*pb.TransactionYearlyAmountFailed, error) {
	query := `SELECT toYear(date) as y, sum(total_failed), sum(total_amount) FROM transaction_stats WHERE toYear(date) = ? AND merchant_id = ? AND status = 'failed' GROUP BY y`
	rows, err := r.conn.Query(ctx, query, req.Year, req.MerchantId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []*pb.TransactionYearlyAmountFailed
	for rows.Next() {
		var item pb.TransactionYearlyAmountFailed
		if err := rows.Scan(&item.Year, &item.TotalFailed, &item.TotalAmount); err != nil {
			return nil, err
		}
		result = append(result, &item)
	}
	return result, nil
}

func (r *clickhouseReaderRepository) GetMonthlyTransactionMethodByMerchantSuccess(ctx context.Context, req *pb.MonthMethodTransactionMerchantRequest) ([]*pb.TransactionMonthlyMethod, error) {
	query := `SELECT formatDateTime(date, '%b') as m, payment_method, count(DISTINCT transaction_id), sum(total_amount) 
			  FROM transaction_stats WHERE toYear(date) = ? AND toMonth(date) = ? AND merchant_id = ? AND status = 'success' GROUP BY m, payment_method`
	rows, err := r.conn.Query(ctx, query, req.Year, req.Month, req.MerchantId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []*pb.TransactionMonthlyMethod
	for rows.Next() {
		var item pb.TransactionMonthlyMethod
		if err := rows.Scan(&item.Month, &item.PaymentMethod, &item.TotalTransactions, &item.TotalAmount); err != nil {
			return nil, err
		}
		result = append(result, &item)
	}
	return result, nil
}

func (r *clickhouseReaderRepository) GetYearlyTransactionMethodByMerchantSuccess(ctx context.Context, req *pb.YearMethodTransactionMerchantRequest) ([]*pb.TransactionYearlyMethod, error) {
	query := `SELECT toYear(date) as y, payment_method, count(DISTINCT transaction_id), sum(total_amount) 
			  FROM transaction_stats WHERE toYear(date) = ? AND merchant_id = ? AND status = 'success' GROUP BY y, payment_method`
	rows, err := r.conn.Query(ctx, query, req.Year, req.MerchantId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []*pb.TransactionYearlyMethod
	for rows.Next() {
		var item pb.TransactionYearlyMethod
		if err := rows.Scan(&item.Year, &item.PaymentMethod, &item.TotalTransactions, &item.TotalAmount); err != nil {
			return nil, err
		}
		result = append(result, &item)
	}
	return result, nil
}

func (r *clickhouseReaderRepository) GetMonthlyTransactionMethodByMerchantFailed(ctx context.Context, req *pb.MonthMethodTransactionMerchantRequest) ([]*pb.TransactionMonthlyMethod, error) {
	query := `SELECT formatDateTime(date, '%b') as m, payment_method, count(DISTINCT transaction_id), sum(total_amount) 
			  FROM transaction_stats WHERE toYear(date) = ? AND toMonth(date) = ? AND merchant_id = ? AND status = 'failed' GROUP BY m, payment_method`
	rows, err := r.conn.Query(ctx, query, req.Year, req.Month, req.MerchantId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []*pb.TransactionMonthlyMethod
	for rows.Next() {
		var item pb.TransactionMonthlyMethod
		if err := rows.Scan(&item.Month, &item.PaymentMethod, &item.TotalTransactions, &item.TotalAmount); err != nil {
			return nil, err
		}
		result = append(result, &item)
	}
	return result, nil
}

func (r *clickhouseReaderRepository) GetYearlyTransactionMethodByMerchantFailed(ctx context.Context, req *pb.YearMethodTransactionMerchantRequest) ([]*pb.TransactionYearlyMethod, error) {
	query := `SELECT toYear(date) as y, payment_method, count(DISTINCT transaction_id), sum(total_amount) 
			  FROM transaction_stats WHERE toYear(date) = ? AND merchant_id = ? AND status = 'failed' GROUP BY y, payment_method`
	rows, err := r.conn.Query(ctx, query, req.Year, req.MerchantId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []*pb.TransactionYearlyMethod
	for rows.Next() {
		var item pb.TransactionYearlyMethod
		if err := rows.Scan(&item.Year, &item.PaymentMethod, &item.TotalTransactions, &item.TotalAmount); err != nil {
			return nil, err
		}
		result = append(result, &item)
	}
	return result, nil
}
