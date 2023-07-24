package main

import "fmt"

func main() {
	const (
		c1 = iota
		c2 = iota
		c3 = iota
	)

	fmt.Println(c1, c2, c3) //printa 0,1,2

	const (
		cc1 = iota
		cc2
		cc3
	)

	fmt.Println(cc1, cc2, cc3) //printa 0,1,2

	const (
		a = iota * 2 //in step of 2 = 0
		b            //2
		c            //4
	)

	const (
		x = -(iota + 2) //-2
		_               //-3
		y               //-4
		z               //-5
	)

}
