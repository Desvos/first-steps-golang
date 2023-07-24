package main

import "fmt"

func main() {
	p := fmt.Println
	f := fmt.Printf

	var employees map[string]string

	f("%#v\n", employees)

	f("Num of pairs: %d\n", len(employees))

	f("The value for key %q is %q", "Dan", employees["Dan"]) //ritorna il 0 value, "" in questo caso, quando non è presente un elemento con quella chiave nella mappa

	var accounts map[string]float64
	f("%#v\n", accounts["0x323"])

	//var m2 map[[]int]string  //slices as key are not admitted
	var m2 map[[5]int]string
	_ = m2

	//employees["Dan"] = "Programmer" //non è possibile aggiungere valori ad una mappa non inizializzata

	people := map[string]float64{} //questa è una mappa inizializzata

	people["John"] = 2.4
	people["Marry"] = 2.3

	p(people)

	map1 := make(map[int]int) //anche questo inizializza una mappa
	map1[4] = 9

	balances := map[string]float64{
		"USD": 323.3,
		"EUR": 999.2,
	}
	p(balances)

	m := map[int]int{1: 10, 2: 20, 3: 30}
	_ = m

	balances["USD"] = 212.1
	balances["GBP"] = 999

	v, ok := balances["RON"] //in questo tipo di assegnazione, la prima variabile rappresenta il valore della chiave, il secondo valore (var ok) rappresenta un booleano che indica se la variabile esiste o no

	if ok {
		p("The RON balance is: ", v)
	} else {
		p("The RON key doesn't exist in the map")
	}

	for k, v := range balances {
		f("Key %#v, Value %#v\n", k, v)
	}

	delete(balances, "USD")

}
