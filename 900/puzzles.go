//https://codeforces.com/problemset/problem/337/A
package main

import (
	"fmt"
	"sort"
)

func main() {
	var count, n int
	fmt.Scanf("%d %d\n", &count, &n)
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &nums[i])
	}
	sort.Ints(nums)

	minimum := nums[count-1] - nums[0]
	for i := 0; i <= n-count; i++ {
		if nums[i+count-1]-nums[i] < minimum {
			minimum = nums[i+count-1] - nums[i]
		}
	}

	fmt.Println(minimum)
}
