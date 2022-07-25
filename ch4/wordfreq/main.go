package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	wordCount := make(map[string]int)
	in := bufio.NewScanner(os.Stdin)

	in.Split(bufio.ScanWords)

	for in.Scan() {
		wordCount[in.Text()]++
	}

	fmt.Println("Word\tcount")
	for word, count := range wordCount {
		fmt.Printf("%q\t%d\n", word, count)
	}

}
