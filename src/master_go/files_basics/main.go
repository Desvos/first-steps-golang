package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	var newFile *os.File //utilizzando i pointer

	fmt.Printf("%T\n", newFile)

	var err error

	newFile, err = os.Create("a.txt")

	if err != nil {
		// p(err)
		// os.Exit(1)	//questo per uscire dal programma una volta avuto l'errore
		log.Fatal(err) //idiomatic method to handle error in go
	}

	err = os.Truncate("a.txt", 0) //svuota il file
	//err = os.Truncate("a.txt", 50) //cancella 50 bytes

	if err != nil {
		log.Fatal(err)
	}
	newFile.Close()

	var file *os.File
	file, err = os.Open("a.txt")
	if err != nil {
		log.Fatal(err)
	}
	file.Close()

	file, err = os.OpenFile("a.txt", os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()

	var fileInfo os.FileInfo
	fileInfo, err = os.Stat("a.txt")

	p := fmt.Println

	p("File Name:", fileInfo.Name())
	p("Size in bytes:", fileInfo.Size())
	p("Last modified:", fileInfo.ModTime())
	p("Is Directory?", fileInfo.IsDir())
	p("Permissions:", fileInfo.Mode())

	fileInfo, err = os.Stat("b.txt")
	if err != nil {
		if os.IsNotExist(err) {
			p("The file does not exist!!")
		}
	}

	oldPath := "a.txt"
	newPath := "aaa.txt"
	err = os.Rename(oldPath, newPath)
	if err != nil {
		log.Fatal(err)
	}

	err = os.Remove("aaa.txt")
	if err != nil {
		log.Fatal(err)
	}
}
