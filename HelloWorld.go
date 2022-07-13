package main

import "fmt"

const ERROR_CODE int = 400

func main() {
	printLn(2)
}

func printLn(x int) {
	a := 1
	sum := 3
	fmt.Println(a+sum, ERROR_CODE)
}
