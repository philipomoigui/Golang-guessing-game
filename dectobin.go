package main

import (
	"fmt"
)

func main() {
	var input int
	fmt.Println("Enter Your Decimal Number")
	fmt.Scan(&input)
	fmt.Printf("The binary of your input is %b\n", input)
}