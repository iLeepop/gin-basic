package idb

import "database/sql"

type IPostgreSQL interface {
	DB() *sql.DB
	Ping() error
	Close() error
}
