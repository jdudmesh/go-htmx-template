package store

import (
	"context"
	"fmt"
	"gohtmx/internal/model"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/gommon/log"
	_ "github.com/mattn/go-sqlite3"
)

type sqliteDatabase struct {
	db            *sqlx.DB
	isDevelopment bool
}

func NewSqlite(databaseURL string, isDevelopment bool) (*sqliteDatabase, error) {
	db, err := sqlx.Connect("sqlite3", databaseURL)
	if err != nil {
		return nil, fmt.Errorf("connecting to database: %w", err)
	}

	return &sqliteDatabase{
		db, isDevelopment,
	}, nil
}

func (d *sqliteDatabase) Close() error {
	return d.db.Close()
}

func (d *sqliteDatabase) CurrentUser(ctx context.Context) (*model.User, error) {
	log.Error("TODO: you need to implement this")
	return nil, nil
}
