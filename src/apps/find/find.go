// Copyright © 2024 Jakub Kapusta <jakub-dev1@protonmail.com>
package find

import (
	"context"
	"os"
)

func Find(ctx context.Context, args []string, fio *FinderOptions, unsafePrint, print0 bool) {
	// fmt.Println("Root dir: ", rootDir)
	// if len(args) > 0 {
	// 	fmt.Println("args: ", args)
	// }

	ph := NewPrintHandler(os.Stdout, unsafePrint, print0)

	f := NewFinder(ctx, ph.getPrintChan(), fio)
	f.Run()
	// First close finder, then printer
	f.Close()
	ph.close()
}
