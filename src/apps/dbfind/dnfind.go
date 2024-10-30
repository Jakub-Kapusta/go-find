// Copyright © 2024 Jakub Kapusta <jakub-dev1@protonmail.com>
package dbfind

import (
	"context"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func DbFind(ctx context.Context, args []string, isSearchPath bool, searchPath string) {
	dbfh, err := newDbFindHandler(ctx, isSearchPath, searchPath)
	if err != nil {
		os.Stderr.WriteString(err.Error() + "\n")
		return
	}

	if err := dbfh.run(); err != nil {
		os.Stderr.WriteString(err.Error() + "\n")
		return
	}

}
