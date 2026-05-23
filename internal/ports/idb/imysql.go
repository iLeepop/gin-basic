package idb

import "database/sql"

type IMySQL interface {
	DB() *sql.DB
	Ping() error
	Close() error
}
