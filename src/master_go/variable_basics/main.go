package main

import (
	"fmt"
	"math"
)

func main() {
	var age int = 30
	fmt.Println("Age:", age)

	var name = "Dan"
	//fmt.Println("Name:", name)
	_ = name

	s := "Learning Golang"
	fmt.Println(s)

	car, cost := "Audi", 50000
	fmt.Println(car, cost)

	//il := può essere usato anche per dichiarare e riassegnare due variabili insieme
	car, year := "BMV", 2021
	//ma non riassegnare variabili già inizializzate
	//car, year := "BMV", 2019
	_ = year

	var opened = false
	opened, file := true, "a.text"
	_, _ = opened, file

	var (
		salary    float64
		firstName string
		gender    bool
	)
	fmt.Print(salary, firstName, gender)

	var a, b, c int
	fmt.Println(a, b, c)

	//multiple assignment
	var i, j int
	i, j = 5, 8
	i, j = j, i // swapping variables
	fmt.Println(i, j)

	sum := 5 + 2.3
	fmt.Println(sum)

	var max_value int       //not OK
	var packetsReceived int //NOT OK, name too long

	const MAX_VALUE = 100 //NOT OK
	const N = 100         // OK

	var Max = 100 // variable to be exported
	_ = Max

	maxValue := 1000   // recommended
	max_values := 1000 // NOT recommended

	writeToDB := true // idiomatic
	writeToDb := true // not idiomatic

	var x = 1
	x++
	x--
	x += 1
	x -= 1
	x *= 1
	x /= 1
	x %= 1
	//fmt.Println(5+x--) non va

	//Overflow
	var p uint = 255
	p++ //ritorna 0 perchè va in overflow

	abc := int8(255 + 1)

	var bbb int8 = -128
	bbb-- //arriva a 127

	f := math.MaxFloat32

	f = f + 1.2 //ritorna +Inf

	type age1 int
	type oldAge age1

}
