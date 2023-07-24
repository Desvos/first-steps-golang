package main

import (
	"fmt"
	"math"
)

func main() {
	//il compilatore non da errore se definiamo una costante senza utilizzarla
	const days int = 7

	var i int
	fmt.Println(i)

	//is mandatory assign a value to a constant
	const pi float64

	x, y := 5, 0

	fmt.Println(x / y) //questo non è un errore a compile time ma a runtime

	//se invece utilizziamo le costanti, lo stesso errore viene individuato a compile time
	const a = 5
	const b = 0

	fmt.Println(a / b)

	const n, m int = 4, 5
	const n1, m1 = 6, 7

	const (
		min1 = -500
		min2 = -300
		min3 = 100
	)

	const (
		min4 = -500
		min5
		min6
	) // saranno tutte uguali a -500

	//non puoi inizializzare una costante
	const power = math.Pow(2, 3)

	//non puoi inizializzare una costante a partire da una variabile
	t := 5
	const tc = t

	const l1 = len("hello") //questo va
}
