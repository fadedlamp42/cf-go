//https://codeforces.com/problemset/problem/116/A
package main

import "fmt"

func main() {
	var n, max, on, off, current int
	fmt.Scanf("%d\n", &n)

	for i := 0; i < n; i++ {
		fmt.Scanf("%d %d\n", &off, &on)
		current += on - off
		if current > max {
			max = current
		}
	}

	fmt.Println(max)
}
