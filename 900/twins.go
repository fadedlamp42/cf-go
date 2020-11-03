//https://codeforces.com/problemset/problem/160/A
package main

import (
	"fmt"
	"sort"
)

func tryTake(n int, coins []int) bool {
	var sum1, sum2 int
	for i, val := range coins {
		if i < n {
			sum1 += val
		} else {
			sum2 += val
		}
	}

	return sum1 > sum2
}

func main() {
	var n int
	fmt.Scanf("%d\n", &n)
	coins := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &coins[i])
	}
	sort.Ints(coins)
	sort.Sort(sort.Reverse(sort.IntSlice(coins)))

	take := 1
	for !tryTake(take, coins) {
		take++
	}

	fmt.Println(take)
}
