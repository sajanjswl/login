package rest

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/sajanjswl/auth/config"
	v1 "github.com/sajanjswl/auth/gen/go/auth/v1"
	restv1 "github.com/sajanjswl/auth/pkg/rest-service/v1"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

// RunServer runs HTTP/REST gateway
func RunServer(ctx context.Context, restServer restv1.RestServer, cfg *config.Config, logger *zap.Logger) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	rmux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	if err := v1.RegisterAuthServiceHandlerFromEndpoint(ctx, rmux, cfg.GRPCHost+":"+cfg.GRPCPort, opts); err != nil {

		logger.Error("failed to start gRPC-Rest gatewy", zap.Error(err))
		return err
	}

	restServer.Mux = http.NewServeMux()
	restServer.Mux.Handle("/", rmux)

	// calling handler
	restServer.InitialRoutes()

	srv := &http.Server{
		Addr:    ":" + cfg.RESTPort,
		Handler: restServer.Mux,
	}

	// graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			// sig is a ^C, handle it
			logger.Warn("shutting down gRPC-Rest Gateway....!!!")

		}

		_, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()

		_ = srv.Shutdown(ctx)
	}()

	logger.Info("stating gRPC-Rest Gateway....!!!")

	return srv.ListenAndServe()

}
