package main

import (
	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
	"github.com/m3o/sentiment-api/comments/handler"
)

func main() {
	// Create the service
	srv := service.New(
		service.Name("comments"),
	)

	// Register Handler
	srv.Handle(handler.NewComments())

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
