package main

import "fmt"

var taskDone bool

func main() {
	var a = 30
	var b = 5.2

	// a = b wrong
	a = int(b)
	fmt.Println(a, b)

	var value int
	var price float64
	var name string
	var done bool

	fmt.Println(value, price, name, done)

	/*
		comment
	*/
}
