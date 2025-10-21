package main

import (
	"log"

	"flakeforge/fin/config"
	"flakeforge/fin/internal/app"

	"github.com/joho/godotenv"
)

func main() {
  if err := godotenv.Load(); err !=  nil{
    log.Printf("Warning: .env file not found: %v", err)
  }

  config, err := config.New()
  if err != nil {
    log.Printf("Warning: .env file not found: %v", err)
  }

  app.Run(config)
}
