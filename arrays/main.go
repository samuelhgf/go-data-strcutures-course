package main

import "fmt"

func main() {
	fmt.Println("Array")

	// Initialize an empty array of integers [0 0 0 0 0 0 0 0] with 8 values
	var myArray [8]int

	// Initialize an populated array on integer [1 2 3 4 5 6 7 8]
	mySecondArray := [8]int{1, 2, 3, 4, 5, 6, 7, 8}

	// Initialize a slice.
	// Different from an array, a slice does not need to pass it length since it automatically resizes it
	mySlice := []int{1, 2, 3, 4, 5, 6, 7, 8}

	// Print the arrays
	fmt.Println(myArray)
	fmt.Println(mySecondArray)

	// Print the index [7] of populated array (8)
	fmt.Println(mySecondArray[7])

	// Iterate over the populated array
	for i, n := range mySecondArray {
		fmt.Printf("Index: %d\n", i)
		fmt.Println(n)
	}

	// Iterate over the populated slice
	for i, n := range mySlice {
		fmt.Printf("Index: %d\n", i)
		fmt.Println(n)
	}
}