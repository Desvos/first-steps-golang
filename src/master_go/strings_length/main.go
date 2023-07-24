package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s1 := "Golang"
	fmt.Println(len(s1))

	name := "Condruta"
	fmt.Println(len(name)) //restituisce 8, la lunghezza in byte della stringa

	n := utf8.RuneCountInString(name)
	fmt.Println(n) //ritorna 7, il numero di caratteri della stringa
}
