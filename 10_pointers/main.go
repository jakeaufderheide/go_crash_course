package main

import "fmt"

func main() {
	a := 5
	b := &a

	fmt.Println(a, b)
	fmt.Printf("%T\n", b)

	// use * to read val from address
	fmt.Println(*b)

	// Change val with pointer
	*b = 10
	fmt.Println(a)

	// research suggests that passing by pointer reference can help performance when passing large structs
}
