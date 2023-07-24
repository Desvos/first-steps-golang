package main

import "fmt"

func main() {
	//INT TYPE
	var i1 int8 = 100
	fmt.Printf("%T\n", i1)

	//UINT TYPE
	var i2 uint16 = 6553
	fmt.Printf("%T\n", i2)

	//FLOAT TYPE
	var f1, f2, f3 float64 = 1.1, -.2, 5. //1.1	-0.2	5.0
	fmt.Printf("%T %T %T \n", f1, f2, f3)

	//RUNE TYPE
	var r rune = 'f'      //rune is a alias for int32
	fmt.Printf("%T\n", r) // prints int32
	fmt.Println(r)        //prints 102
	fmt.Printf("%x\n", r) //prints 66, decimal ASCII code for f

	//BOOL TYPE
	var b bool = true
	fmt.Printf("%T\n", b)

	//STRING TYPE
	var s string = "Hello Go!"
	fmt.Printf("%T\n", s)

	//ARRAY TYPE
	var numbers = [4]int{2, 3, -4, 5}
	_ = numbers

	//SLICE TYPE
	var cities = []string{"London", "Tokyo", "NewYork"}
	_ = cities

	//MAP TYPE
	balances := map[string]float64{
		"USD": 232.2,
		"EUR": 292.1,
	}
	_ = balances

	//STRUCT TYPE
	type Person struct {
		name string
		age  int
	}
	var you Person
	_ = you

	//POINTER TYPE
	var x int = 2
	ptr := &x
	_ = ptr

	//FUNCTION TYPE
	fmt.Printf("%T\n", f)

	type speed uint
	var s1 speed = 10
	var s2 speed = 20
	_ = s1 - s2

	var xx uint
	//xx = s1 error

	xx = uint(s1)
	_ = xx

	var s3 speed
	s3 = speed(xx)
	_ = s3

	//alias
	type second = uint

}

func f() {

}
