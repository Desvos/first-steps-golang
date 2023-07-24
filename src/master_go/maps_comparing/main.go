package main

import "fmt"

func main() {
	p := fmt.Println

	a := map[string]string{"A": "X"}

	b := map[string]string{"B": "Y"}

	//p(a == b)		//error

	s1 := fmt.Sprintf("%s", a)
	s2 := fmt.Sprintf("%s", b)

	p(s1, s2)

	if s1 == s2 {
		fmt.Println("Maps are equal")
	} else {
		p("Maps are not equal")
	}

}
