// Copyright © 2024 Jakub Kapusta <jakub-dev1@protonmail.com>
package dbfind

import (
	"context"
	"database/sql"
	"sync"
)

type dbFindHandler struct {
	ctx          context.Context
	wg           sync.WaitGroup
	db           *sql.DB
	c            chan string
	isSearchPath bool
	searchPath   string
}

func newDbFindHandler(ctx context.Context, isSearchPath bool, searchPath string) (*dbFindHandler, error) {
	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		return nil, err
	}

	return &dbFindHandler{
		ctx:          ctx,
		db:           db,
		c:            make(chan string, 32),
		isSearchPath: isSearchPath,
		searchPath:   searchPath,
	}, nil
}

func (dbh *dbFindHandler) run() error {
	return nil
}

func (dbfh *dbFindHandler) getChan() <-chan string {
	return dbfh.c
}

// Do not call directly.
func (dbfh *dbFindHandler) close(rollback bool) error {
	dbfh.wg.Wait()

	if err := dbfh.db.Close(); err != nil {
		return err
	}

	return nil
}
