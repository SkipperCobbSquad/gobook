package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	defer printStack()
	f(3)

}

func printStack() {
	var buff [4096]byte
	n := runtime.Stack(buff[:], false)
	os.Stdout.Write(buff[:n])
}

func f(x int) {
	fmt.Printf("f(%d)\n", x+0/x)
	defer fmt.Printf("defer %d\n", x)
	f(x - 1)
}
