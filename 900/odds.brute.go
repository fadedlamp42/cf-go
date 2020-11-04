//https://codeforces.com/problemset/problem/318/A
package main

import "fmt"

func main() {
	var k, n int
	x := 1
	fmt.Scanf("%d %d", &n, &k)
	for i := 1; i != k; i++ {
		x += 2
		if x > n {
			x = 2
		}
	}

	fmt.Println(x)
}
