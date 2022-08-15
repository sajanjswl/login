package rest

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	v1 "github.com/sajanjswl/auth/pkg/api/v1"
	restv1 "github.com/sajanjswl/auth/pkg/rest-service/v1"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

// RunServer runs HTTP/REST gateway
func RunServer(ctx context.Context, restServer restv1.RestServer, grpcPort, grpcHost, httpPort string) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	rmux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	if err := v1.RegisterUserServiceHandlerFromEndpoint(ctx, rmux, grpcHost+":"+grpcPort, opts); err != nil {

		log.Fatalf("failed to start HTTP gateway: %v", err)
	}

	restServer.Mux = http.NewServeMux()
	restServer.Mux.Handle("/", rmux)

	// // calling handler
	restServer.InitialRoutes()

	srv := &http.Server{
		Addr:    ":" + httpPort,
		Handler: restServer.Mux,
	}

	// graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			// sig is a ^C, handle it

			log.Println("shutting down HTTP/REST gateway...!!!")
		}

		_, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()

		_ = srv.Shutdown(ctx)
	}()

	log.Println("starting HTTP/REST gateway...!!!")
	return srv.ListenAndServe()
}
