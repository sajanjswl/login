package main

import (
	// "github.com/sajanjswl/auth/pkg/cmd"
	// "github.com/joho/godotenv"
	// log "github.com/sirupsen/logrus"
	// "github.com/sajanjswl/auth/config"

	"flag"
	"fmt"
)


func main() {
	
	fmt.Println("Hello")

	cfg :=config.NewConfig()

	flag.StringVar(&cfg.AbsoluteLogPath,"","SDfsdfds" )
	// flag.StringVar(&cfg.GRPCPort, "grpc-port", "", "gRPC port to bind")
	flag.Parse()

	fmt.Println(cfg.AbsoluteLogPath)
	// _ = cmd.RunServer()

}
