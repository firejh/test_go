package main

import (
	"time"
	"fmt"
)

func main() {

	t := time.Now()
	layout := "20060102 15:04:05"

	ts := t.Format(layout)
	td, _ := time.Parse(layout, ts)

	fmt.Println(t.Unix())
	fmt.Println(ts)
	fmt.Println(td.Unix())
}
