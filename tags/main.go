package main

import (
	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
	"github.com/m3o/sentiment-api/tags/handler"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("tags"),
	)

	// Register Handler
	srv.Handle(handler.NewTags())

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
