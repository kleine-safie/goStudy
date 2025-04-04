package main

import (
	"fmt"
)

/*
Implicit Pointer
-slice
-map
-function
-channel
-interface

Pointer start value is nil
*/

func modMap(m map[int]string) {
	m[2] = "hello"
	m[3] = "goodbye"
	delete(m, 1)
}

func createMap() {

	m := map[int]string{
		1: "first",
		2: "second",
	}

	m2 := m

	fmt.Printf("m: %v\n", m)
	fmt.Printf("m2: %v\n", m2)

	modMap(m)

	fmt.Printf("m: %v\n", m)
	fmt.Printf("m2: %v\n", m2)

}