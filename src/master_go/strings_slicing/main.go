package main

import "fmt"

func main() {
	s1 := "I love Golang!"
	fmt.Println(s1[2:5])

	s2 := "すぐにすぐにすぐに"
	fmt.Println(s1[0:2]) //torna lo slice dei byte

	rs := []rune(s2) //per tornare i caratteri nel console log bisogna convertire la stringa in uno slice di rune
	fmt.Printf("%T\n", rs)
	fmt.Println(string(rs)) //e poi convertirlo in string

}
