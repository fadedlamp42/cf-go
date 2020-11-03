//https://codeforces.com/problemset/problem/71/A
package main

import (
	"fmt"
)

func main() {
	var n int
	var word string

	fmt.Scanf("%d\n", &n)

	for i := 0; i < n; i++ {
		fmt.Scanf("%s\n", &word)
		if len(word) > 10 {
			fmt.Printf("%c%d%c\n", word[0], len(word)-2, word[len(word)-1])
		} else {
			fmt.Println(word)
		}
	}
}
