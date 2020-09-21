package main

import "fmt"

func main() {
	// Using var

	var age int32 = 37
	const isCool = true

	// shorthand
	name := "Brad"
	size := 1.3

	fmt.Println(name, age, isCool)
	fmt.Printf("%T\n", size)
}
