package article

import (
	"context"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/uacademy/blogpost/article_service/models"
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

	err := s.stg.AddArticle(id.String(), models.CreateArticleModel{
		Content: models.Content{
			Title: req.Content.Title,
			Body:  req.Content.Body,
		},
		AuthorId: req.AuthorId,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.AddArticle: %s", err.Error())
	}

	article, err := s.stg.ReadArticleById(id.String())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.ReadArticleById: %s", err.Error())
	}

	var updatedAt string
	if article.Updated_at != nil {
		updatedAt = article.Updated_at.String()
	}

	return &articleproto.Article{
		Id: article.Id,
		Content: &articleproto.Content{
			Title: article.Content.Title,
			Body:  article.Content.Body,
		},
		AuthorId:  article.Author.Id,
		CreatedAt: article.Created_at.String(),
		UpdatedAt: updatedAt,
	}, nil
}

func (s *articleService) UpdateArticle(ctx context.Context, req *articleproto.UpdateArticleRequest) (*articleproto.Article, error) {
	err := s.stg.UpdateArticle(models.UpdateArticleModel{
		Id: req.Id,
		Content: models.Content{
			Title: req.Content.Title,
			Body:  req.Content.Body,
		},
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.UpdateArticle: %s", err.Error())
	}

	article, err := s.stg.ReadArticleById(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.ReadArticleById: %s", err.Error())
	}
	var updatedAt string
	if article.Updated_at != nil {
		updatedAt = article.Updated_at.String()
	}

	return &articleproto.Article{
		Id: article.Id,
		Content: &articleproto.Content{
			Title: article.Content.Title,
			Body:  article.Content.Body,
		},
		AuthorId:  article.Author.Id,
		CreatedAt: article.Created_at.String(),
		UpdatedAt: updatedAt,
	}, nil
}

func (s *articleService) DeleteArticle(ctx context.Context, req *articleproto.DeleteArticleRequest) (*articleproto.Article, error) {
	article, err := s.stg.ReadArticleById(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.ReadArticleById: %s", err.Error())
	}

	err = s.stg.DeleteArticle(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.DeleteArticle: %s", err.Error())
	}

	var updatedAt string
	if article.Updated_at != nil {
		updatedAt = article.Updated_at.String()
	}

	return &articleproto.Article{
		Id: article.Id,
		Content: &articleproto.Content{
			Title: article.Content.Title,
			Body:  article.Content.Body,
		},
		AuthorId:  article.Author.Id,
		CreatedAt: article.Created_at.String(),
		UpdatedAt: updatedAt,
	}, nil
}
func (s *articleService) GetArticleList(ctx context.Context, req *articleproto.GetArticleListRequest) (*articleproto.GetArticleListResponse, error) {
	res := &articleproto.GetArticleListResponse{
		Articles: make([]*articleproto.Article, 0),
	}

	articleList, err := s.stg.ReadListArticle(int(req.Offset), int(req.Limit), req.Search)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.ReadListArticle: %s", err.Error())
	}

	for _, v := range articleList {
		var updatedAt string
		if v.Updated_at != nil {
			updatedAt = v.Updated_at.String()
		}

		res.Articles = append(res.Articles, &articleproto.Article{
			Id: v.Id,
			Content: &articleproto.Content{
				Title: v.Content.Title,
				Body:  v.Content.Body,
			},
			AuthorId:  v.AuthorId,
			CreatedAt: v.Created_at.String(),
			UpdatedAt: updatedAt,
		})
	}
	return res, nil
}

func (s *articleService) GetArticleById(ctx context.Context, req *articleproto.GetArticleByIdRequest) (*articleproto.GetArticleByIdResponse, error) {
	article, err := s.stg.ReadArticleById(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.ReadArticleById: %s", err.Error())
	}

	var updatedAt string
	if article.Updated_at != nil {
		updatedAt = article.Updated_at.String()
	}

	var deletedAt string
	if article.Deleted_at != nil {
		deletedAt = article.Deleted_at.String()
	}

	var authorUpdatedAt string
	if article.Author.Updated_at != nil {
		authorUpdatedAt = article.Author.Updated_at.String()
	}

	return &articleproto.GetArticleByIdResponse{
		Id: article.Id,
		Content: &articleproto.Content{
			Title: article.Content.Title,
			Body:  article.Content.Body,
		},
		Author: &articleproto.GetArticleByIdResponse_Author{
			Id:        article.Author.Id,
			Fullname:  article.Author.Fullname,
			CreatedAt: article.Author.Created_at.String(),
			UpdatedAt: authorUpdatedAt,
		},
		CreatedAt: article.Created_at.String(),
		UpdatedAt: updatedAt,
		DeletedAt: deletedAt,
	}, nil
}
