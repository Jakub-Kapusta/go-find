// Copyright © 2024 Jakub Kapusta <jakub-dev1@protonmail.com>
package find

import "os"

func Find(args []string, rootDir string, unsafePrint bool, print0 bool) {
	// fmt.Println("Root dir: ", rootDir)
	// if len(args) > 0 {
	// 	fmt.Println("args: ", args)
	// }

	f := newFinder(os.Stdout, rootDir, unsafePrint, print0)
	f.run()

}
