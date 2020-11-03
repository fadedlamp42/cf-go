//https://codeforces.com/problemset/problem/282/A
package main

import "fmt"

func main() {
	var x, n int
	var line string

	fmt.Scanf("%d\n", &n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%s\n", &line)
		if line[1] == '+' {
			x++
		} else {
			x--
		}
	}

	fmt.Println(x)
}
