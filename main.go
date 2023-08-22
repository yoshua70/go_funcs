package main

import "fmt"

func main() {
	getFib := FibonacciDynamic()
	fmt.Println("Fibonacci of 25", getFib(25))
	fmt.Println("Fibonacci of 24", getFib(24))

	for i := 0; i <= 20; i++ {
		fmt.Println(getFib(i))
	}
}
