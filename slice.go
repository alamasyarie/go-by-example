package main

import "fmt"

func Slice() {
	s := make([]string, 3) // To create an empty slice with non-zero length, use the builtin make
	fmt.Println("emp:", s)

	// To set and get slice element
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	fmt.Println("len:", len(s)) // Built-in function to get length slice

	// To append element in slice
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	// You can copy slice
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	// Slices support a “slice” operator with the syntax slice[low:high]
	l := s[2:5]
	fmt.Println("sl1:", l)

	l = s[:5]
	fmt.Println("sl2:", l)

	l = s[2:]
	fmt.Println("sl3:", l)

	// Shorthand declaration
	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	// Multi dimensional slice
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
