package author

import (
	"context"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/uacademy/blogpost/article_service/models"
	authorproto "github.com/uacademy/blogpost/article_service/proto-gen/blogpost"
	"github.com/uacademy/blogpost/article_service/storage"
)

// We define a AuthorService struct that implements the server interface.
type authorService struct {
	stg storage.StorageI
	authorproto.UnimplementedAuthorServiceServer
}

// NewAuthorService ...
func NewAuthorService(stg storage.StorageI) *authorService {
	return &authorService{
		stg: stg,
	}
}

// We implement the SayHello method of the AuthorService interface.
func (s *authorService) SayHello(ctx context.Context, in *authorproto.HelloRequest) (*authorproto.HelloReply, error) {
	return &authorproto.HelloReply{
		Message: "Hello, " + in.GetName()}, nil
}

func (s *authorService) CreateAuthor(ctx context.Context, req *authorproto.CreateAuthorRequest) (*authorproto.Author, error) {
	id := uuid.New()

	err := s.stg.AddAuthor(id.String(), models.CreateAuthorModel{
		Fullname: req.Fullname,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.AddAuthor: %s", err.Error())
	}

	author, err := s.stg.ReadAuthorById(id.String())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.ReadAuthorById: %s", err.Error())
	}

	var updatedAt string
	if author.Updated_at != nil {
		updatedAt = author.Updated_at.String()
	}
	
	return &authorproto.Author{
		Id: author.Id,
		Fullname: author.Fullname,
		CreatedAt: author.Created_at.String(),
		UpdatedAt: updatedAt,
	}, nil
}

func (s *authorService) UpdateAuthor(ctx context.Context, req *authorproto.UpdateAuthorRequest) (*authorproto.Author, error) {
	err := s.stg.UpdateAuthor(models.UpdateAuthorModel{
		Id: req.Id,
		Fullname: req.Fullname,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.UpdateAuthor: %s", err.Error())
	}

	author, err := s.stg.ReadAuthorById(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.ReadAuthorById: %s", err.Error())
	}

	var updatedAt string
	if author.Updated_at != nil {
		updatedAt = author.Updated_at.String()
	}
	
	return &authorproto.Author{
		Id: author.Id,
		Fullname: author.Fullname,
		CreatedAt: author.Created_at.String(),
		UpdatedAt: updatedAt,
	}, nil
}

func (s *authorService) DeleteAuthor(ctx context.Context, req *authorproto.DeleteAuthorRequest) (*authorproto.Author, error) {
	author, err := s.stg.ReadAuthorById(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.ReadAuthorById: %s", err.Error())
	}
	
	err = s.stg.DeleteAuthor(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.DeleteAuthor: %s", err.Error())
	}

	var updatedAt string
	if author.Updated_at != nil {
		updatedAt = author.Updated_at.String()
	}
	
	return &authorproto.Author{
		Id: author.Id,
		Fullname: author.Fullname,
		CreatedAt: author.Created_at.String(),
		UpdatedAt: updatedAt,
	}, nil
}

func (s *authorService) GetAuthorList(ctx context.Context, req *authorproto.GetAuthorListRequest) (*authorproto.GetAuthorListResponse, error) {
	res := &authorproto.GetAuthorListResponse{
		Authors: make([]*authorproto.Author, 0),
	}

	authorList, err := s.stg.ReadListAuthor(int(req.Offset), int(req.Limit), req.Search)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.ReadListAuthor: %s", err.Error())
	}

	for _, v := range authorList {
		var updatedAt string
		if v.Updated_at != nil {
			updatedAt = v.Updated_at.String()
		}

		res.Authors = append(res.Authors, &authorproto.Author{
			Id: v.Id,
			Fullname: v.Fullname,
			CreatedAt: v.Created_at.String(),
			UpdatedAt: updatedAt,
		})
	}
	return res, nil
}

func (s *authorService) GetAuthorById(ctx context.Context, req *authorproto.GetAuthorByIdRequest) (*authorproto.GetAuthorByIdResponse, error) {
	author, err := s.stg.ReadAuthorById(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.ReadAuthorById: %s", err.Error())
	}

	var updatedAt string
	if author.Updated_at != nil {
		updatedAt = author.Updated_at.String()
	}

	if author.Deleted_at != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.ReadAuthorById: author with id: %s not found", req.Id)
	}

	return &authorproto.GetAuthorByIdResponse{
		Id: author.Id,
		Fullname: author.Fullname,
		CreatedAt: author.Created_at.String(),
		UpdatedAt: updatedAt,
	}, nil
}
