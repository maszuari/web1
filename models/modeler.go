package modeler

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Datasource interface {
	FindOrgByName(shortname string) (string, error)
}

type DB struct {
	*sqlx.DB
}

func (db *DB) FindOrgByName(shortname string) (string, error) {
	var name string
	err := db.Get(&name, "SELECT fullname FROM organizations WHERE shortname = $1", shortname)
	if err != nil {
		return "none", err
	}
	return name, nil
}

func Connect() (*DB, error) {
	db, err := sqlx.Connect("postgres", "host=172.17.0.1 user=web1 dbname=testwebdb sslmode=disable password=secret")
	if err != nil {
		return nil, err
	}
	return &DB{db}, nil
}
