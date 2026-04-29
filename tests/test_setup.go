package tests

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net"
	"os"
	"path/filepath"
	"time"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/clickhouse"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/modules/redis"
	"github.com/testcontainers/testcontainers-go/wait"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
	"google.golang.org/grpc"
	"github.com/jackc/pgx/v5/pgxpool"
	goredis "github.com/redis/go-redis/v9"

	"github.com/MamangRust/microservice-ecommerce-pkg/adapter"
	"github.com/MamangRust/microservice-ecommerce-pkg/auth"
	db "github.com/MamangRust/microservice-ecommerce-pkg/database/schema"
	"github.com/MamangRust/microservice-ecommerce-pkg/hash"
	"github.com/MamangRust/microservice-ecommerce-pkg/logger"
	"github.com/MamangRust/microservice-ecommerce-shared/cache"
	"github.com/MamangRust/microservice-ecommerce-shared/observability"
	"github.com/MamangRust/microservice-ecommerce-shared/pb"
	sdklog "go.opentelemetry.io/otel/sdk/log"

	role_handler "github.com/MamangRust/microservice-ecommerce-grpc-role/handler"
	role_repo "github.com/MamangRust/microservice-ecommerce-grpc-role/repository"
	role_service "github.com/MamangRust/microservice-ecommerce-grpc-role/service"
	user_handler "github.com/MamangRust/microservice-ecommerce-grpc-user/handler"
	user_repo "github.com/MamangRust/microservice-ecommerce-grpc-user/repository"
	user_service "github.com/MamangRust/microservice-ecommerce-grpc-user/service"
)

type TestSuite struct {
	PGContainer    *postgres.PostgresContainer
	RedisContainer *redis.RedisContainer
	CHContainer    *clickhouse.ClickHouseContainer
	DBURL          string
	RedisURL       string
	CHURL          string
	Ctx            context.Context

	UserAdapter     adapter.UserAdapter
	CardAdapter     adapter.CardAdapter
	MerchantAdapter adapter.MerchantAdapter
	SaldoAdapter    adapter.SaldoAdapter

	// Local gRPC Clients for Auth/Identity testing
	UserQueryClient   pb.UserQueryServiceClient
	UserCommandClient pb.UserCommandServiceClient
	RoleQueryClient   pb.RoleQueryServiceClient
	RoleCommandClient pb.RoleCommandServiceClient

	// Aliases for convenience
	UserClient *LocalUserClient
	RoleClient *LocalRoleClient

	// Shared resources
	Logger        logger.LoggerInterface
	CacheStore    *cache.CacheStore
	Observability observability.TraceLoggerObservability
	Hashing       hash.HashPassword
	TokenManager  auth.TokenManager
}

