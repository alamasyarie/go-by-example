package main

import "fmt"

func Array() {
	var a [5]int // Declare array with 5 elements (by default each arary is zero)
	fmt.Println("emp:", a)

	a[4] = 100 // How to set element of array
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	fmt.Println("len:", len(a)) // HOw to check array length

	b := [5]int{1, 2, 3, 4, 5} // Shorthand to declare array

	fmt.Println("dcl:", b)

	var twoD [2][3]int // Multi-dimensional array
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
