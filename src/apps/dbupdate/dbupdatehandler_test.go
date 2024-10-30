// Copyright © 2024 Jakub Kapusta <jakub-dev1@protonmail.com>
package dbupdate

import (
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
	"testing"
)

var shortTeststring = []byte("Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua")
var longTeststring = []byte("Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.")

func BenchmarkMd5sumShort(b *testing.B) {
	for range b.N {
		_ = md5.Sum([]byte(shortTeststring))
	}
}

func BenchmarkSha224Short(b *testing.B) {
	for range b.N {
		_ = sha256.Sum224([]byte(shortTeststring))
	}
}

func BenchmarkSha256Short(b *testing.B) {
	for range b.N {
		_ = sha256.Sum256([]byte(shortTeststring))
	}
}

func BenchmarkSha512Short(b *testing.B) {
	for range b.N {
		_ = sha512.Sum512([]byte(shortTeststring))
	}
}

func BenchmarkMd5sumLong(b *testing.B) {
	for range b.N {
		_ = md5.Sum([]byte(longTeststring))
	}
}

func BenchmarkSha224Long(b *testing.B) {
	for range b.N {
		_ = sha256.Sum224([]byte(longTeststring))
	}
}
func BenchmarkSha256Long(b *testing.B) {
	for range b.N {
		_ = sha256.Sum256([]byte(longTeststring))
	}
}

func BenchmarkSha512Long(b *testing.B) {
	for range b.N {
		_ = sha512.Sum512([]byte(longTeststring))
	}
}
