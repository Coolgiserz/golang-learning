package main

import (
	"fmt"
	"regexp"
)

// const text = "Googlo's email is googalo@gmail.com"
// const text = "Googlo's email is googalo@gmail.com@eas.com" //bad case
const text = `
Googlo's email is googalo@gmail.com@eas.com
Googlo's email is sadf@gmail.co
Googlo's email is rfeda@gmail.com.ds@eas.com
` //bad case

//正则表达式使用
func main() {
	// re := regexp.MustCompile("googlo@gmail.com")
	// re := regexp.MustCompile(".+@.+\\..+")//wrong email pattern
	// re := regexp.MustCompile(`[a-zA-Z0-9]+@[a-zA-Z0-9]+\.[a-zA-Z]+`) //避免使用转义字符\
	re := regexp.MustCompile(`([a-zA-Z0-9]+)@([a-zA-Z0-9]+)(\.[a-zA-Z]+)+`) //避免使用转义字符\

	// 正则表达式提取：字符串

	// match := re.FindString(text)
	// match := re.FindAllString(text, -1)

	// 正则表达式提取：子匹配
	match := re.FindAllStringSubmatch(text, -1)

	fmt.Println(match)

}
