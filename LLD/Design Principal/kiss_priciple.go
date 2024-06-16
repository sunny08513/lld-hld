package main

//Keep it Simple, Stupid

//Overly Complex:

func Factorial1(n int) int {
	if n < 0 {
		return -1 // Factorial is not defined for negative numbers
	}
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result
}

//KISS-aligned:

// use function if available already as library
