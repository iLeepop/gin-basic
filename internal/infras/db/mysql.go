package db

import (
	"database/sql"
	"fmt"

	"gin-basic/internal/cfg"
	"gin-basic/internal/ports/idb"

	_ "github.com/go-sql-driver/mysql"
)

type mysql struct {
	db *sql.DB
}

func NewMySQL(c cfg.Mysql, cp cfg.ConnectPool) (idb.IMySQL, error) {
	if c.Host == "" {
		return nil, nil
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		c.Username, c.Password, c.Host, c.Port, c.Database)

	sqlDB, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("open mysql: %w", err)
	}

	configurePool(sqlDB, cp)

	if err := sqlDB.Ping(); err != nil {
		_ = sqlDB.Close()
		return nil, fmt.Errorf("ping mysql: %w", err)
	}

	return &mysql{db: sqlDB}, nil
}

func (m *mysql) DB() *sql.DB {
	return m.db
}

func (m *mysql) Ping() error {
	return m.db.Ping()
}

func (m *mysql) Close() error {
	return m.db.Close()
}
