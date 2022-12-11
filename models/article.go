package models

import "time"

type Content struct {
	Title string `json:"title" binding:"required"`
	Body  string `json:"body" binding:"required"`
}

type Article struct {
	Id         string     `json:"id"`
	Content    Content    `json:"content"`
	AuthorId   string     `json:"author_id" binding:"required"`
	Created_at time.Time  `json:"created_at"`
	Updated_at *time.Time `json:"updated_at"`
	Deleted_at *time.Time `json:"-"`
}

type CreateArticleModel struct {
	Content  Content `json:"content"`
	AuthorId string  `json:"author_id" binding:"required"`
}

type PackedArticleModel struct {
	Id         string     `json:"id"`
	Content    Content    `json:"content"`
	Author     Author     `json:"author"`
	Created_at time.Time  `json:"created_at"`
	Updated_at *time.Time `json:"updated_at"`
	Deleted_at *time.Time `json:"deleted_at"`
}

type UpdateArticleModel struct {
	Id       string  `json:"id" binding:"required"`
	Content  Content `json:"content"`
}
