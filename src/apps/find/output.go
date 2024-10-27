// Copyright © 2024 Jakub Kapusta <jakub-dev1@protonmail.com>
package find

import (
	"bufio"
	"fmt"
	"os"
	"sync"
	"unicode"
)

type printHandler struct {
	wg          sync.WaitGroup
	w           *bufio.Writer
	printer     chan<- *FileInfo // Close when done.
	unsafePrint bool
	print0      bool
	lineEnding  string
}

func NewPrintHandler(f *os.File, unsafePrint, print0 bool) *printHandler {

	ph := &printHandler{
		// 4 MiB buffer
		w:           bufio.NewWriterSize(f, 2048*2048),
		unsafePrint: unsafePrint,
		print0:      print0,
	}

	if print0 {
		ph.lineEnding = nullString
	} else {
		ph.lineEnding = newlineString
	}

	ph.run()

	return ph
}

func (ph *printHandler) safePrinter(c <-chan *FileInfo) {
	defer ph.wg.Done()
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
			ret = append(ret, []byte(ph.lineEnding)...)

			_, err := ph.w.Write(ret)
			if err != nil {
				fmt.Println(err)
			}

		}
	}
}
func (ph *printHandler) unsafePrinter(c <-chan *FileInfo) {
	defer ph.wg.Done()
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
			ret = append(ret, []byte(ph.lineEnding)...)

			_, err := ph.w.Write(ret)
			if err != nil {
				// TODO implement better logging.
				fmt.Println(err)
			}
		}
	}
}

func (ph *printHandler) run() {
	c := make(chan *FileInfo, 32)

	if ph.unsafePrint {
		// Slower but prevents unexpected things to happen to our terminal.
		ph.wg.Add(1)
		go ph.safePrinter(c)
	} else {
		// Fast but unexpected things can happen to our terminal.
		ph.wg.Add(1)
		go ph.unsafePrinter(c)
	}

	ph.printer = c
}

func (ph *printHandler) close() {
	close(ph.printer)
	ph.wg.Wait()
}
