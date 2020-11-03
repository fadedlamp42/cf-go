//https://codeforces.com/problemset/problem/112/A
package main

import (
	"fmt"
	"strings"
)

func main() {
	var a, b string
	fmt.Scanf("%s\n%s", &a, &b)
	a = strings.ToLower(a)
	b = strings.ToLower(b)

	fmt.Println(strings.Compare(a, b))
}
