//https://codeforces.com/problemset/problem/133/A
package main

import (
	"fmt"
	"strings"
)

func main() {
	var line string
	fmt.Scanln(&line)
	if strings.ContainsAny(line, "HQ9") {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
