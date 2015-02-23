/*
C++ version for performance comparison

Kyumins-MBP:sherlock-and-array math4tots$ g++ main.cpp --std=c++11
Kyumins-MBP:sherlock-and-array math4tots$ time eval 'cat input.txt | ./a.out'
NO
NO
NO
YES
YES
YES
YES
YES
YES
YES

real  0m0.631s
user  0m0.624s
sys 0m0.009s

Kyumins-MBP:sherlock-and-array math4tots$ g++ main.cpp --std=c++11 -O3
Kyumins-MBP:sherlock-and-array math4tots$ time eval 'cat input.txt | ./a.out'
NO
NO
NO
YES
YES
YES
YES
YES
YES
YES

real  0m0.599s
user  0m0.585s
sys 0m0.008s

Performance boost not as great as I thought it would be.

*/
#include <iostream>
#include <vector>
using namespace std;

/*
Shouldn't be that slow. Consider implicit move semantics.
*/
vector<int> read() {
  int N;
  vector<int> A;
  cin >> N;
  for (int i = 0; i < N; i++) {
    int a;
    cin >> a;
    A.push_back(a);
  }
  return A;
}

bool solve(vector<int> A) {
  int total = 0;
  for (auto a : A)
    total += a;
  int so_far = 0;
  for (auto a : A) {
    if (so_far == total - so_far - a)
      return true;
    so_far += a;
  }
  return false;
}

int main() {
  int T;
  cin >> T;
  for (int i = 0; i < T; i++) {
    cout << (solve(read()) ? "YES" : "NO") << endl;
  }
}
