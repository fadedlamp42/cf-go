//https://codeforces.com/problemset/problem/118/A
package main

import (
	"fmt"
	"strings"
)

func main() {
	var in string
	fmt.Scanf("%s", &in)

	in = strings.ToLower(in)

	r := strings.NewReplacer(
		"a", "",
		"e", "",
		"i", "",
		"o", "",
		"u", "",
		"y", "",
	)
	in = r.Replace(in)

	for _, c := range in {
		fmt.Printf(".%c", c)
	}
}
