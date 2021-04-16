package cmd

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/dezhab-service/pkg/protocol/grpc"
	"github.com/dezhab-service/pkg/protocol/rest"
	restv1 "github.com/dezhab-service/pkg/rest-service/v1"
	v1 "github.com/dezhab-service/pkg/service/v1"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	log "github.com/sirupsen/logrus"
)

type Config struct {
	GRPCPort string
}

func RunServer() error {

	// initialising postgress db endpoints
	postgres_endpoints := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_NAME"), os.Getenv("DB_PASS"),
		os.Getenv("DB_SSLMODE"))

	fmt.Println("printing postgress endpoint", postgres_endpoints)

	log.Println("connecting to db...")
	db, err := gorm.Open(os.Getenv("DB_DIALECT"), postgres_endpoints)
	defer db.Close()
	if err != nil {
		log.Warning("Failded to connect to DB", err)
	}

	if err = db.DB().Ping(); err != nil {
		log.Fatal("failed to connect to the database")
	}
	log.Println("connected to db...!!!")
	ctx := context.Background()
	var cfg Config
	flag.StringVar(&cfg.GRPCPort, "grpc-port", "", "gRPC port to bind")
	flag.Parse()
	if len(cfg.GRPCPort) == 0 {
		return fmt.Errorf("invalid TCP port for gRPC server: '%s'", cfg.GRPCPort)
	}

	//  passing DB connection to Grpc
	v1API := v1.NewUserServiceServer(db)

	//passing DB connection to Rest
	restServer := restv1.RestServer{Db: db}

	// run HTTP gateway
	go func() {
		_ = rest.RunServer(ctx, restServer, cfg.GRPCPort, os.Getenv("GRPC_HOST"), os.Getenv("REST_PORT"))
	}()

	// run grpc server
	return grpc.RunServer(ctx, v1API, cfg.GRPCPort)
}
