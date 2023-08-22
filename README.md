# Go Funcs

A collection of functions written in the Go programming language.

## FibonacciDynamic

```go
FibonacciDynamic() func(y int) int
```

Compute the nth element of the Fibonacci sequence using dynamic programming.

### Usage

This is a function which returns another function. You first need to store the returned function in a variable and then use it to compute the elements of the Fibonacci sequence.

```go
getFib := FibonacciDynamic()
fmt.Println("Fibonacci of 25", getFib(25))
fmt.Println("Fibonacci of 24", getFib(24))
```
