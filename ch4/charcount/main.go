package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	counts := make(map[rune]int)
	var utfllen [utf8.UTFMax + 1]int
	invalid := 0

	in := bufio.NewReader(os.Stdin)

	for {
		r, n, err := in.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "Charcount: %v\n", err)
			os.Exit(1)
		}

		if r == unicode.ReplacementChar && n == 1 {
			invalid++
		}
		counts[r]++
		utfllen[n]++
	}
	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Print("\nlen\tcount\n")
	for i, n := range utfllen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}

	if invalid > 0 {
		fmt.Printf("Invalid: %d\n", invalid)
	}
}
