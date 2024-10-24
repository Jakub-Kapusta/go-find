// Copyright © 2024 Jakub Kapusta <jakub-dev1@protonmail.com>
package find

import (
	"bufio"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/Jakub-Kapusta/go-find/internal/output"
)

func Test(args []string, rootDir string, unsafePrint bool) {
	fmt.Println("Root dir: ", rootDir)
	if len(args) > 0 {
		fmt.Println("args: ", args)
	}

	// 4 MiB buffer
	w := bufio.NewWriterSize(os.Stdout, 2048*2048)

	filepath.WalkDir(rootDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			os.Stderr.WriteString(err.Error() + "\n")
			return nil
		}

		if unsafePrint {
			output.UnsafePrint(w, path)
		} else {
			output.SafePrint(w, path)
		}
		return nil
	})

	if err := w.Flush(); err != nil {
		fmt.Println(err)
	}
}
