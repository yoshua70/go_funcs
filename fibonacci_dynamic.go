package main

func FibonacciDynamic() func(y int) int {
	var fib []int
	fib = append(fib, 0, 1)

	return func(y int) int {
		if y < 0 {
			panic("Negative number not supported.")
		}
		// We already computed the value.
		if y < int(len(fib)) {
			return fib[y]
		}

		// We have the two previous values needed
		// for computation.
		if y-2 >= 0 && y < len(fib) {
			fib = append(fib, fib[y-1]+fib[y-2])
			return fib[y]
		}

		// We compute a new value.
		for i := len(fib); i <= y; i++ {
			fib = append(fib, fib[i-1]+fib[i-2])
		}
		return fib[y]
	}
}
