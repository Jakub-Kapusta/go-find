// Copyright © 2024 Jakub Kapusta <jakub-dev1@protonmail.com>
package find

import (
	"context"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

type FileInfo struct {
	path string
	d    fs.DirEntry
}

type Finder struct {
	ctx          context.Context
	wg           sync.WaitGroup
	p            *printHandler
	rootDir      string
	isSearchPath bool
	searchPath   string
}

func NewFinder(ctx context.Context, f *os.File, rootDir string, isSearchPath bool, searchPath string, unsafePrint, print0 bool) *Finder {
	var fi = &Finder{
		ctx:          ctx,
		rootDir:      rootDir,
		isSearchPath: isSearchPath,
		searchPath:   searchPath,
	}

	fi.p = NewPrintHandler(f, unsafePrint, print0)

	return fi
}

func (f *Finder) run() error {

	err := filepath.WalkDir(f.rootDir, func(path string, d fs.DirEntry, err error) error {
		select {
		case <-f.ctx.Done():
			return f.ctx.Err()
		default:
			if err != nil {
				os.Stderr.WriteString(err.Error() + newlineString)
				return nil
			}
			fi := &FileInfo{
				path: path,
				d:    d,
			}

			if f.isSearchPath {
				if strings.Contains(path, f.searchPath) {
					f.p.printer <- fi
				}
			} else {
				f.p.printer <- fi
			}
			return nil
		}
	})
	if err != nil {
		os.Stderr.WriteString("\n" + err.Error() + "\n")
	}
	return f.close()
}

func (f *Finder) close() error {
	f.p.close()
	f.wg.Wait()

	if err := f.p.w.Flush(); err != nil {
		return err
	}
	return nil
}