func SetupTestSuite() (*TestSuite, error) {
	ctx := context.Background()

	// Setup PostgreSQL
	pgContainer, err := postgres.Run(ctx,
		"postgres:16-alpine",
		postgres.WithDatabase("testdb"),
		postgres.WithUsername("testuser"),
		postgres.WithPassword("testpass"),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).
				WithStartupTimeout(60*time.Second)),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to start postgres container: %w", err)
	}

	dbURL, err := pgContainer.ConnectionString(ctx, "sslmode=disable")
	if err != nil {
		return nil, fmt.Errorf("failed to get postgres connection string: %w", err)
	}

	// Setup Redis
	redisContainer, err := redis.Run(ctx, "redis:7-alpine")
	if err != nil {
		return nil, fmt.Errorf("failed to start redis container: %w", err)
	}

	redisURL, err := redisContainer.ConnectionString(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get redis connection string: %w", err)
	}

	// Setup ClickHouse
	chContainer, err := clickhouse.Run(ctx,
		"clickhouse/clickhouse-server:24.3-alpine",
		clickhouse.WithDatabase("testdb"),
		clickhouse.WithUsername("testuser"),
		clickhouse.WithPassword("testpass"),
		testcontainers.WithWaitStrategy(
			wait.ForHTTP("/ping").WithPort("8123/tcp").WithStartupTimeout(30*time.Second)),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to start clickhouse container: %w", err)
	}

	chURL, err := chContainer.ConnectionString(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get clickhouse connection string: %w", err)
	}

	ts := &TestSuite{
		PGContainer:    pgContainer,
		RedisContainer: redisContainer,
		CHContainer:    chContainer,
		DBURL:          dbURL,
		RedisURL:       redisURL,
		CHURL:          chURL,
		Ctx:            ctx,
	}

	// Find project root
	root := ts.FindRootDir()
	if root == "" {
		return nil, fmt.Errorf("could not find root directory")
	}

	// Optional: Core migrations if needed
	// migrationsDir := filepath.Join(root, "pkg", "database", "migrations")
	// if err := ts.RunMigrations(migrationsDir); err != nil {
	// 	ts.Teardown()
	// 	return nil, fmt.Errorf("failed to run core migrations: %w", err)
	// }

	// Initialize Logging, Cache and Observability
	logger.ResetInstance()
	lp := sdklog.NewLoggerProvider()
	ts.Logger, _ = logger.NewLogger("test-integration", lp)

	redisOpts, err := goredis.ParseURL(redisURL)
	if err != nil {
		return nil, fmt.Errorf("failed to parse redis url: %w", err)
	}
	redisClient := goredis.NewClient(redisOpts)
	cacheMetrics, _ := observability.NewCacheMetrics("test-integration")
	ts.CacheStore = cache.NewCacheStore(redisClient, ts.Logger, cacheMetrics)
	ts.Observability, _ = observability.NewObservability("test-integration", ts.Logger)
	ts.Hashing = hash.NewHashingPassword()
	ts.TokenManager, _ = auth.NewManager("test-secret-key")

	// Initialize Dependencies for Adapters/LocalClients
	pool, err := pgxpool.New(ctx, dbURL)
	if err != nil {
		return nil, fmt.Errorf("failed to open pgxpool: %w", err)
	}
	queries := db.New(pool)

	// User & Role Local Setup
	roleRepos := role_repo.NewRepositories(queries)
	roleSvc := role_service.NewService(&role_service.Deps{
		Repository:    roleRepos,
		Logger:        ts.Logger,
		Cache:         nil, // Can be improved
		Observability: ts.Observability,
	})
	roleHandlerInstance := role_handler.NewHandler(&role_handler.Deps{
		Service: roleSvc,
		Logger:  ts.Logger,
	})
	rClient := &LocalRoleClient{Handler: roleHandlerInstance}
	ts.RoleQueryClient = rClient
	ts.RoleCommandClient = rClient
	ts.RoleClient = rClient

	userRepos := user_repo.NewRepositories(queries, rClient) // Use local role client
	userServiceInstance := user_service.NewService(&user_service.Deps{
		Repositories:  userRepos,
		Logger:        ts.Logger,
		Hash:          ts.Hashing,
		Cache:         nil, // Can be improved
		Observability: ts.Observability,
	})
	userHandlerInstance := user_handler.NewHandler(&user_handler.Deps{
		Service: userServiceInstance,
		Logger:  ts.Logger,
	})
	uClient := &LocalUserClient{Handler: userHandlerInstance}
	ts.UserQueryClient = uClient
	ts.UserCommandClient = uClient
	ts.UserClient = uClient

	// Adapters (using local queries/clients)
	userQueryRepo := user_repo.NewUserQueryRepository(queries)
	ts.UserAdapter = adapter.NewLocalUserAdapter(userQueryRepo)
	ts.CardAdapter = adapter.NewLocalCardAdapter(nil, nil)     // Placeholder
	ts.MerchantAdapter = adapter.NewLocalMerchantAdapter(nil) // Placeholder
	ts.SaldoAdapter = adapter.NewLocalSaldoAdapter(nil)       // Placeholder

	return ts, nil
}

func (ts *TestSuite) FindRootDir() string {
	cwd, _ := os.Getwd()
	root := cwd
	for {
		if _, err := os.Stat(filepath.Join(root, "justfile")); err == nil {
			return root
		}
		parent := filepath.Dir(root)
		if parent == root {
			return ""
		}
		root = parent
	}
}

func (ts *TestSuite) RunServiceMigrations(serviceName string) error {
	root := ts.FindRootDir()
	migrationsDir := filepath.Join(root, "service", serviceName, "database", "migrations")

	// Use service-specific migration table to avoid conflicts between decentralized services
	goose.SetTableName(fmt.Sprintf("goose_db_version_%s", serviceName))

	return ts.RunMigrations(migrationsDir)
}

func (ts *TestSuite) RunMigrations(migrationsDir string) error {
	db, err := sql.Open("pgx", ts.DBURL)
	if err != nil {
		return err
	}
	defer db.Close()

	goose.SetDialect("postgres")

	if err := goose.Up(db, migrationsDir); err != nil {
		return err
	}

	newVersion, _ := goose.GetDBVersion(db)
	fmt.Printf("DEBUG: Migrations finished for %s | New Version: %d\n", migrationsDir, newVersion)

	return nil
}

func (ts *TestSuite) DBPool() *pgxpool.Pool {
	pool, _ := pgxpool.New(ts.Ctx, ts.DBURL)
	return pool
}

func (ts *TestSuite) RedisClient() *goredis.Client {
	opts, _ := goredis.ParseURL(ts.RedisURL)
	return goredis.NewClient(opts)
}

func (ts *TestSuite) Teardown() {
	if ts.PGContainer != nil {
		if err := ts.PGContainer.Terminate(ts.Ctx); err != nil {
			log.Printf("failed to terminate postgres container: %v", err)
		}
	}
	if ts.RedisContainer != nil {
		if err := ts.RedisContainer.Terminate(ts.Ctx); err != nil {
			log.Printf("failed to terminate redis container: %v", err)
		}
	}
	if ts.CHContainer != nil {
		if err := ts.CHContainer.Terminate(ts.Ctx); err != nil {
			log.Printf("failed to terminate clickhouse container: %v", err)
		}
	}
}

type GRPCServerRunner interface {
	Serve(lis net.Listener) error
	Stop()
	GracefulStop()
}

func RunGRPCServer(server *grpc.Server) (string, error) {
	lis, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		return "", err
	}
	go func() {
		if err := server.Serve(lis); err != nil {
			log.Printf("grpc server error: %v", err)
		}
	}()
	return lis.Addr().String(), nil
}
