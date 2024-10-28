// Copyright © 2024 Jakub Kapusta <jakub-dev1@protonmail.com>
package dbupdate

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func DbUpdate(ctx context.Context, args []string, rootDir string, isSearchPath bool, searchPath string) {
	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		os.Stderr.WriteString(err.Error() + "\n")
	}
	defer db.Close()

	var version string
	err = db.QueryRow("SELECT SQLITE_VERSION()").Scan(&version)

	if err != nil {
		os.Stderr.WriteString(err.Error() + "\n")
	}

	fmt.Println(version)
}
