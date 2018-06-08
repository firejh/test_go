package main

import (
	"fmt"
	"time"
)

func main() {
	var t map[string] []interface{}
	t = make(map[string] []interface{})

	for {
		t["test"] = append(t["test"], "1")
		t["test"] = append(t["test"], "2")
		t["test"] = append(t["test"], "3")
		t["test"] = append(t["test"], "4")
		t["test"] = append(t["test"], "5")


		fmt.Printf("len=%d\n", len(t["test"]))

		for k, v := range t {
			t[k] = t[k][:0]
			fmt.Printf("len=%v\n", len(v))
		}
		//t["test"] = t["test"][:0]
		fmt.Printf("len=%d\n", len(t["test"]))
		time.Sleep(time.Second * 1)
	}


}
