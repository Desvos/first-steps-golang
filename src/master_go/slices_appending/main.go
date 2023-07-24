package main

import "fmt"

func main() {
	numbers := []int{2, 3}

	numbers = append(numbers, 10) //append crea un nuovo slice
	_ = numbers

	numbers = append(numbers, 20, 30, 40)

	n := []int{100, 200}
	numbers = append(numbers, n...)

	src := []int{10, 20, 30}
	dst := make([]int, 2)
	nn := copy(dst, src)
	fmt.Println(src, dst, nn)
}
