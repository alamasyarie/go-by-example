package main

import "fmt"

func Variables() {
	a := "ini adalah tipe data string"
	fmt.Println(a)

	var b, c int = 1, 2 // You can declare multiple variables
	fmt.Println(b, c)

	d := true // Go will infer the type of initialized variables
	fmt.Println(d)

	var e int // Variables declared without assign value are zero-valued.
	fmt.Println(e)

	f := "apple" // Short assingment
	fmt.Println(f)
}
