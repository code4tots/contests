/*
C version for comparison

Kyumins-MBP:sherlock-and-array math4tots$ gcc main.c
Kyumins-MBP:sherlock-and-array math4tots$ time eval 'cat input.txt | ./a.out'
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

real  0m0.141s
user  0m0.135s
sys 0m0.007s

OK. So C version is more than 10x faster than the faster Go version.

*/
#include <stdio.h>

int A[100000], N;

void read() {
  int i;
  scanf("%d", &N);
  for (i = 0; i < N; i++)
    scanf("%d", A+i);
}

int solve() {
  int i, total = 0, so_far = 0;
  for (i = 0; i < N; i++)
    total += A[i];
  for (i = 0; i < N; i++) {
    if (so_far == total - so_far - A[i])
      return 1;
    so_far += A[i];
  }
  return 0;
}

int main() {
  int T, i;
  scanf("%d", &T);
  for (i = 0; i < T; i++) {
    read();
    printf("%s\n", solve() ? "YES" : "NO");
  }
}
