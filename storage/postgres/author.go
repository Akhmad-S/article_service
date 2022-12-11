package postgres

import (
	"github.com/uacademy/article/models"

	"errors"
)

func (stg Postgres) AddAuthor(id string, input models.CreateAuthorModel) error {
	_, err := stg.db.Exec(`INSERT INTO author (id, fullname) VALUES ($1, $2)`, id, input.Fullname)
	if err != nil {
		return err
	}
	return nil
}

func (stg Postgres) ReadAuthorById(id string) (models.Author, error) {
	var res models.Author
	err := stg.db.QueryRow(`SELECT id, fullname, created_at, updated_at, deleted_at FROM author WHERE id=$1`, id).Scan(
		&res.Id, &res.Fullname, &res.Created_at, &res.Updated_at, &res.Deleted_at,
	)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (stg Postgres) ReadListAuthor(offset, limit int, search string) (list []models.Author, err error) {
	rows, err := stg.db.Queryx(`SELECT
	id,
	fullname,
	created_at,
	updated_at,
	deleted_at
	FROM author WHERE deleted_at IS NULL AND (fullname ILIKE '%' || $1 || '%')
	LIMIT $2
	OFFSET $3
	`, search, limit, offset)

	if err != nil {
		return list, err
	}
	for rows.Next() {
		var a models.Author

		err := rows.Scan(
			&a.Id,
			&a.Fullname,
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

func (stg Postgres) UpdateAuthor(input models.UpdateAuthorModel) error {
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
