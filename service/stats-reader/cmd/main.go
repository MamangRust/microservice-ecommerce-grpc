package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/MamangRust/microservice-ecommerce-shared/pb"
	"github.com/MamangRust/microservice-ecommerce-pkg/clickhouse"
	"github.com/MamangRust/microservice-ecommerce-pkg/dotenv"
	"github.com/MamangRust/microservice-ecommerce-pkg/logger"
	"github.com/MamangRust/microservice-ecommerce-grpc/service/stats-reader/handler"
	"github.com/MamangRust/microservice-ecommerce-grpc/service/stats-reader/repository"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	if err := dotenv.Viper(); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load configuration: %v\n", err)
	}
	log, _ := logger.NewLogger("stats-reader", nil)
	
	chConn, err := clickhouse.NewClient(log)
	if err != nil {
		log.Fatal("Failed to connect to ClickHouse", zap.Error(err))
	}

	repo := repository.NewClickHouseReaderRepository(chConn)
	
	// Initialize Handlers
	categoryHandler := handler.NewCategoryStatsHandler(repo, log)
	categoryByMerchantHandler := handler.NewCategoryStatsByMerchantHandler(repo, log)
	categoryByIdHandler := handler.NewCategoryStatsByIdHandler(repo, log)
	orderHandler := handler.NewOrderStatsHandler(repo, log)
	transactionHandler := handler.NewTransactionStatsHandler(repo, log)
	transactionByMerchantHandler := handler.NewTransactionStatsByMerchantHandler(repo, log)

	server := grpc.NewServer()
	
	// Register Category Stats Services
	pb.RegisterCategoryStatsServiceServer(server, categoryHandler)
	pb.RegisterCategoryStatsByMerchantServiceServer(server, categoryByMerchantHandler)
	pb.RegisterCategoryStatsByIdServiceServer(server, categoryByIdHandler)
	
	// Register Order Stats Service
	pb.RegisterOrderStatsServiceServer(server, orderHandler)
	
	// Register Transaction Stats Services
	pb.RegisterTransactionStatsServiceServer(server, transactionHandler)
	pb.RegisterTransactionStatsByMerchantServiceServer(server, transactionByMerchantHandler)

	reflection.Register(server)

	port := ":50071"
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("Failed to listen", zap.Error(err))
	}

	log.Info("Stats Reader starting", zap.String("port", port))

	go func() {
		if err := server.Serve(lis); err != nil {
			log.Fatal("Failed to serve", zap.Error(err))
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Info("Shutting down Stats Reader...")
	server.GracefulStop()
}
