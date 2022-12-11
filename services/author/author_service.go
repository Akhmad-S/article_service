package author

import (
	"context"
	authorproto "github.com/uacademy/blogpost/article_service/proto-gen/blogpost"
)
// We define a AuthorService struct that implements the server interface.
type AuthorService struct{
	authorproto.UnimplementedAuthorServiceServer
}

// We implement the SayHello method of the AuthorService interface.
func (s *AuthorService) SayHello(ctx context.Context, in *authorproto.HelloRequest) (*authorproto.HelloReply, error) {
	return &authorproto.HelloReply{
		Message: "Hello, " + in.GetName()}, nil
}
