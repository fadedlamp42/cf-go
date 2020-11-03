//https://codeforces.com/problemset/problem/50/A
package main

import "fmt"

func main() {
	var m, n int
	fmt.Scanf("%d %d\n", &m, &n)
	fmt.Println((m/2)*n + (m%2)*n/2)
}
