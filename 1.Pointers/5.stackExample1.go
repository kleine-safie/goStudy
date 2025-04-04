package main

func createAndUpdateValue1() {
	x := 10
	updateValue1(&x)
	return
}

func updateValue1(a *int) {
	*a = 11
}