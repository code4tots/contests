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

real	0m2.783s
user	0m1.694s
sys	0m1.401s

OK... So this is even slower...

So couldn't get the Go version to pass the judge...
*/
package main

import "fmt"

var A [100000]int
var N int

func read() {
	fmt.Scanf("%v", &N)
	for i := 0; i < N; i++ {
		fmt.Scanf("%v", &A[i])
	}
}

func print(answer bool) {
	if answer {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func solve() bool {
	total := 0
	for i := 0; i < N; i++ {
		total += A[i]
	}
	so_far := 0
	for i := 0; i < N; i++ {
		if so_far == total - so_far - A[i] {
			return true
		}
		so_far += A[i]
	}
	return false
}

func main() {
	var T int
	fmt.Scanf("%v", &T)
	for t := 0; t < T; t++ {
		read()
		print(solve())
	}
}
