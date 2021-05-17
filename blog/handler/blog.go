package handler

import (
	"context"

	proto "github.com/m3o/sentiment-api/blog/proto"
	comments "github.com/m3o/sentiment-api/comments/proto"
	posts "github.com/m3o/sentiment-api/posts/proto"
	tags "github.com/m3o/sentiment-api/tags/proto"
)

type Blog struct {
	ps posts.PostsService
	cs comments.CommentsService
	ts tags.TagsService
}

func NewBlog(ps posts.PostsService,
	cs comments.CommentsService,
	ts tags.TagsService) *Blog {
	return &Blog{
		ps: ps,
		cs: cs,
		ts: ts,
	}
}

func (e *Blog) Latest(ctx context.Context, req *proto.LatestRequest, rsp *proto.LatestResponse) error {
	resp, err := e.ps.Query(ctx, &posts.QueryRequest{Limit: 1})
	if err != nil {
		return err
	}

	if len(resp.Posts) == 0 {
		return nil
	}

	rsp.Latest = resp.Posts[0]

	return nil
}

func (e *Blog) Posts(ctx context.Context, req *proto.PostsRequest, rsp *proto.PostsResponse) error {
	resp, err := e.ps.Query(ctx, &posts.QueryRequest{
		Limit:  req.Limit,
		Offset: req.Offset,
	})

	if err != nil {
		return err
	}

	rsp.Posts = resp.Posts
	return nil
}
