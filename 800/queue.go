//https://codeforces.com/problemset/problem/266/B
package main

import (
	"fmt"
)

func move(lineP *[]byte) {
	line := *lineP

	for i := 0; i < len(line)-1; i++ {
		if line[i] == 'B' && line[i+1] == 'G' {
			line[i], line[i+1] = line[i+1], line[i]
			i++
		}
	}
}

func main() {
	var n, t int
	var line []byte
	fmt.Scanf("%d %d\n%s", &n, &t, &line)
	for i := 0; i < t; i++ {
		move(&line)
	}
	fmt.Printf("%s\n", line)
}
