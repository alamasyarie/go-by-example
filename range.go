package main

import "fmt"

func Range() {
	nums := []int{2, 3, 4}

	// Use range to iterate and find sum of slice/array
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	// Range on arrays and slices provides both the index and value
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	// Range on map iterates over key/value pairs.
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	// Range can also iterate over just the keys of a map.
	for k := range kvs {
		fmt.Println("key:", k)
	}

	// Range on strings iterates over Unicode code points
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
