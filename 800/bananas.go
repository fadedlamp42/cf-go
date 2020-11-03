//https://codeforces.com/problemset/problem/546/A
package main

import (
	"fmt"
	"math"
)

func main() {
	var price, have, count, total int
	fmt.Scanf("%d %d %d", &price, &have, &count)
	for i := 1; i <= count; i++ {
		total += i * price
	}

	fmt.Println(int(math.Max(0, float64(total-have))))
}
