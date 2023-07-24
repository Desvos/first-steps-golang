package main

import "fmt"

func main() {
	var nums []int
	fmt.Printf("%#v\n", nums)

	fmt.Printf("Length: %d, Capacity: %d \n", len(nums), cap(nums)) // 0, 0 da ricordarsi che capacity è la dimensione del backing array

	nums = append(nums, 1, 2)                                       //append crea un nuovo backing array con maggiore capacità
	fmt.Printf("Length: %d, Capacity: %d \n", len(nums), cap(nums)) // 2, 2

	nums = append(nums, 3)
	fmt.Printf("Length: %d, Capacity: %d \n", len(nums), cap(nums)) // 3, 4

	nums = append(nums, 4)
	fmt.Printf("Length: %d, Capacity: %d \n", len(nums), cap(nums)) // 4, 4, questo perchè la capacità MASSIMA viene incrementata esponenzialmente in base 2, quindi 2,4,8,16,32

	nums = append(nums, 5)
	fmt.Printf("Length: %d, Capacity: %d \n", len(nums), cap(nums)) // 5, 8

	nums = append(nums[0:4], 200, 300, 400, 500, 600)
	fmt.Printf("Length: %d, Capacity: %d \n", len(nums), cap(nums)) // 9, 16
	fmt.Println(nums)

	letters := []string{"A", "B", "C", "D", "E", "F"}
	letters = append(letters[0:1], "X", "Y")

	//fmt.Println(letters[4]) //out of range
	fmt.Println(letters[3:6])
}
