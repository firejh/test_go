package main

import (
	"fmt"
)

type Data struct {
	a 	int
	b 	int
	c 	string
}

func main() {
	var data Data
	data.a = 0
	data.b = 1
	data.c = "0,1"

	test := data
	test.a = 1
	test.b = 2
	test.c = "1,2"

	fmt.Println(data)
	fmt.Println(test)

}
