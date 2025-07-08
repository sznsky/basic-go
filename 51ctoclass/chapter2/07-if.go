package main

import "fmt"

func IFTest() {

	a := 10
	if a > 10 {
		fmt.Printf("a biger than 10")
	} else if a == 10 {
		fmt.Printf("a equal 10")
	} else {
		fmt.Printf("a less than 10")
	}
}
