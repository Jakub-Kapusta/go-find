// Copyright © 2024 Jakub Kapusta <jakub-dev1@protonmail.com>
package find

import (
	"context"
	"os"

	"github.com/Jakub-Kapusta/go-find/internal/printer"
)

func Find(ctx context.Context, fio *FinderOptions) {
	ph := printer.NewPrintHandler(os.Stdout, fio.UnsafePrint, fio.Print0)

	f := NewFinder(ctx, ph.GetPrintChan(), fio)
	f.Run()
	// First close finder, then printer
	f.Close()
	ph.Close()
}
