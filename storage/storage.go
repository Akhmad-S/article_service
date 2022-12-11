package storage

import "github.com/uacademy/article/models"

type StorageI interface {
	AddArticle(id string, input models.CreateArticleModel) error
	ReadArticleById(id string) (models.PackedArticleModel, error)
	ReadListArticle(offset, limit int, search string) (list []models.Article, err error)
	UpdateArticle(input models.UpdateArticleModel) error
	DeleteArticle(id string) error
	AddAuthor(id string, input models.CreateAuthorModel) error
	ReadAuthorById(id string) (models.Author, error)
	ReadListAuthor(offset, limit int, search string) (list []models.Author, err error)
	UpdateAuthor(input models.UpdateAuthorModel) error
	DeleteAuthor(id string) error
}
