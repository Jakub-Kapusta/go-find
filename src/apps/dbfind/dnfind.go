// Copyright © 2024 Jakub Kapusta <jakub-dev1@protonmail.com>
package dbfind

import (
	"context"
	"os"

	"github.com/Jakub-Kapusta/go-find/internal/printer"
	_ "github.com/mattn/go-sqlite3"
)

func DbFind(ctx context.Context, args []string, unsafePrint, print0, isSearchPath bool, searchPath string) {
	ph := printer.NewPrintHandler(os.Stdout, unsafePrint, print0)

	dbfh, err := newDbFindHandler(ctx, ph.GetPrintChan(), isSearchPath, searchPath)
	if err != nil {
		os.Stderr.WriteString(err.Error() + "\n")
		return
	}

	if err := dbfh.run(); err != nil {
		os.Stderr.WriteString(err.Error() + "\n")
		return
	}

}
