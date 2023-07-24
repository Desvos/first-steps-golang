package main

import "fmt"

func main() {
	var numbers [4]int

	fmt.Printf("%v\n", numbers)
	fmt.Printf("%#v\n", numbers)

	var a1 = [4]float64{}
	fmt.Printf("%#v\n", a1)

	var a2 = [3]int{-10, 1, 100}
	fmt.Printf("%#v\n", a2)

	a3 := [4]string{"Dan", "Diana", "Paul"}
	fmt.Printf("%#v\n", a3)

	//ellipsis operator
	a4 := [...]int{1, 2, 3, 4, 5}
	fmt.Printf("%#v\n", a4)
	fmt.Printf("%d\n", len(a4))

	balances := [2][3]int{
		{5, 6, 7},
		[3]int{8, 9, 10},
	}
	fmt.Println(balances)

	m := [3]int{1, 2, 3}
	n := m
	fmt.Println(n == m)
	m[0] = -1
	fmt.Println(n == m)

	//puoi
	grades := [3]int{
		1: 10,
		0: 5,
		2: 7,
	}
	fmt.Println((grades))

	accounts := [3]int{2: 50}
	fmt.Println(accounts)

	names := [...]string{
		5: "Dan",
	}
	fmt.Println(names)

	cities := [...]string{
		5: "Paris",
		"London",
		1: "NYC",
	}
	fmt.Printf("%#v", cities) //london è sull'indice 6

	weekend := [7]bool{5: true, 6: true}
	fmt.Println(weekend)
}
