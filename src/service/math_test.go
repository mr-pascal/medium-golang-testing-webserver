package main

import (
	"testing"
)

func TestSum(t *testing.T) {
	// Simple test case, checking one specific case
	total := Sum(5, 5)
	if total != 10 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
	}
}

func TestMultiply(t *testing.T) {

	// Create the test table
	tables := []struct {
		x int
		y int
		r int
	}{
		{x: 1, y: 2, r: 2}, // normal multiplication with named struct parameters
		{2, 3, 6},          // normal multiplication with anonymous struct parameters
		{0, 0, 0},          // both zero
		{2, 0, 0},          // x is zero
		{0, 2, 0},          // y is zero
		{-5, 2, -10},       // x is negative
		{5, -2, -10},       // y is negative
		{-5, -2, 10},       // x and y are negative
	}

	for _, table := range tables {
		// Get the return value of the function
		total := Multiply(table.x, table.y)

		// Assert the result with our expected value
		if total != table.r {
			// Doesn't match our expected result
			// -> Print error indicating what was the input, what we got and what we have expected
			t.Errorf("Multiplication of ( %d * %d ) was incorrect, got: %d, want: %d.", table.x, table.y, total, table.r)
		}
	}
}
