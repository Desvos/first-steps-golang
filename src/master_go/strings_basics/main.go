package main

import "fmt"

func main() {
	s1 := "Hi there go" //in double quotes è chiamato string literal
	fmt.Println(s1)

	fmt.Println("He says: \"Hello\"")

	fmt.Println(`He says: "Hello!"`) //raw strings

	s2 := `I like \n Go` //i caratteri speciali non funzionano
	fmt.Println(s2)

	//in compenso ti calcola l'accapo
	fmt.Println(`
		Price 100
		Brand Nike
	`)

	fmt.Println(`C:\Users\devcot`)   //in raw i back slash si possono usare singolarmente
	fmt.Println("C:\\Users\\devcot") // in literal devi fare l'escape dei backslash

	var s3 = "I love " + "Go " + "Programming"
	fmt.Println(s3 + "!")
	fmt.Println(s3[0]) //ritorna 73 (in decimal), questo perchè in go le stringhe sono slices of bytes

	//s3[5] = "x"  //errore, perchè le stringhe sono immutabili e non puoi cambiare un carattere senza creare una nuova stringa

}
