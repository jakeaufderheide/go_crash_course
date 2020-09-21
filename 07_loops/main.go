package main

import "fmt"

func main() {
	// long version
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	// condensed
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d\n", i)
	}
}
