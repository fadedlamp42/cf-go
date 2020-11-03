//https://codeforces.com/problemset/problem/281/A
package main

import (
	"fmt"
	"strings"
)

func main() {
	var word string
	fmt.Scanf("%s", &word)
	wordNew := fmt.Sprintf("%c%s", strings.ToUpper(word)[0], word[1:])
	fmt.Println(wordNew)
}
