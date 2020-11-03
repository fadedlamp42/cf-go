//https://codeforces.com/problemset/problem/158/A
package main

import "fmt"

func main() {
	var n, k, total, kScore int

	fmt.Scanf("%d %d\n", &n, &k)
	var scores []int = make([]int, n)
	for i := range scores {
		fmt.Scanf("%d", &scores[i])

		if kScore > scores[i] {
			break
		}

		if (i + 1) == k {
			kScore = scores[i]
		}

		if scores[i] > 0 {
			total++
		}
	}

	fmt.Println(total)
}
