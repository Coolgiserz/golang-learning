package main

import (
	"encoding/json"
	"fmt"
)

//需要注意首字母大小写，决定着字段能否被外部包（包括json.Marsgal)访问
type Order struct { //可以为结构体打上json tag，tag会被json.Marshal看到
	ID         string       `json:"id"`
	Item       *[]OrderItem `json:"item"`
	Quantity   int          `json:"quantity"`
	TotalPrice string       `json:"total_price"`
}

type OrderItem struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

//指针、指针数组、
func main() {
	o := Order{
		ID: "no1",
		Item: &[]OrderItem{
			{
				Name:  "book",
				Price: 4.2,
			}, {
				Name:  "course",
				Price: 94.2,
			},
		},
		Quantity: 2,
	}
	fmt.Printf("%+v\n", o)    //%+v 打印结构体
	b, err := json.Marshal(o) //将结构体序列化成可以在网络上传输的字节。（e.g.json格式)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", b)

	//对json字符串整理到struct类型中
	s := `{"id":"no1","item":{"id":"","name":"book","Price":4.2},"quantity":2}`
	var result Order
	json.Unmarshal([]byte(s), &result)
	fmt.Printf("%+v\n", result) //%+v 打印结构体
	fmt.Println(result.Item)

}
