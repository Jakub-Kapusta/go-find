// Copyright © 2024 Jakub Kapusta <jakub-dev1@protonmail.com>
package find

import (
	"context"
	"fmt"
	"os"
)

func Find(ctx context.Context, args []string, rootDir string, isSearchPath bool, searchPath string, unsafePrint, print0 bool) {
	// fmt.Println("Root dir: ", rootDir)
	// if len(args) > 0 {
	// 	fmt.Println("args: ", args)
	// }

	f := newFinder(ctx, os.Stdout, rootDir, isSearchPath, searchPath, unsafePrint, print0)
	if err := f.run(); err != nil {
		fmt.Println(err)
	}

}
