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

func main() {
	for i := 1; i < 100; i++ {
		sumOfSquares := sumOfSquares(i)
		squareOfSums := squareOfSums(i)
		fmt.Printf("%-3d %6.0f %8.0f %8.0f\n",
			i, sumOfSquares, squareOfSums, squareOfSums-sumOfSquares)
	}
}
