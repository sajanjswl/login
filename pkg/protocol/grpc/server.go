package grpc

import (
	"context"
	"net"
	"os"
	"os/signal"

	v1 "github.com/sajanjswl/auth/pkg/api/v1"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func RunServer(ctx context.Context, v1API v1.UserServiceServer, port string) error {

	listen, err := net.Listen(os.Getenv("GRPC_NETWORK_TYPE"), ":"+port)
	if err != nil {
		return err
	}

	// register service
	server := grpc.NewServer()
	v1.RegisterUserServiceServer(server, v1API)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			// sig is a ^C, handle it
			log.Println("shutting down gRPC server...!!!")

			server.GracefulStop()

			<-ctx.Done()
		}
	}()

	// start gRPC server
	log.Println("starting gRPC server...")
	return server.Serve(listen)
}
