package main

import (
	"github.com/dezhab-service/pkg/cmd"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("No .env file found")
	}
}
func main() {
	if err := cmd.RunServer(); err != nil {
		log.Fatal(err)

	}

}
