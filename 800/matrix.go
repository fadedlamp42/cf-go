//https://codeforces.com/problemset/problem/263/A
package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	for row := 0; row < 5; row++ {
		for col := 0; col < 5; col++ {
			fmt.Scanf("%d", &n)
			if n == 1 {
				fmt.Println(math.Abs(float64(row)-2) + math.Abs(float64(col)-2))
				return
			}
		}
		fmt.Scanln()
	}
}
