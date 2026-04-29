module github.com/MamangRust/monolith-ecommerce-grpc/service/stats-reader

go 1.25.1

require (
	github.com/ClickHouse/clickhouse-go/v2 v2.34.0
	github.com/MamangRust/monolith-ecommerce-pkg v0.0.0
	github.com/MamangRust/monolith-ecommerce-shared v0.0.0
	github.com/spf13/viper v1.21.0
	go.uber.org/zap v1.27.1
	google.golang.org/grpc v1.80.0
)

replace github.com/MamangRust/monolith-ecommerce-pkg => ../../pkg
replace github.com/MamangRust/monolith-ecommerce-shared => ../../shared
