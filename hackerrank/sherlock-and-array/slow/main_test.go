package main

import "testing"

func TestSolve(t *testing.T) {
	cs := []struct {
		A []int
		expected bool
	} {
		{[]int {1, 2, 3}, false },
		{[]int {1, 2, 3, 3}, true},
	}

	for _, c := range cs {
		if solve(c.A) != c.expected {
			t.Errorf("solve(%v) != %v", c.A, c.expected)
		}
	}
}
