/*
Kyumins-MBP:sherlock-and-array math4tots$ time eval 'cat input.txt | go run main.go'
NO
NO
NO
NO
NO
NO
NO
NO
NO
NO

real	0m1.706s
user	0m0.860s
sys	0m1.133s
*/
package main

import "fmt"

func read() []int {
	var N int
	fmt.Scanf("%v", &N)
	A := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scanf("%v", &A[i])
	}
	return A
}

func print(answer bool) {
	if answer {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func solve(A []int) bool {
	total := 0
	for _, a := range A {
		total += a
	}
	so_far := 0
	for _, a := range A {
		if so_far == total - so_far - a {
			return true
		}
		so_far += a
	}
	return false
}

func main() {
	var T int
	fmt.Scanf("%v", &T)
	for t := 0; t < T; t++ {
		print(solve(read()))
	}
}
