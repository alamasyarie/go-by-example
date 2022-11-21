package main

import "fmt"

func Map() {
	// To create an empty map, use the built-in make function
	m := make(map[string]int)

	// Set  key/value
	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m)

	// Get value with key
	v1 := m["k1"]
	fmt.Println("v1: ", v1)

	// Built in length function
	fmt.Println("len:", len(m))

	// Built in delete function
	delete(m, "k2")
	fmt.Println("map:", m)

	// Shorthand declaration
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)
}
