package main

import "fmt"

func main() {
	//定义Map
	m := map[string]string{
		"GCN":  "图卷积网络",
		"GAT":  "图注意力网络",
		"AGCN": "自适应图卷积网络",
		"HGCN": "双曲图卷积网络",
	}
	m2 := make(map[string]int) //empty map
	var m3 map[string]int      //nil

	fmt.Println(m, m2, m3)

	//遍历Map
	for k, v := range m {
		fmt.Println(k, v)
	}

	//获取Map中的元素;Key不存在时不会报错，而是返回Value类型的初始值
	if transformer, isExist := m["transformer"]; isExist {
		fmt.Println(transformer)
	} else {
		fmt.Println(isExist)
	}

	//删除元素
	delete(m, "GCN")
	fmt.Println(m, len(m))

}
