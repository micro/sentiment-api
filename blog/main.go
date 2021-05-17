package main

import (
	"github.com/m3o/sentiment-api/blog/handler"
	proto "github.com/m3o/sentiment-api/blog/proto"
	comments "github.com/m3o/sentiment-api/comments/proto"
	posts "github.com/m3o/sentiment-api/posts/proto"
	tags "github.com/m3o/sentiment-api/tags/proto"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("blog"),
	)

	// Register handler
	proto.RegisterBlogHandler(srv.Server(), handler.NewBlog(
		posts.NewPostsService("posts", srv.Client()),
		comments.NewCommentsService("comments", srv.Client()),
		tags.NewTagsService("tags", srv.Client()),
	))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
