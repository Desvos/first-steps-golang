package main

import "fmt"

func main() {
	// if price {
	//error
	// }
	i := 0
todo:
	if i < 5 {
		i++
		fmt.Println(i)
		goto todo
	}

}
