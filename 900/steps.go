//https://codeforces.com/problemset/problem/580/A
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var (
		n, length, longest, last, now int
		err                           error
	)
	scanner := bufio.NewScanner(os.Stdin)
	longest = 1

	fmt.Scanf("%d\n", &n)
	scanner.Split(bufio.ScanWords)

	for i := 0; i < n; i++ {
		if !scanner.Scan() {
			break
		}

		last = now
		now, err = strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println(err.Error())
		}

		if now >= last {
			length++
			if length > longest {
				longest = length
			}
		} else {
			length = 1
		}
	}

	fmt.Println(longest)
}
