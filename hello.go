package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("Hello, World!\n")

	fmt.Printf("Param 1 => ", os.Args[1])
	fmt.Printf("Param 2 => ", os.Args[2])
}
