package main

import (
	"fmt"
)

func main() {
	test := map[int] int {1:1, 2:2, 3:3}

	for k,_ := range test {
		test[k] = 0
	}

	for k,v := range test {
		fmt.Print(k, ":", v, "\n")
	}
}