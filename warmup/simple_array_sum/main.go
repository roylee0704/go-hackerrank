package main

import (
	"fmt"
	"os"
)

func main() {
	var n, sum, x int
	fmt.Fscan(os.Stdin, &n)
	for i := 0; i < n; i++ {

		count, _ := fmt.Fscan(os.Stdin, &x)
		if count < 0 {
			break
		}
		sum += x
	}
	fmt.Println(sum)
}
