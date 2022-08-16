package cmd

import (
	"context"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/sajanjswl/auth/config"
	"github.com/sajanjswl/auth/pkg/protocol/grpc"
	"github.com/sajanjswl/auth/pkg/protocol/rest"
	restv1 "github.com/sajanjswl/auth/pkg/rest-service/v1"
	grpcv1 "github.com/sajanjswl/auth/pkg/service/v1"
	"go.uber.org/zap"
)

type Config struct {
	GRPCPort string
}

func RunServer(cfg *config.Config, logger *zap.Logger) {

	// initialising postgress db endpoints
	// postgres_endpoints := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
	// 	os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_NAME"), os.Getenv("DB_PASS"),
	// 	os.Getenv("DB_SSLMODE"))

	// fmt.Println("printing postgress endpoint", postgres_endpoints)

	// log.Println("connecting to db...")
	// db, err := gorm.Open(os.Getenv("DB_DIALECT"), postgres_endpoints)
	// defer db.Close()
	// if err != nil {
	// 	log.Warning("Failded to connect to DB", err)
	// }

	// if err = db.DB().Ping(); err != nil {
	// 	log.Fatal("failed to connect to the database")
	// }
	// log.Println("connected to db...!!!")

	ctx := context.Background()

	var db *gorm.DB
	//  passing DB connection to Grpc
	grpcAuthServerApi := grpcv1.NewAuthServiceServer(db, logger, cfg)

	// //passing DB connection to Rest
	restAuthServerApi := restv1.RestServer{Db: db}

	// // run HTTP gateway
	go func() {
		_ = rest.RunServer(ctx, restAuthServerApi, cfg, logger)
	}()

	// run grpc server
	grpc.RunServer(ctx, grpcAuthServerApi, cfg, logger)
}
