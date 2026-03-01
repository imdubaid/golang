package main

import "fmt"


func main() {
	var variable string = "Variable declaration"
	shorthand := "Short hand variable declaration"
	const constant string = "Constant declaration"
	const (
		port = 8000
		host = "localhost"
	)


	fmt.Println("Hello, World!")
	fmt.Println(variable, "and", shorthand)
	fmt.Println(constant, "cannot be changed", port, "and", host)
}