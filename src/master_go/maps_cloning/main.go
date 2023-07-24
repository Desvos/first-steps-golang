package main

import "fmt"

func main() {
	p := fmt.Println
	f := fmt.Printf

	friends := map[string]int{"Dan": 40, "Maria": 25}

	neighbors := friends //questa non è una copia, ma un reindirizzamento di memoria

	friends["Dan"] = 50

	p(neighbors)

	people := make(map[string]int)

	for k, v := range friends {
		people[k] = v
	}

	people["Anne"] = 18

	f("%#v ----------------- %#v", people, friends)

}
