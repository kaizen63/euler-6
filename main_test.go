package main

import (
	"math"
	"testing"
)

func TestSumOfSquares(t *testing.T) {
	s := math.Pow(1, 2) +
		math.Pow(2, 2) +
		math.Pow(3, 2) +
		math.Pow(4, 2) +
		math.Pow(5, 2) +
		math.Pow(6, 2) +
		math.Pow(7, 2) +
		math.Pow(8, 2) +
		math.Pow(9, 2) +
		math.Pow(10, 2) +
		math.Pow(11, 2) +
		math.Pow(12, 2)

	if r := sumOfSquares(12); r != s {
		t.Errorf("sumOfSquares(12) = %f, expected %f", r, s)
	}
}

func TestSumOfSquaresNegativeArgument(t *testing.T) {
	if r := sumOfSquares(-1); r != 0 {
		t.Errorf("sumOfSquares(-1) = %f, expected 0", r)
	}
}
func TestSquareOfSums(t *testing.T) {
	s := math.Pow(float64(1+2+3+4+5+6+7+8+9+10+11+12), 2)
	if r := squareOfSums(12); r != s {
		t.Errorf("squareOfSums(12) = %f, expected %f", r, s)
	}
}
func TestSquareOfSumsNegativeArgument(t *testing.T) {
	if r := squareOfSums(-1); r != 0 {
		t.Errorf("squareOfSums(-1) = %f, expected 0f", r)
	}
}
