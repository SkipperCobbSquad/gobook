package main

import "fmt"

func main() {
	//Opernads
	// var x uint8 = 1<<1 | 1<<5
	// var y uint8 = 1<<1 | 1<<2

	// fmt.Printf("%08b\n", x)
	// fmt.Printf("%08b\n", y)

	// fmt.Printf("1. %08b\n", x&y)
	// fmt.Printf("2. %08b\n", x|y)
	// fmt.Printf("3. %08b\n", x^y)
	// fmt.Printf("4. %08b\n", x&^y)

	// for i := uint(0); i < 8; i++ {
	// 	if x&(1<<i) != 0 {
	// 		fmt.Println(i)
	// 	}
	// }
	// fmt.Printf("%08b\n", x<<1)
	// fmt.Printf("%08b\n", x>>1)

	fmt.Println("------------------------------")
	//Printing
	o := 0666
	fmt.Printf("%d, %[1]o, %#[1]o\n", o)

	x := 0xdeadbeef
	fmt.Printf("%d, %#[1]x, %#[1]X\n", x)

}
