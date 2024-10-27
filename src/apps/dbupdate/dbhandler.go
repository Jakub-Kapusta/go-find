// Copyright © 2024 Jakub Kapusta <jakub-dev1@protonmail.com>
package dbupdate

import (
	"bufio"
	"context"
	"os"
	"sync"
)

type dbHandler struct {
	ctx          context.Context
	wg           sync.WaitGroup
	w            *bufio.Writer
	rootDir      string
	isSearchPath bool
	searchPath   string
}

func newDbHandler(ctx context.Context, f *os.File, rootDir string, isSearchPath bool, searchPath string) *dbHandler {
	return &dbHandler{
		ctx: ctx,
		// 4 MiB buffer
		w:            bufio.NewWriterSize(f, 2048*2048),
		rootDir:      rootDir,
		isSearchPath: isSearchPath,
		searchPath:   searchPath,
	}
}

func (dbh *dbHandler) run() error {
	return dbh.close()
}

func (dbh *dbHandler) close() error {
	dbh.wg.Wait()

	return nil
}
