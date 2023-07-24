package main

import (
	"fmt"
)

func main() {
	fmt.Printf("Your age is %d\n", 21)

	fmt.Printf("x is %d, y is %f", 1, 2.5) // %d for decimals, %f for float

	figure := "Circle"
	radius := 5
	pi := 3.14159

	_, _ = figure, pi

	fmt.Printf("Radius is %+d\n", radius) // %+d per visionare anche il segno + ( se il segno è - rimane -)

	fmt.Printf("The figure is %s", figure) // string

	fmt.Printf("The figure is %q", figure) // for quoted strings -> "Circle"

	fmt.Printf("The figure is %v", figure) // %v for any values

	fmt.Printf("The figure is of type %T", figure) // %T for show the type of a variable

	// %t for bool
	closed := true
	fmt.Printf("File closed: %t\n", closed)

	// % -> base 2, byte
	fmt.Printf("%b \n", 55) // Result: 110111

	// %(0num)b -> base 2 with leading 0
	fmt.Printf("%08b \n", 55) // Result: 00110111

	// %x -> hexadecimal
	fmt.Printf("%x \n", 100) // Result: 64

	x := 3.4
	y := 6.9

	// per mostrare 3 decimal after point
	fmt.Printf("x * y = %.3f\n", x*y)
}
