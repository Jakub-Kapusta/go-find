// Copyright © 2024 Jakub Kapusta <jakub-dev1@protonmail.com>
package dbupdate

import (
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
	"testing"
)

// goos: linux
// goarch: amd64
// pkg: github.com/Jakub-Kapusta/go-find/apps/dbupdate
// cpu: 13th Gen Intel(R) Core(TM) i7-13620H
// BenchmarkMd5sumShort-16          6591486               179.2 ns/op
// BenchmarkSha224Short-16         10657720               112.3 ns/op
// BenchmarkSha256Short-16         10908380               109.1 ns/op
// BenchmarkSha512Short-16          3952369               304.4 ns/op
// BenchmarkMd5sumLong-16           2524539               477.2 ns/op
// BenchmarkSha224Long-16           4867201               246.5 ns/op
// BenchmarkSha256Long-16           4876554               245.0 ns/op
// BenchmarkSha512Long-16           2057421               588.5 ns/op

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
