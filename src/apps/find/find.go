// Copyright © 2024 Jakub Kapusta <jakub-dev1@protonmail.com>
package find

import "fmt"

func Test(args []string, rootDir string) {
	fmt.Println("Root dir: ", rootDir)
	if len(args) > 0 {
		fmt.Println("args: ", args)
	}
}
