package postgres

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Postgres struct{
	db *sqlx.DB
}

func InitDb(psqlConfig string) (*Postgres, error){
	var err error

	tempDb, err := sqlx.Connect("postgres", psqlConfig)
	if err != nil{
		return nil, err
	}

	return &Postgres{
		db: tempDb,
	}, nil
}
