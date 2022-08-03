package data

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

func GetDbConnection() (*sql.DB, error) {
	db, err := sql.Open("mysql", os.Getenv("DSN"))
	if err != nil {
		return nil, errors.Wrap(err, "(GetDbConnection) sql.Open")
	}
	return db, nil
}
