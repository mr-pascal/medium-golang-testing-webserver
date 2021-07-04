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
	// Using test tables
	tables := []struct {
		x int
		y int
		r int
	}{
		{x: 1, y: 1, r: 1},
		{1, 2, 2},
		{0, 0, 0},
		{2, 0, 0},
		{-5, 2, -10},
	}

	for _, table := range tables {
		total := Multiply(table.x, table.y)
		if total != table.r {
			t.Errorf("Multiplication of ( %d * %d ) was incorrect, got: %d, want: %d.", table.x, table.y, total, table.r)
		}
	}
}
