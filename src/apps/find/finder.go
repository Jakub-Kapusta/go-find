// Copyright © 2024 Jakub Kapusta <jakub-dev1@protonmail.com>
package find

import (
	"bufio"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sync"

	"github.com/Jakub-Kapusta/go-find/internal/output"
)

type finder struct {
	wg          sync.WaitGroup
	w           *bufio.Writer
	rootDir     string
	unsafePrint bool
	print0      bool
}

func newFinder(f *os.File, rootDir string, unsafePrint bool, print0 bool) *finder {
	return &finder{
		// 4 MiB buffer
		w:           bufio.NewWriterSize(f, 2048*2048),
		rootDir:     rootDir,
		unsafePrint: unsafePrint,
		print0:      print0,
	}
}

func (f *finder) run() {
	filepath.WalkDir(f.rootDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			os.Stderr.WriteString(err.Error() + newlineString)
			return nil
		}

		if f.unsafePrint {
			if f.print0 {
				output.UnsafePrint(f.w, path, nullString)
			} else {
				output.UnsafePrint(f.w, path, newlineString)
			}

		} else {
			if f.print0 {
				output.SafePrint(f.w, path, nullByte)
			} else {
				output.SafePrint(f.w, path, newlineByte)
			}
		}
		return nil
	})

	if err := f.w.Flush(); err != nil {
		fmt.Println(err)
	}
}
