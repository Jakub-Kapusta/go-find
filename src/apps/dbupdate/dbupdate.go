// Copyright © 2024 Jakub Kapusta <jakub-dev1@protonmail.com>
package dbupdate

import (
	"context"
	"fmt"
	"os"
)

func DbUpdate(ctx context.Context, args []string, rootDir string, isSearchPath bool, searchPath string) {

	dbh := newDbHandler(ctx, os.Stdout, rootDir, isSearchPath, searchPath)
	if err := dbh.run(); err != nil {
		fmt.Println(err)
	}

}
