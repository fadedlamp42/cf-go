//https://codeforces.com/problemset/problem/617/A
package main

import "fmt"

func main() {
	steps := 0
	var x int
	fmt.Scanf("%d", &x)

	steps = x / 5
	x %= 5
	if x != 0 {
		steps++
	}

	fmt.Println(steps)
}
