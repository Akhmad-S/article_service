package article

import (
	"context"
	articleproto "github.com/uacademy/blogpost/article_service/proto-gen/blogpost"
)
// We define a ArticleService struct that implements the server interface.
type ArticleService struct{
	articleproto.UnimplementedArticleServiceServer
}

// We implement the SayHello method of the ArticleService interface.
func (s *ArticleService) SayHello(ctx context.Context, in *articleproto.HelloRequest) (*articleproto.HelloReply, error) {
	return &articleproto.HelloReply{
		Message: "Hello, " + in.GetName()}, nil
}
