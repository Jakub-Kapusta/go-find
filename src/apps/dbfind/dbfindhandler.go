// Copyright © 2024 Jakub Kapusta <jakub-dev1@protonmail.com>
package dbfind

import (
	"context"
	"database/sql"
	"sync"

	"github.com/Jakub-Kapusta/go-find/apps/types"
)

type dbFindHandler struct {
	ctx          context.Context
	wg           sync.WaitGroup
	db           *sql.DB
	sink         chan<- *types.FileInfo
	isSearchPath bool
	searchPath   string
}

func newDbFindHandler(ctx context.Context, sink chan<- *types.FileInfo, isSearchPath bool, searchPath string) (*dbFindHandler, error) {
	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		return nil, err
	}

	return &dbFindHandler{
		ctx:          ctx,
		db:           db,
		sink:         sink,
		isSearchPath: isSearchPath,
		searchPath:   searchPath,
	}, nil
}

func (dbh *dbFindHandler) run() error {
	return nil
}

// Do not call directly.
func (dbfh *dbFindHandler) close(rollback bool) error {
	dbfh.wg.Wait()

	if err := dbfh.db.Close(); err != nil {
		return err
	}

	return nil
}
