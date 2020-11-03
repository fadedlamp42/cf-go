//https://codeforces.com/problemset/problem/96/A
package main

import (
	"fmt"
	"strings"
)

func main() {
	var team string
	fmt.Scanf("%s", &team)
	if strings.Contains(team, "1111111") ||
		strings.Contains(team, "0000000") {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
