//https://codeforces.com/problemset/problem/266/A
package main

import "fmt"

func main() {
	var n int
	cNow := 'n'
	cLast := 'l'
	stones := 0

	fmt.Scanf("%d\n", &n)
	for i := 0; i < n; i++ {
		cLast = cNow
		fmt.Scanf("%c", &cNow)
		if cNow == cLast {
			stones++
		}
	}
	fmt.Println(stones)
}
