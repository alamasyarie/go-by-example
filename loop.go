package main

import "fmt"

func Loop() {
	i := 1 // Basic type loop with a single condition
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	for j := 7; j <= 9; j++ { // A classic for loop
		fmt.Println(j)
	}

	for { // Without condition loop cannot be stop until you break or return something
		fmt.Println("loop")
		break
	}

	for n := 0; n <= 5; n++ { // You can also continue to the next iteration of the loop.
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
