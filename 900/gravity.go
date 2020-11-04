//https://codeforces.com/problemset/problem/405/A
package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scanf("%d\n", &n)
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &nums[i])
	}
	sort.Ints(nums)
	for _, x := range nums {
		fmt.Printf("%d ", x)
	}
}
