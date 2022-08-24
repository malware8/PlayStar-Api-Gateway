package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type Repository struct {
	DataBase *sqlx.DB
}

func (r *Repository) NewConnection(host string, port string, user string, password string, dbname string) {
	credentials := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host,
		port,
		user,
		password,
		dbname)
	if conn, err := sqlx.Connect("postgres", credentials); err != nil {
		logrus.Fatalf("Database connection error. Error: %s", err.Error())
	} else {
		r.DataBase = conn
	}
}
