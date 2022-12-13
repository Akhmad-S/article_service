package storage

import (
	"github.com/uacademy/blogpost/article_service/proto-gen/blogpost"
)

type StorageI interface {
	AddArticle(id string, input *blogpost.CreateArticleRequest) error
	ReadArticleById(id string) (*blogpost.GetArticleByIdResponse, error)
	ReadListArticle(offset, limit int, search string) (resp *blogpost.GetArticleListResponse, err error)
	UpdateArticle(input *blogpost.UpdateArticleRequest) error
	DeleteArticle(id string) error

	AddAuthor(id string, input *blogpost.CreateAuthorRequest) error
	ReadAuthorById(id string) (*blogpost.GetAuthorByIdResponse, error)
	ReadListAuthor(offset, limit int, search string) (resp *blogpost.GetAuthorListResponse, err error)
	UpdateAuthor(input *blogpost.UpdateAuthorRequest) error
	DeleteAuthor(id string) error
}
