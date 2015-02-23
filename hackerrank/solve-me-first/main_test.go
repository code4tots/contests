package main

import "testing"

func TestSolveMeFirst(t *testing.T) {
	cases := []struct {
		a, b, c uint32
	}{
		{1, 2, 3},
		{2, 3, 5},
	}

	for _, c := range cases {
		got := solveMeFirst(c.a, c.b)
		if got != c.c {
			t.Errorf("solveMeFirst(%d, %d) == %d, want %d", c.a, c.b, got, c.c)
		}
	}
}
