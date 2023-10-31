package main

import "fmt"

func main() {
	fmt.Println("OK1") // OK1
	goto label
	fmt.Println("OK2")
	fmt.Println("OK3")
	fmt.Println("OK4")
label:
	fmt.Println("OK5") // OK5
	fmt.Println("OK6") // OK6
}
