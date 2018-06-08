package main

import "fmt"

func main() {
	var test []int
	test = append(test, 1)
	test = append(test, 2)

	fmt.Println(test[0])
	fmt.Println(test[1])
	fmt.Println(test[2])

	fmt.Println("stop")

}
