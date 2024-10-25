// Copyright © 2024 Jakub Kapusta <jakub-dev1@protonmail.com>
package find

import (
	"bufio"
	"context"
	"io/fs"
	"os"
	"path/filepath"
	"sync"
)

type finder struct {
	ctx         context.Context
	wg          sync.WaitGroup
	w           *bufio.Writer
	printer     chan<- string // Close when done.
	rootDir     string
	unsafePrint bool
	print0      bool
}

func newFinder(ctx context.Context, f *os.File, rootDir string, unsafePrint bool, print0 bool) *finder {
	return &finder{
		ctx: ctx,
		// 4 MiB buffer
		w:           bufio.NewWriterSize(f, 2048*2048),
		rootDir:     rootDir,
		unsafePrint: unsafePrint,
		print0:      print0,
	}
}

func (f *finder) run() error {
	var lineEnding string
	if f.print0 {
		lineEnding = nullString
	} else {
		lineEnding = newlineString
	}

	if f.unsafePrint {
		f.printer = NewUnsafePrinter(f, lineEnding)
	} else {
		f.printer = NewSafePrinter(f, lineEnding)
	}

	err := filepath.WalkDir(f.rootDir, func(path string, d fs.DirEntry, err error) error {
		select {
		case <-f.ctx.Done():
			return f.ctx.Err()
		default:
			if err != nil {
				os.Stderr.WriteString(err.Error() + newlineString)
				return nil
			}

			f.printer <- path
			return nil
		}
	})
	if err != nil {
		os.Stderr.WriteString("\n" + err.Error() + "\n")
	}
	return f.close()
}

func (f *finder) close() error {
	close(f.printer)
	f.wg.Wait()

	if err := f.w.Flush(); err != nil {
		return err
	}
	return nil
}
