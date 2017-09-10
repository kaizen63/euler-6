package main

import (
	"fmt"
	"math"
)

func sumOfSquares(n int) float64 {
	var sum float64
	for i := 1; i <= n; i++ {
		sum += math.Pow(float64(i), 2)
	}
	return sum
}

func squareOfSums(n int) float64 {
	if n < 0 {
		return 0
	}
	// Gauss: sum(1..n) = n*(n+1)/2
	return math.Pow(float64(n)*float64(n+1)/2.0, 2.0)
}

/*
The sum of the squares of the first ten natural numbers is,

1^2 + 2^2 + ... + 10^2 = 385
The square of the sum of the first ten natural numbers is,

(1 + 2 + ... + 10)^2 = 55^2 = 3025
Hence the difference between the sum of the squares of the first ten natural numbers and the square of the sum is 3025 − 385 = 2640.

Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum.
*/

func main() {
	for i := 1; i < 100; i++ {
		sumOfSquares := sumOfSquares(i)
		squareOfSums := squareOfSums(i)
		fmt.Printf("%-3d %6.0f %8.0f %8.0f\n",
			i, sumOfSquares, squareOfSums, squareOfSums-sumOfSquares)
	}
}
