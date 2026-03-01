package main

import "fmt"
import "time"

func main() {

	// if statement with short variable declaration
	if age := 25; age > 18 {
		fmt.Println("You are an adult.")
	} else if age > 12 {
		fmt.Println("You are a teenager.")
	} else {
		fmt.Println("You are a child.")
	}

	// switch statement
	switch i := 1; {
		case i % 2 == 0:
			fmt.Println("i is even")
		default:
			fmt.Println("i is odd")
	}


	// switch statement with multiple cases
	day := time.Now().Weekday()

	switch day {
		case time.Saturday, time.Sunday:
			fmt.Println("It's weekend!")
		default:
			fmt.Println("Working day.")
	}
}