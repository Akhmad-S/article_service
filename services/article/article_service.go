package article

import (
	"context"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	articleproto "github.com/uacademy/blogpost/article_service/proto-gen/blogpost"
	"github.com/uacademy/blogpost/article_service/storage"
)

// We define a articleService struct that implements the server interface.

type articleService struct {
	stg storage.StorageI
	articleproto.UnimplementedArticleServiceServer
}

// NewArticleService ...
func NewArticleService(stg storage.StorageI) *articleService {
	return &articleService{
		stg: stg,
	}
}

// We implement the SayHello method of the ArticleService interface.
func (s *articleService) SayHello(ctx context.Context, in *articleproto.HelloRequest) (*articleproto.HelloReply, error) {
	return &articleproto.HelloReply{
		Message: "Hello, " + in.GetName()}, nil
}

func (s *articleService) CreateArticle(ctx context.Context, req *articleproto.CreateArticleRequest) (*articleproto.Article, error) {
	id := uuid.New()

	err := s.stg.AddArticle(id.String(), req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.AddArticle: %s", err.Error())
	}

	article, err := s.stg.ReadArticleById(id.String())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.ReadArticleById: %s", err.Error())
	}

	return &articleproto.Article{
		Id: article.Id,
		Content: article.Content,
		AuthorId:  article.Author.Id,
		CreatedAt: article.CreatedAt,
		UpdatedAt: article.UpdatedAt,
	}, nil
}

func (s *articleService) UpdateArticle(ctx context.Context, req *articleproto.UpdateArticleRequest) (*articleproto.Article, error) {
	err := s.stg.UpdateArticle(req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.UpdateArticle: %s", err.Error())
	}

	article, err := s.stg.ReadArticleById(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.ReadArticleById: %s", err.Error())
	}

	return &articleproto.Article{
		Id: article.Id,
		Content: article.Content,
		AuthorId:  article.Author.Id,
		CreatedAt: article.CreatedAt,
		UpdatedAt: article.UpdatedAt,
	}, nil
}

func (s *articleService) DeleteArticle(ctx context.Context, req *articleproto.DeleteArticleRequest) (*articleproto.Article, error) {
	article, err := s.stg.ReadArticleById(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.ReadArticleById: %s", err.Error())
	}

	err = s.stg.DeleteArticle(article.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.DeleteArticle: %s", err.Error())
	}

	return &articleproto.Article{
		Id: article.Id,
		Content: article.Content,
		AuthorId:  article.Author.Id,
		CreatedAt: article.CreatedAt,
		UpdatedAt: article.UpdatedAt,
	}, nil
}

func (s *articleService) GetArticleList(ctx context.Context, req *articleproto.GetArticleListRequest) (*articleproto.GetArticleListResponse, error) {
	res, err := s.stg.ReadListArticle(int(req.Offset), int(req.Limit), req.Search)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.ReadListArticle: %s", err.Error())
	}

	return res, nil
}

func (s *articleService) GetArticleById(ctx context.Context, req *articleproto.GetArticleByIdRequest) (*articleproto.GetArticleByIdResponse, error) {
	article, err := s.stg.ReadArticleById(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.ReadArticleById: %s", err.Error())
	}

	return article, nil
}
