//https://codeforces.com/problemset/problem/69/A
package main

import "fmt"

func main() {
	var (
		in, sum [3]int64
		n       int
	)

	fmt.Scanf("%d\n", &n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d %d %d\n", &in[0], &in[1], &in[2])
		sum[0] += in[0]
		sum[1] += in[1]
		sum[2] += in[2]
	}

	if sum[0] == 0 && sum[1] == 0 && sum[2] == 0 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
