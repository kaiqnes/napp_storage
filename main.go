package main

import (
	"log"
	"storage/internal/frameworks/app"
)

func main() {
	server := app.Setup()
	if err := server.Run(); err != nil {
		log.Fatalf("failed to start server - err %v", err)
	}
}
