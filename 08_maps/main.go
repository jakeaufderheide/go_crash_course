package main

import "fmt"

func main() {
	dict := make(map[string]string)
	dict["test"] = "test"
	dict2 := map[string]string{"test2": "test2"}
	fmt.Println(dict)
	fmt.Println(dict2)
	fmt.Println(len(dict))
}
