// Copyright © 2024 Jakub Kapusta <jakub-dev1@protonmail.com>
package find

import (
	"fmt"
	"os"
	"unicode"
)

// A printer should not check if f.ctx is cancelled.
// The sender will stop sending and close the channel if the ctx is cancelled.
// The printer should then process its buffer and exit on chanel close.
// Slower but prevents unexpected things to happen to our terminal.
func NewSafePrinter(f *finder, lineEnding string) chan<- *fileInfo {
	c := make(chan *fileInfo, 32)

	f.wg.Add(1)
	go func(c <-chan *fileInfo, f *finder) {
		defer f.wg.Done()
		for {
			select {
			case fi, ok := <-c:
				if !ok {
					return
				}
				rs := []rune(fi.path)

				// Most runes will be 2 bytes so allocate at least len(rs)*2.
				// *4 should prevent extra allocations in for all utf-8 strings.

				ret := make([]byte, 0, len(rs)*2)

				for _, r := range rs {
					if !unicode.IsControl(r) {
						ret = append(ret, byte(r))
					} else {
						// Control characters will be replaced with a string representation of their unicode code.
						// Example: U+0090 will be printed as the string literal U+0090, and not as the actual unicode code point.
						os.Stderr.WriteString(fmt.Sprintf("String contains unicode control characters: %q\n", fi.path))
						ret = append(ret, []byte(fmt.Sprintf("%U", r))...)
					}

				}

				// Append trailing / for directories
				if fi.d.IsDir() {
					ret = append(ret, '/')
				}
				// Append line ending.
				ret = append(ret, []byte(lineEnding)...)

				_, err := f.w.Write(ret)
				if err != nil {
					fmt.Println(err)
				}

			}
		}

	}(c, f)

	return c
}

// A printer should not check if f.ctx is cancelled.
// The sender will stop sending and close the channel if the ctx is cancelled.
// The printer should then process its buffer and exit on chanel close.
// Fast but unexpected things can happen to our terminal.
func NewUnsafePrinter(f *finder, lineEnding string) chan<- *fileInfo {
	c := make(chan *fileInfo, 32)

	f.wg.Add(1)
	go func(c <-chan *fileInfo, f *finder) {
		defer f.wg.Done()
		for {
			select {
			case fi, ok := <-c:
				if !ok {
					return
				}
				ret := []byte(fi.path)

				// Append trailing / for directories
				if fi.d.IsDir() {
					ret = append(ret, '/')
				}
				ret = append(ret, []byte(lineEnding)...)

				_, err := f.w.Write(ret)
				if err != nil {
					// TODO implement better logging.
					fmt.Println(err)
				}
			}
		}
	}(c, f)

	return c
}
