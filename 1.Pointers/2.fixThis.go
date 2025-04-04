package main

import "fmt"

func createAndUpdateValue() {
	x := 10
	updateValue(&x)
	fmt.Println(x)
}

func updateValue(a *int) {
	b := *a + 1
	*a = b
}
