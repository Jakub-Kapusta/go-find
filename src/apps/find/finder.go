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
	Path string
	D    fs.DirEntry
}

type Finder struct {
	ctx context.Context
	wg  sync.WaitGroup
	// Be sure to close after use.
	printChan    chan<- *FileInfo
	rootDir      string
	isSearchPath bool
	searchPath   string
}

func NewFinder(ctx context.Context, printChan chan<- *FileInfo, rootDir, searchPath string, isSearchPath bool) *Finder {
	var fi = &Finder{
		ctx:          ctx,
		printChan:    printChan,
		rootDir:      rootDir,
		isSearchPath: isSearchPath,
		searchPath:   searchPath,
	}

	return fi
}

func (f *Finder) Run() {
	f.wg.Add(1)
	go func(f *Finder) {
		defer f.wg.Done()
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
					Path: path,
					D:    d,
				}

				if f.isSearchPath {
					if strings.Contains(path, f.searchPath) {
						f.printChan <- fi
					}
				} else {
					f.printChan <- fi
				}
				return nil
			}
		})
		if err != nil {
			os.Stderr.WriteString("\n" + err.Error() + "\n")
		}

		close(f.printChan)
	}(f)
}

// Do not call directly.
func (f *Finder) Close() {
	f.wg.Wait()
}
