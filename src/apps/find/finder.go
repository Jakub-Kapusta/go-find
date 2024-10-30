// Copyright © 2024 Jakub Kapusta <jakub-dev1@protonmail.com>
package find

import (
	"context"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/Jakub-Kapusta/go-find/apps/fileinfo"
	"github.com/Jakub-Kapusta/go-find/internal/printer"
)

type FinderOptions struct {
	RootDir      string
	IsSearchPath bool
	SearchPath   string
}

type Finder struct {
	ctx context.Context
	wg  sync.WaitGroup
	// Be sure to close after use.
	printChan chan<- *fileinfo.FileInfo
	fio       *FinderOptions
}

func NewFinder(ctx context.Context, printChan chan<- *fileinfo.FileInfo, fio *FinderOptions) *Finder {
	var fi = &Finder{
		ctx:       ctx,
		printChan: printChan,
		fio:       fio,
	}

	return fi
}

func (f *Finder) Run() {
	f.wg.Add(1)
	go func(f *Finder) {
		defer f.wg.Done()
		err := filepath.WalkDir(f.fio.RootDir, func(path string, d fs.DirEntry, err error) error {
			select {
			case <-f.ctx.Done():
				return f.ctx.Err()
			default:
				if err != nil {
					os.Stderr.WriteString(err.Error() + printer.NewlineString)
					return nil
				}
				fi := &fileinfo.FileInfo{
					Path: path,
					D:    d,
				}

				if f.fio.IsSearchPath {
					if strings.Contains(path, f.fio.SearchPath) {
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
