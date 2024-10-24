// Copyright © 2024 Jakub Kapusta <jakub-dev1@protonmail.com>
package output

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

// Slower but prevents unexpected things to happen to our terminal.
func SafePrint(w *bufio.Writer, s string, lineEnding byte) {
	rs := []rune(s)

	// Most runes will be 2 bytes so allocate at least len(rs)*2.
	// *4 should prevent extra allocations in for all utf-8 strings.

	ret := make([]byte, 0, len(rs)*2)

	for _, r := range rs {
		if !unicode.IsControl(r) {
			ret = append(ret, byte(r))
		} else {
			// Control characters will be replaced with a string representation of their unicode code.
			// Example: U+0090 will be printed as the string literal U+0090, and not as the actual unicode code point.
			os.Stderr.WriteString(fmt.Sprintf("String contains unicode control characters: %q\n", s))
			ret = append(ret, []byte(fmt.Sprintf("%U", r))...)
		}

	}

	// Add line ending.
	ret = append(ret, lineEnding)

	_, err := w.Write(ret)
	if err != nil {
		fmt.Println(err)
	}

}

// Fast but unexpected things can happen to our terminal.
func UnsafePrint(w *bufio.Writer, s string, lineEnding string) {
	_, err := w.WriteString(s + lineEnding)
	if err != nil {
		// TODO implement better logging.
		fmt.Println(err)
	}
}
