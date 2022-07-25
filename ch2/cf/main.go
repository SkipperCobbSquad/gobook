package main

import (
	"bufio"
	"fmt"
	"gobook/ch2/tempconv"
	"os"
	"strconv"
)

func main() {
	if len(os.Args[1:]) == 0 {
		input := bufio.NewScanner(os.Stdin)
		input.Scan()
		t, err := strconv.ParseFloat(input.Text(), 64)

		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		f := tempconv.Fahrenheit(t)
		c := tempconv.Celsius(t)
		fmt.Printf("%s == %s, %s == %s\n", f, tempconv.FtoC(f), c, tempconv.CtoF(c))

	} else {
		for _, arg := range os.Args[1:] {
			t, err := strconv.ParseFloat(arg, 64)
			if err != nil {
				fmt.Fprintf(os.Stderr, "cf: %v\n", err)
				os.Exit(1)
			}
			f := tempconv.Fahrenheit(t)
			c := tempconv.Celsius(t)
			fmt.Printf("%s == %s, %s == %s\n", f, tempconv.FtoC(f), c, tempconv.CtoF(c))
		}
	}

}
