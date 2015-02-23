package main

import "fmt"

func read() []int {
	N := 0

	fmt.Scanf("%v", &N)

	sticks := make([]int, N)

	for i := 0; i < N; i++ {
		fmt.Scanf("%v", &sticks[i])
	}

	return sticks
}

func print(xs []int) {
	for _, x := range xs {
		fmt.Println(x)
	}
}

func min(xs []int) int {
	m := xs[0]
	for _, x := range xs {
		if x < m {
			m = x
		}
	}
	return m
}

func cutTheSticks(sticks []int) []int {
	cuts := make([]int, 0)
	for len(sticks) > 0 {
		cuts = append(cuts, len(sticks))
		smallest := min(sticks)
		remaining := make([]int, 0)
		for _, value := range sticks {
			if value-smallest > 0 {
				remaining = append(remaining, value-smallest)
			}
		}
		sticks = remaining
	}
	return cuts
}

func main() {
	print(cutTheSticks(read()))
}
