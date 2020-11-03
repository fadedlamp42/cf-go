//https://codeforces.com/problemset/problem/339/A
package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	var s string
	fmt.Scanln(&s)
	fields := strings.Split(s, "+")
	sort.Sort(sort.StringSlice(fields))
	fmt.Println(strings.Join(fields, "+"))
}
