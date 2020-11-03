//https://codeforces.com/problemset/problem/791/A
package main

import "fmt"

func main() {
	var a, b, years int
	fmt.Scanf("%d %d\n", &a, &b)

	for a <= b {
		years++
		a *= 3
		b *= 2
	}

	fmt.Println(years)
}
