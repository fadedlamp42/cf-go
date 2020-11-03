//https://codeforces.com/problemset/problem/236/A
package main

import "fmt"

func main() {
	chars := map[rune]bool{}
	var word string
	fmt.Scanf("%s\n", &word)
	for _, c := range word {
		chars[c] = true
	}

	if len(chars)%2 == 0 {
		fmt.Println("CHAT WITH HER!")
	} else {
		fmt.Println("IGNORE HIM!")
	}
}
