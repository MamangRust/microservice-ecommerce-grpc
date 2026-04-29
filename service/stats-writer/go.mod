module github.com/MamangRust/microservice-ecommerce-grpc/service/stats-writer

go 1.25.1

require (
	github.com/ClickHouse/clickhouse-go/v2 v2.34.0
	github.com/IBM/sarama v1.46.3
	github.com/MamangRust/microservice-ecommerce-pkg v0.0.0
	github.com/MamangRust/microservice-ecommerce-shared v0.0.0
	github.com/spf13/viper v1.21.0
	go.uber.org/zap v1.27.1
)

replace github.com/MamangRust/microservice-ecommerce-pkg => ../../pkg
replace github.com/MamangRust/microservice-ecommerce-shared => ../../shared
