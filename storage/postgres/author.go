package postgres

import (
	"github.com/uacademy/blogpost/article_service/proto-gen/blogpost"

	"errors"
)

func (stg Postgres) AddAuthor(id string, input *blogpost.CreateAuthorRequest) error {
	_, err := stg.db.Exec(`INSERT INTO author (id, fullname) VALUES ($1, $2)`, id, input.Fullname)
	if err != nil {
		return err
	}
	return nil
}

func (stg Postgres) ReadAuthorById(id string) (*blogpost.GetAuthorByIdResponse, error) {
	res := &blogpost.GetAuthorByIdResponse{}
	var updatedAt *string

	err := stg.db.QueryRow(`SELECT id, fullname, created_at, updated_at FROM author WHERE id=$1 AND deleted_at IS NULL`, id).Scan(
		&res.Id, &res.Fullname, &res.CreatedAt, &updatedAt,
	)

	if err != nil{
		return nil, errors.New("author not found")
	}

	if updatedAt != nil{
		res.UpdatedAt = *updatedAt
	}

	return res, nil
}

func (stg Postgres) ReadListAuthor(offset, limit int, search string) (*blogpost.GetAuthorListResponse, error) {
	resp := &blogpost.GetAuthorListResponse{
		Authors: make([]*blogpost.Author, 0),
	}
	
	rows, err := stg.db.Queryx(`SELECT
	id,
	fullname,
	created_at,
	updated_at
	FROM author WHERE deleted_at IS NULL AND (fullname ILIKE '%' || $1 || '%')
	LIMIT $2
	OFFSET $3
	`, search, limit, offset)

	if err != nil {
		return resp, err
	}
	for rows.Next() {
		a := &blogpost.Author{}
		var updatedAt *string

		err := rows.Scan(
			&a.Id,
			&a.Fullname,
			&a.CreatedAt,
			&updatedAt,
		)
		if err != nil {
			return resp, err
		}

		if updatedAt != nil{
			a.UpdatedAt = *updatedAt
		}

		resp.Authors = append(resp.Authors, a)
	}

	return resp, err
}

func (stg Postgres) UpdateAuthor(input *blogpost.UpdateAuthorRequest) error {
	res, err := stg.db.NamedExec("UPDATE author  SET fullname=:fn, updated_at=now() WHERE deleted_at IS NULL AND id=:id", map[string]interface{}{
		"id": input.Id,
		"fn": input.Fullname,
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

	return errors.New("author not found")
}

func (stg Postgres) DeleteAuthor(id string) error {
	res, err := stg.db.Exec("UPDATE author  SET deleted_at=now() WHERE id=$1 AND deleted_at IS NULL", id)
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
	return errors.New("author not found")
}
