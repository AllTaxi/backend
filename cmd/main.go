package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"gitlab.com/golang-team-template/monolith/api"
	"gitlab.com/golang-team-template/monolith/pkg/middleware"
	"go.uber.org/zap"

	"github.com/blendle/zapdriver"
	"gitlab.com/golang-team-template/monolith/service"
	rds "gitlab.com/golang-team-template/monolith/storage/redis"

	"gitlab.com/golang-team-template/monolith/configs"
	devicerepo "gitlab.com/golang-team-template/monolith/storage/device"
	userrepo "gitlab.com/golang-team-template/monolith/storage/user"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/gomodule/redigo/redis"
	_ "github.com/joho/godotenv/autoload"                 // load .env file automatically
	_ "gitlab.com/golang-team-template/monolith/api/docs" //init swagger docs

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4/pgxpool"

	_ "github.com/lib/pq"
)

// @title Monolith Sample API
// @version 1.0
// @description This is a sample MONOLITHAPP server.
// @termsOfService

// @contact.name API Support
// @contact.url https://novalabtech.com/
// @contact.email kholboevdostonbek@gmail.com

// @license.name Apache Licence
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @BasePath /api
func main() {
	// =========================================================================
	// Config
	conf := configs.Load()

	if err := conf.Validate(); err != nil {
		panic(err)
	}

	logger, err := zapdriver.NewDevelopment() // with `development` set to `true`

	// =========================================================================
	// Postgres
	psqlString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		conf.PostgresHost,
		conf.PostgresPort,
		conf.PostgresUser,
		conf.PostgresPassword,
		conf.PostgresDatabase,
	)

	pgPool, err := pgxpool.Connect(context.Background(), psqlString)
	if err != nil {
		logger.Fatal("error in connection to postgres: ", zap.Error(err))
	}
	defer pgPool.Close()
	// =========================================================================
	// Migrations
	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", conf.PostgresUser, conf.PostgresPassword, conf.PostgresHost, conf.PostgresPort, conf.PostgresDatabase)
	m, err := migrate.New("file://pkg/migrations", dbURL)
	if err != nil {
		logger.Fatal("error in creating migrations: ", zap.Error(err))
	}
	fmt.Printf("")
	if err := m.Up(); err != nil {
		logger.Info("error updating migrations: ", zap.Error(err))
	}
	// =========================================================================
	// Redis
	rdsPool := redis.Pool{
		// Maximum number of idle connections in the pool.
		MaxIdle: 80,
		// max number of connections
		MaxActive: 12000,
		// Dial is an application supplied function for creating and
		// configuring a connection.
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", fmt.Sprintf("%s:%d", conf.RedisHost, conf.RedisPort))
			if err != nil {
				panic(err.Error())
			}
			return c, err
		},
	}
	// =========================================================================
	// Storage
	userRepo := userrepo.NewRepository(pgPool)
	deviceRepo := devicerepo.NewRepository(pgPool)
	redisRepo := rds.NewRedisRepo(&rdsPool)
	// =========================================================================
	// Sevice
	userService := service.NewUserService(userRepo)
	deviceService := service.NewDeviceService(deviceRepo)

	// =========================================================================
	// HTTP
	root := mux.NewRouter()
	// =========================================================================
	//Middleware
	root.Use(middleware.PanicRecovery)
	root.Use(middleware.Logging)
	casbinJWTRoleAuthorizer, err := middleware.NewCasbinJWTRoleAuthorizer(&conf)
	if err != nil {
		logger.Fatal("Could not initialize Cabin JWT Role Authorizer", zap.Error(err))
	}
	root.Use(casbinJWTRoleAuthorizer.Middleware)

	// API
	api.Init(root, userService, deviceService, redisRepo, logger)

	errChan := make(chan error, 1)
	osSignals := make(chan os.Signal, 1)
	signal.Notify(osSignals, os.Interrupt, syscall.SIGTERM)

	httpServer := http.Server{
		Addr:    conf.HTTPPort,
		Handler: root,
	}

	// http server
	go func() {
		errChan <- httpServer.ListenAndServe()
	}()

	// Blocking main and waiting for shutdown.
	select {
	case err := <-errChan:
		logger.Fatal("error: ", zap.Error(err))

	case <-osSignals:
		logger.Info("main : recieved os signal, shutting down")
		_ = httpServer.Shutdown(context.Background())
		return
	}
}
