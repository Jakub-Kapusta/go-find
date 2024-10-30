// Copyright © 2024 Jakub Kapusta <jakub-dev1@protonmail.com>
package types

import "io/fs"

type Runner interface {
	Run() error
}

type Closer interface {
	Close(rollback bool) error
}

type FileInfoSinker interface {
	GetChan() chan<- *FileInfo
}

type FileInfo struct {
	Path string
	D    fs.DirEntry
}
