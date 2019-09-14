
package db

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	// "os"
)

type Postgres struct {
	sqlx *sqlx.DB
}

func NewPostgres() *Postgres {
	return &Postgres{}
}

func (postgres *Postgres) Open() *DB {

	db, err := sqlx.Open("postgres", "host=127.0.0.1 port=5432 user=hoge password=hogemaru dbname=todo_graphql_api_development sslmode=disable")
	if err != nil {
		panic(err.Error())
	}

	return &DB{SqlxDB: db}
}
