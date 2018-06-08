package main

import (
	"fmt"
	"encoding/json"
)
// Product 商品信息
type Product struct {
	Name      string
	ProductID int64
	Number    int
	price     float64
	IsOnSale  bool
}
type Productun struct {
	Name      string
	ProductID int64
	Number    int
	price     float64
	IsOnSale  bool
	Test 	string
}

var test = "{\"beat\": { \"hostname\": \"localhost\"},\"type\": \"topic_go_test\"}"

type Test struct {
	Type	string 	`json:"type""`
}

func main() {
	p := Product{}
	p.Name = "Xiao mi 6"
	p.IsOnSale = true
	p.Number = 10000
	p.price = 2499.00
	p.ProductID = 1
	data, _ := json.Marshal(p)
	fmt.Println(string(data))

	var info = &Productun{}
	json.Unmarshal(data, info)

	newdata, _ := json.Marshal(info)
	fmt.Println(string(newdata))

	var t = &Test{}
	json.Unmarshal([]byte(test),t)
	fmt.Println(t.Type)
}
