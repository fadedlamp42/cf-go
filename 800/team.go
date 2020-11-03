//https://codeforces.com/problemset/problem/231/A
package main

import (
	"fmt"
)

func main() {
	var team [3]int
	var n, total int
	fmt.Scanf("%d\n", &n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d %d %d\n", &team[0], &team[1], &team[2])
		if (team[0] + team[1] + team[2]) >= 2 {
			total++
		}
	}

	fmt.Println(total)
}
