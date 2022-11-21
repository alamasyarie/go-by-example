package main

import (
	"fmt"
)

const s string = "constant" // Const declare a constant value (cannot change)

func Constant() {
	fmt.Println(s)

	const n = 500000000

	const d = 3e20 / n // Constant expressions perform arithmetic with arbitrary precision.
	fmt.Println(d)

	fmt.Println(int64(d)) // Numeric constant has no type until you give it
}
