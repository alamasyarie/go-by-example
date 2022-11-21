package main

import "fmt"

func Condition() {
	if 7%2 == 0 { // Basic conditional
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	if 8%4 == 0 { // You cah have single condition without else
		fmt.Println("8 is divisible by 4")
	}

	if num := 9; num < 0 { // A statement can precede conditionals
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}
