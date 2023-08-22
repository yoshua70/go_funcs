package main

import "testing"

func TestFibonacciOf0ShouldReturn0(t *testing.T) {
	x := 0
	want := 0

	getFibDyn := FibonacciDynamic()
	have := getFibDyn(x)

	if want != have {
		t.Fatalf(`FibonacciDynamic(%d) = %d, want equal for %d`, x, have, want)
	}
}

func TestFibonacciOf1ShouldReturn1(t *testing.T) {
	x := 1
	want := 1

	getFibDyn := FibonacciDynamic()
	have := getFibDyn(x)

	if want != have {
		t.Fatalf(`FibonacciDynamic(%d) = %d, want equal for %d`, x, have, want)
	}
}

func TestFibonacciOf10ShouldReturn55(t *testing.T) {
	x := 10
	want := 55

	getFibDyn := FibonacciDynamic()
	have := getFibDyn(x)

	if want != have {
		t.Fatalf(`FibonacciDynamic(%d) = %d, want equal for %d`, x, have, want)
	}
}
