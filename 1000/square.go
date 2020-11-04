//https://codeforces.com/problemset/problem/1/A
package main

import (
	"fmt"
	"math"
)

func main() {
	var n, m, size float64
	fmt.Scanf("%f %f %f", &n, &m, &size)
	x := math.Ceil(n / size)
	y := math.Ceil(m / size)
	fmt.Println(int64(x * y))
}
