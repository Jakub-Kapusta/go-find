// Copyright © 2024 Jakub Kapusta <jakub-dev1@protonmail.com>
package dbupdate

import (
	"context"
	"os"

	"github.com/Jakub-Kapusta/go-find/apps/find"
	_ "github.com/mattn/go-sqlite3"
)

func DbUpdate(ctx context.Context, args []string, rootDir string, isSearchPath bool, searchPath string) {
	dbh, err := newDbUpdateHandler(ctx, rootDir, isSearchPath, searchPath)
	if err != nil {
		os.Stderr.WriteString(err.Error() + "\n")
		return
	}

	fi := find.NewFinder(
		ctx,
		dbh.getChan(),
		&find.FinderOptions{
			RootDir:      rootDir,
			IsSearchPath: isSearchPath,
			SearchPath:   searchPath,
		},
	)
	fi.Run()

	if err := dbh.run(); err != nil {
		os.Stderr.WriteString(err.Error() + "\n")
		return
	}

	fi.Close()
}
