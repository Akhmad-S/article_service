package postgres

import (
	"errors"

	"github.com/uacademy/article/models"
)

func (stg Postgres) AddArticle(id string, input models.CreateArticleModel) error {
	_, err := stg.ReadAuthorById(input.AuthorId)
	if err != nil {
		return err
	}

	_, err = stg.db.Exec(`INSERT INTO article (id, title, body, author_id) VALUES ($1, $2, $3, $4)`, id, input.Content.Title, input.Content.Body, input.AuthorId)
	if err != nil {
		return err
	}
	return nil
}

func (stg Postgres) ReadArticleById(id string) (models.PackedArticleModel, error) {
	var res models.PackedArticleModel

	err := stg.db.QueryRow(`SELECT
		ar.id, ar.title, ar.body, ar.created_at, ar.updated_at, ar.deleted_at,
		au.id, au.fullname, au.created_at, au.updated_at, au.deleted_at  
		FROM article ar JOIN author au ON ar.author_id = au.id WHERE ar.id = $1`, id).Scan(
			&res.Id, &res.Content.Title, &res.Content.Body, &res.Created_at, &res.Updated_at, &res.Deleted_at, &res.Author.Id, &res.Author.Fullname, &res.Author.Created_at, &res.Author.Updated_at, &res.Author.Deleted_at,
		)
	if err != nil{
		return res, err
	}
	
	return res, nil
}

func (stg Postgres) ReadListArticle(offset, limit int, search string) (list []models.Article, err error) {
	rows, err := stg.db.Queryx(`SELECT
	id,
	title,
	body,
	author_id,
	created_at,
	updated_at,
	deleted_at
	FROM article WHERE deleted_at IS NULL AND ((title ILIKE '%' || $1 || '%') OR (body ILIKE '%' || $1 || '%'))
	LIMIT $2
	OFFSET $3
	`, search, limit, offset)

	if err != nil{
		return list, err
	}
	for rows.Next() {
		var a models.Article

		err := rows.Scan(
			&a.Id,
			&a.Content.Title,
			&a.Content.Body,
			&a.AuthorId,
			&a.Created_at,
			&a.Updated_at,
			&a.Deleted_at,
		)
		if err != nil {
			return list, err
		}
		list = append(list, a)
	}

	return list, err
}

func (stg Postgres) UpdateArticle(input models.UpdateArticleModel) error {
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
