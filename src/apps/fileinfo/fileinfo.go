// Copyright © 2024 Jakub Kapusta <jakub-dev1@protonmail.com>
package fileinfo

import "io/fs"

type FileInfo struct {
	Path string
	D    fs.DirEntry
}
