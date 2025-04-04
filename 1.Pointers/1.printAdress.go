package main

import (
	"fmt"
)

//Explicit Pointer:
//  * + concrete type (primitives, struct or array)
//  Access it by getting the address (&)

func printAddress() {

	x := 10
	y := true
	pointerToX := &x
	pointerToY := &y
	fmt.Println(pointerToX)
	fmt.Println(*pointerToX)
	fmt.Println(pointerToY)
	fmt.Println(*pointerToY)

	z := 3 + *pointerToX
	fmt.Println(z)

	// wrong
	//z := 3 + pointerToX
}

/*

type定義とパラメータ定義に * が pointer
ex:
var x *int

関数内に * が deference (pointが指す先を参照する)
ex:
*x = 100

&はアドレスを取得する

// ex:
func makeX100 (x *int) {
 *x = 100
}

x := 10
makeX100(&x)

// 42行の *intは、int型のポインタを表す
// 43行の *xは、ポインタxが指す先の値を参照することを意味する
// 47行の &xは、変数xのアドレスを取得することを意味する
*/