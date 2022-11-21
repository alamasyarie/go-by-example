package main

import "fmt"

// Function that takes two ints and returns their sum as an int.
func plus(a int, b int) int {
	return a + b
}

// When you have multiple consecutive parameters of the same type, just declare once
func plusPlus(a, b, c int) int {
	return a + b + c
}

// Functions with multiple return
func vals() (int, int) {
	return 3, 7
}

// Function with variadic args
func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0

	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}
