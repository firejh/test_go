package main

import (
	"fmt"
	"math"
)

func main() {
	var n uint64
	n = -1 ^ -1 << 64
	fmt.Println(n)

	fmt.Println(uint64(math.MaxUint64))

	te := uint64(1 << 64 - 1)

	fmt.Println(te)

	var(
		t1 uint32
		t2 uint32
		t3 uint32
	)
	t1 = 999
	t2 = 99

	t3 = (t1 << 8) | t2

	fmt.Println(t3)

	ne := uint64(476771499054874725)
	fmt.Println("get:")
	fmt.Println(ne)

	t := 9
	ss := fmt.Sprintf("%010d", t)

	fmt.Println(ss)
}
