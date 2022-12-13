package postgres

import (
	"errors"
	"time"

	"github.com/uacademy/blogpost/article_service/proto-gen/blogpost"
)

func (stg Postgres) AddArticle(id string, input *blogpost.CreateArticleRequest) error {
	_, err := stg.ReadAuthorById(input.AuthorId)
	if err != nil {
		return err
	}

	if input.Content == nil{
		input.Content = &blogpost.Content{}
	}

	_, err = stg.db.Exec(`INSERT INTO article (id, title, body, author_id) VALUES ($1, $2, $3, $4)`, id, input.Content.Title, input.Content.Body, input.AuthorId)
	if err != nil {
		return err
	}
	return nil
}

func (stg Postgres) ReadArticleById(id string) (*blogpost.GetArticleByIdResponse, error) {
	res := &blogpost.GetArticleByIdResponse{
		Content: &blogpost.Content{},
		Author: &blogpost.GetArticleByIdResponse_Author{},
	}
	var deletedAt *time.Time
	var updatedAt, authorUpdatedAt *string

	err := stg.db.QueryRow(`SELECT
		ar.id, ar.title, ar.body, ar.created_at, ar.updated_at, ar.deleted_at,
		au.id, au.fullname, au.created_at, au.updated_at  
		FROM article ar JOIN author au ON ar.author_id = au.id WHERE ar.id = $1`, id).Scan(
		&res.Id, &res.Content.Title, &res.Content.Body, &res.CreatedAt, &updatedAt, &deletedAt, &res.Author.Id, &res.Author.Fullname, &res.Author.CreatedAt, &authorUpdatedAt,
	)
	if err != nil {
		return res, err
	}

	if updatedAt != nil{
		res.UpdatedAt = *updatedAt
	}

	if authorUpdatedAt != nil{
		res.Author.UpdatedAt = *authorUpdatedAt
	}

	if deletedAt != nil{
		return res, errors.New("article not found")
	}

	return res, nil
}

func (stg Postgres) ReadListArticle(offset, limit int, search string) (*blogpost.GetArticleListResponse, error) {
	resp := &blogpost.GetArticleListResponse{
		Articles: make([]*blogpost.Article, 0),
	}

	rows, err := stg.db.Queryx(`SELECT
	id,
	title,
	body,
	author_id,
	created_at,
	updated_at
	FROM article WHERE deleted_at IS NULL AND ((title ILIKE '%' || $1 || '%') OR (body ILIKE '%' || $1 || '%'))
	LIMIT $2
	OFFSET $3
	`, search, limit, offset)

	if err != nil {
		return resp, err
	}
	for rows.Next() {
		a := &blogpost.Article{
			Content: &blogpost.Content{},
		}

		var updatedAt *string

		err := rows.Scan(
			&a.Id,
			&a.Content.Title,
			&a.Content.Body,
			&a.AuthorId,
			&a.CreatedAt,
			&updatedAt,
		)
		if err != nil {
			return resp, err
		}

		if updatedAt != nil{
			a.UpdatedAt = *updatedAt
		}
		resp.Articles = append(resp.Articles, a)
	}

	return resp, err
}

func (stg Postgres) UpdateArticle(input *blogpost.UpdateArticleRequest) error {
	if input.Content == nil{
		input.Content = &blogpost.Content{}
	}

	res, err := stg.db.NamedExec("UPDATE article  SET title=:t, body=:b, updated_at=now() WHERE deleted_at IS NULL AND id=:id", map[string]interface{}{
		"id": input.Id,
		"t":  input.Content.Title,
		"b":  input.Content.Body,
	})
	if err != nil {
		return err
	}

	n, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if n > 0 {
		return nil
	}

	return errors.New("article not found")
}

func (stg Postgres) DeleteArticle(id string) error {
	res, err := stg.db.Exec("UPDATE article  SET deleted_at=now() WHERE id=$1 AND deleted_at IS NULL", id)
	if err != nil {
		return err
	}

	n, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if n > 0 {
		return nil
	}
	return errors.New("article not found")
}
