package main

import (
	"fmt"
	"unsafe"
)

func main() {
	s1 := []int{10, 20, 30, 40, 50}

	s3, s4 := s1[0:2], s1[1:3]

	s3[1] = 600
	//in questo caso entrambi gli slices avranno l'elemento che corrisponde al 20 modificato in 600
	fmt.Println(s1)
	fmt.Println(s3)
	fmt.Println(s4)

	arr1 := [4]int{10, 20, 30, 40}
	slice1, slice2 := arr1[0:2], arr1[1:3]

	//anche in questo caso gli elementi corrispondenti al 20 vengono sostituiti in 2,
	//perchè il backing array arr1, l'array che si forma come copia per creare lo slice, viene modificato
	arr1[1] = 2
	_, _ = slice1, slice2



	//NEW SLICES
	//to create a new slice from another slice we should use append method
	cars := []string{"Ford", "Honda", "audi", "fiat"}
	newCars := []string{}

	newCars = append(newCars, cars[0:2]...)

	cars[0] = "Nissan"

	fmt.Println(cars, newCars)

	slicee := []int{10,20,30,40,50}
	newSlicee := slicee[0:3]
	fmt.Println(len(newSlicee), cap(newSlicee))	//torna 3 5, perchè il capacity conta tutti gli elementi dall'indice di inizio fino alla fine

	newSlicee = slicee[2:3]
	fmt.Println(len(newSlicee), cap(newSlicee))	//torna 3 3, perchè il capacity conta tutti gli elementi dall'indice di inizio fino alla fine

	//con queste istruzioni si stampano gli indirizzi di memoria
	//si vuol dimostrare che anche se due slices hanno due indirizzi di memoria differenti, essi potrebbero avere lo stesso backing array
	fmt.Printf("%p\n", &slicee)

	fmt.Printf("%p %p\n", &slicee, &newSlicee)

	newSlicee[0] = 1000
	fmt.Println("slicee: ",slicee)

	//ora vediamo quale tipo di variabile (array o slice) è più pesante
	a := [5]int{1,2,3,4,5}
	s := []int{1,2,3,4,5}
	fmt.Printf("a size in bytes is: %d \n", unsafe.Sizeof(a))
	fmt.Printf("s size in bytes is: %d \n", unsafe.Sizeof(s))
	//gli array occupano più spazio degli slice

	

}
