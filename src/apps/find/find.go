// Copyright © 2024 Jakub Kapusta <jakub-dev1@protonmail.com>
package find

import (
	"context"
	"os"

	"github.com/Jakub-Kapusta/go-find/internal/printer"
)

func Find(ctx context.Context, fio *FinderOptions, unsafePrint, print0 bool) {
	ph := printer.NewPrintHandler(os.Stdout, unsafePrint, print0)

	f := NewFinder(ctx, ph.GetPrintChan(), fio)
	f.Run()
	// First close finder, then printer
	f.Close()
	ph.Close()
}
