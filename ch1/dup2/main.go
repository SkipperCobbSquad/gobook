package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fileMap := make(map[string]map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, fileMap)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, fileMap)
			f.Close()
		}
		for fileName, lineMap := range fileMap {
			fileNamePrinted := false
			for line, n := range lineMap {
				if fileMap[fileName][line] > 1 {
					if !fileNamePrinted {
						fmt.Println(fileName)
						fileNamePrinted = true
					}
					fmt.Printf("%d\t%s\n", n, line)
				}
			}
		}
	}
}

func countLines(f *os.File, fileMap map[string]map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		if fileMap[f.Name()] == nil {
			fileMap[f.Name()] = make(map[string]int)
		}
		fileMap[f.Name()][input.Text()]++

	}
}
