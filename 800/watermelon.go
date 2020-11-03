//https://codeforces.com/problemset/problem/4/A
package main

import (
	"fmt"
)

func main() {
	var weight int
	fmt.Scanf("%d", &weight)

	if weight < 3 {
		fmt.Println("NO")
		return
	}

	if weight%2 == 0 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
