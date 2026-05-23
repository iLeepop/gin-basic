package db

import (
	"database/sql"
	"gin-basic/internal/cfg"
	"time"
)

const (
	defaultMaxOpenConns    = 25
	defaultMaxIdleConns    = 5
	defaultConnMaxLifetime = 5 * time.Minute
)

func configurePool(db *sql.DB, cp cfg.ConnectPool) {
	if cp.MaxOpenConns == 0 {
		db.SetMaxOpenConns(defaultMaxOpenConns)
	} else {
		db.SetMaxOpenConns(cp.MaxOpenConns)
	}
	if cp.MaxIdleConns == 0 {
		db.SetMaxIdleConns(defaultMaxIdleConns)
	} else {
		db.SetMaxIdleConns(cp.MaxIdleConns)
	}
	if cp.ConnMaxLifetime == 0 {
		db.SetConnMaxLifetime(defaultConnMaxLifetime)
	} else {
		db.SetConnMaxLifetime(time.Duration(cp.ConnMaxLifetime) * time.Second)
	}
}
