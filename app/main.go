package main

import (
	"fmt"

	"app/practice"
)

func exec_practice() {
	practice.Practice()
}

func main() {
	fmt.Println("Hello golang from docker!")
	// exec_practice()
	vec := []byte("hello")
	fmt.Println(vec)
}
