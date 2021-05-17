package main

import (
	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
	"github.com/m3o/sentiment-api/posts/handler"
	tags "github.com/m3o/sentiment-api/tags/proto"
)

func main() {
	// Create the service
	srv := service.New(
		service.Name("posts"),
	)

	// Register Handler
	srv.Handle(handler.NewPosts(
		tags.NewTagsService("tags", srv.Client()),
	))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
