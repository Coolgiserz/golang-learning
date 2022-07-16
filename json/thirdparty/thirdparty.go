package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	s := `{"data":[
		{
		 "synonym": "",
		 "weight": "0.800000",
		 "tag": "品牌",
		 "word": "三只松鼠"
		},
		{
		 "synonym": "",
		 "weight": "0.100000",
		 "tag": "普通词",
		 "word": " "
		},
		{
		 "synonym": "开心面具",
		 "weight": "1.000000",
		 "tag": "品类",
		 "word": "开心果"
		},
		{
		 "synonym": "",
		 "weight": "0.100000",
		 "tag": "普通词",
		 "word": " "
		},
		{
		 "synonym": "",
		 "weight": "1.000000",
		 "tag": "品类",
		 "word": "零食"
		},
		{
		 "synonym": "",
		 "weight": "1.000000",
		 "tag": "品类",
		 "word": "坚果"
		},
		{
		 "synonym": "",
		 "weight": "1.000000",
		 "tag": "品类",
		 "word": "炒货"
		},
		{
		 "synonym": "",
		 "weight": "0.600000",
		 "tag": "修饰",
		 "word": "无漂白"
		},
		{
		 "synonym": "",
		 "weight": "0.600000",
		 "tag": "修饰",
		 "word": "原味"
		},
		{
		 "synonym": "",
		 "weight": "0.100000",
		 "tag": "普通词",
		 "word": "健康"
		},
		{
		 "synonym": "",
		 "weight": "1.000000",
		 "tag": "品类",
		 "word": "食品"
		},
		{
		 "synonym": "",
		 "weight": "1.000000",
		 "tag": "品类",
		 "word": "小吃"
		}
	   ]}`

	//====方式一：解析到json
	// m := make(map[string]interface{}) //json
	// json.Unmarshal([]byte(s), &m)
	// // fmt.Printf("%+v\n", m["data"].([]interface{})[2])
	// fmt.Printf("%+v\n", m["data"].([]interface{})[2].(map[string]interface{})["word"]) //需要通过类型断言告知编译器m["data"]属于哪个类型

	//====方式二：解析到自定义结构体
	m := struct {
		Data []struct {
			Synonym string `json:"synonym"`
			Tag     string
			Weight  float64
			Word    string
		} `json:"data"`
	}{}

	json.Unmarshal([]byte(s), &m)
	fmt.Printf("%+v %+v\n", m.Data[2].Synonym, m.Data[2].Word)

	//TODO: 通过http向阿里云NLP服务器发送Post请求，对返回的结果进行Unmarshal处理；https://help.aliyun.com/document_detail/177235.html
	//refer to:
	// 1. https://help.aliyun.com/document_detail/176979.html
	// 2. https://github.com/aliyun/alibaba-cloud-sdk-go/blob/master/README-CN.md?spm=a2c4g.11186623.0.0.489e3c6bwX0rUG&file=README-CN.md
	// 3. https://github.com/wArrest/aliyun-nlp-sdk/blob/master/client.go
}
