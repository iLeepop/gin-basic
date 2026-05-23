package db

import (
	"database/sql"
	"fmt"

	"gin-basic/internal/cfg"
	"gin-basic/internal/ports/idb"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type postgresql struct {
	db *sql.DB
}

func NewPostgreSQL(c cfg.PostgreSQL, cp cfg.ConnectPool) (idb.IPostgreSQL, error) {
	if c.Host == "" {
		return nil, nil
	}

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		c.Username, c.Password, c.Host, c.Port, c.Database)

	sqlDB, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, fmt.Errorf("open postgresql: %w", err)
	}

	configurePool(sqlDB, cp)

	if err := sqlDB.Ping(); err != nil {
		_ = sqlDB.Close()
		return nil, fmt.Errorf("ping postgresql: %w", err)
	}

	return &postgresql{db: sqlDB}, nil
}

func (p *postgresql) DB() *sql.DB {
	return p.db
}

func (p *postgresql) Ping() error {
	return p.db.Ping()
}

func (p *postgresql) Close() error {
	return p.db.Close()
}
