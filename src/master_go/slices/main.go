package main

import "fmt"

func main() {
	//####################################
	//SLICES

	var cc []string
	fmt.Println(cc == nil)
	fmt.Printf("cc %#v\n", cc)

	cc[0] = "London"

	numbers := []int{2, 3, 4, 5}
	_ = numbers

	nnn := make([]int, 2)
	_ = nnn

	type names []string
	friends := names{"Dan", "Maria"}

	myFriend := friends[0]
	_ = myFriend

	for index, value := range numbers {
		fmt.Printf("Index %v, value %v\n", index, value)
	}

	var n []int
	fmt.Println(n == nil)

	m := []int{}
	fmt.Println(m == nil)

	a, b := []int{1, 2, 3}, []int{1, 2, 3}

	//error -> fmt.Println(a==b)
	var eq bool = true

	if len(a) != len(b) {
		eq = false
	} else {
		for i, valueA := range a {
			if valueA != b[i] {
				eq = false
				break
			}
		}
	}

	if eq {
		fmt.Println("a and b are equal slices")
	} else {
		fmt.Println("a and b are not equal slices")
	}

}
