//https://codeforces.com/problemset/problem/318/A
package main

import (
	"fmt"
	"math"
)

func main() {
	var k, n float64
	fmt.Scanf("%f %f", &n, &k)

	if k <= math.Ceil(n/2) {
		fmt.Println(int64(1 + 2*(k-1)))
	} else {
		fmt.Println(int64(2 * (k - math.Ceil(n/2))))
	}
}
