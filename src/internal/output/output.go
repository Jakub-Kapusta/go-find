// Copyright © 2024 Jakub Kapusta <jakub-dev1@protonmail.com>
package output

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func SafePrint(w *bufio.Writer, s string) {
	rs := []rune(s)

	// Most runes will be 2 bytes so allocate at least len(rs)*2.
	// *4 should prevent extra allocations in for all utf-8 strings.

	ret := make([]byte, 0, len(rs)*2)

	for _, r := range rs {
		if !unicode.IsControl(r) {
			ret = append(ret, byte(r))
		} else {
			os.Stderr.WriteString(fmt.Sprintf("String contains unicode control characters: %q\n", s))
			ret = append(ret, []byte(fmt.Sprintf("%U", r))...)
		}

	}

	// Add line ending.
	ret = append(ret, '\n')

	_, err := w.Write(ret)
	if err != nil {
		fmt.Println(err)
	}

}

func UnsafePrint(w *bufio.Writer, s string) {
	_, err := w.WriteString(s + "\n")
	if err != nil {
		fmt.Println(err)
	}
}
