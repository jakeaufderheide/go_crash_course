package main

import "fmt"

func main() {
	ids := []int{33, 76, 54, 23, 11, 2}

	// using index
	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	// not using index
	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}

	// sum
	sum := 0
	for _, item := range ids {
		sum += item
	}
	fmt.Println(sum)
}
