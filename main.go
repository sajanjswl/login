package main

import (
	"context"
	"flag"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sajanjswl/auth/config"
	"github.com/sajanjswl/auth/database"
	"github.com/sajanjswl/auth/models"
	"github.com/sajanjswl/auth/pkg/protocol/grpc"
	"github.com/sajanjswl/auth/pkg/protocol/rest"
	restv1 "github.com/sajanjswl/auth/pkg/rest-service/v1"
	grpcv1 "github.com/sajanjswl/auth/pkg/service/v1"
	"go.uber.org/zap"
)

const service string = "auth-service"

func main() {
	cfg := config.NewConfig()

	// grpc-server configs
	flag.StringVar(&cfg.GRPCHost, "grpc-host", "localhost", "grpc host")
	flag.StringVar(&cfg.GRPCPort, "grpc-port", "8000", "grpc port")
	flag.StringVar(&cfg.GRPCNetworkType, "grpc-network-type", "tcp", "grpc network type")

	// rest-server configs
	flag.StringVar(&cfg.RESTHost, "rest-host", "localhost", "rest host")
	flag.StringVar(&cfg.RESTPort, "rest-port", "9000", "rest port")
	flag.StringVar(&cfg.LoginUIEndPoint, "login-ui", "/v1/login/ui/", "Oauth login ui endpoint")
	flag.StringVar(&cfg.HTMLTemplate, "html-templates", "./templates", "html template path")
	flag.StringVar(&cfg.AbsoluteLogPath, "logpath", "logs.log", "application logs path")

	// database configs
	flag.StringVar(&cfg.DBHost, "db-host", "localhost", "database host")
	flag.StringVar(&cfg.DBPort, "db-port", "3305", "database port")
	flag.StringVar(&cfg.DBUserName, "db-user-name", "root", "db username")
	flag.StringVar(&cfg.DBPassword, "db-password", "example", "db password")
	flag.StringVar(&cfg.DBName, "db-name", "user", "database name")

	// google Oauth configs
	flag.StringVar(&cfg.GoogleLoginEnpoint, "google-login-enpoint", "/auth/google/login", "Oauth google login enpoint")
	flag.StringVar(&cfg.GoogleCallbackEnpoint, "google-callback-enpoint", "/auth/google/callback", "Outh google callback enpoint")
	flag.StringVar(&cfg.GoogleRedirectURl, "google-redirect-uri", "http://localhost:9000/auth/google/callback", "Oauth google redirect URI")
	flag.StringVar(&cfg.GoogleClientId, "google-client-id", "997183094499-fidl2a9ol2gutlmjbpaadrnndbura862.apps.googleusercontent.com", "Oauth google client id")
	flag.StringVar(&cfg.GoogleClientSecret, "google-client-secret", "-7rHfJbWJW-RS50McT6g8iEG", "Oauth google client secret")
	flag.Parse()

	logger, _ := zap.NewProduction()
	defer logger.Sync()

	db := database.InitDb(logger, cfg)
	db.AutoMigrate(&models.User{})

	ctx := context.Background()
	grpcAuthServerApi := grpcv1.NewAuthServiceServer(db, logger, cfg)
	// //passing DB connection to Rest
	restAuthServerApi := restv1.NewRestServer(db, cfg, logger)

	// // // run HTTP gateway
	go func() {
		_ = rest.RunServer(ctx, restAuthServerApi, cfg, logger)
	}()
	grpc.RunServer(ctx, grpcAuthServerApi, cfg, logger)

}
