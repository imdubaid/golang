package main

import "fmt"

func main() {
	fmt.Println("For loop example:")
	sum := 0
    // init; condition; post
    for i := 0; i < 5; i++ {
        sum += i
    }
	fmt.Println(sum)

	fmt.Println("While loop example:")
	count := 0
	for count < 5 {
		fmt.Println(count)
		count++
	}

	fmt.Println("Range loop example:")
	for i := range 5 {
		fmt.Println(i)
	}
}