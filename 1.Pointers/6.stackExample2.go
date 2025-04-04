package main

func createValue2() {
	var pointX *int
	pointX = createValue()
	print(pointX)
	return
}

func createValue() *int {
	x := 10
	innerPointer := &x
	return innerPointer
}
