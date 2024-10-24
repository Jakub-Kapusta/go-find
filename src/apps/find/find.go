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

func Test(args []string, rootDir string, unsafePrint bool, print0 bool) {
	// fmt.Println("Root dir: ", rootDir)
	// if len(args) > 0 {
	// 	fmt.Println("args: ", args)
	// }

	// 4 MiB buffer

	w := bufio.NewWriterSize(os.Stdout, 2048*2048)

	filepath.WalkDir(rootDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			os.Stderr.WriteString(err.Error() + newlineString)
			return nil
		}

		if unsafePrint {
			if print0 {
				output.UnsafePrint(w, path, nullString)
			} else {
				output.UnsafePrint(w, path, newlineString)
			}

		} else {
			if print0 {
				output.SafePrint(w, path, nullByte)
			} else {
				output.SafePrint(w, path, newlineByte)
			}
		}
		return nil
	})

	if err := w.Flush(); err != nil {
		fmt.Println(err)
	}
}
