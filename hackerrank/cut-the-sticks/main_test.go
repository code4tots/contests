package main

import "testing"

func equal(as, bs []int) bool {
	if len(as) != len(bs) {
		return false
	}

	for i, a := range as {
		if a != bs[i] {
			return false
		}
	}

	return true
}

func TestCutTheSticks(t *testing.T) {
	cases := []struct {
		in, want []int
	}{
		{
			[]int{1, 2, 3},
			[]int{3, 2, 1},
		},
		{
			[]int{5, 4, 4, 2, 2, 8},
			[]int{6, 4, 2, 1},
		},
		{
			[]int{1, 2, 3, 4, 3, 3, 2, 1},
			[]int{8, 6, 4, 1},
		},
	}

	for _, c := range cases {
		got := cutTheSticks(c.in)
		if !equal(got, c.want) {
			t.Errorf("cutTheSticks(%v) is %v, want %v", c.in, got, c.want)
		}
	}
}
