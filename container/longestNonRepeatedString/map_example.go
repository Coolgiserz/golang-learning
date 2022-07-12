package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	//处理特殊情况:s为空串时，返回0
	if len(s) == 0 {
		return 0
	}
	maxLength := 1
	start := 0 //最长不重复子串起时索引
	lastOccurred := make(map[byte]int)

	for i, ch := range ([]byte)(s) {
		pos, hasOccured := lastOccurred[ch]
		if hasOccured && pos >= start {
			start = pos + 1
		} else {
			//ch之前没出现过
			if i-start+1 > maxLength {
				maxLength = i - start + 1
			}
		}
		lastOccurred[ch] = i
	}
	return maxLength
}
func lengthOfLongestSubstring1(s string) int {
	//处理特殊情况:s为空串时，返回0
	if len(s) == 0 {
		return 0
	}
	maxLength := 1
	start := 0 //最长不重复子串起时索引
	lastOccurred := make(map[rune]int)

	for i, ch := range []rune(s) {
		pos, hasOccured := lastOccurred[ch]
		if hasOccured && pos >= start {
			start = pos + 1
		} else {
			//ch之前没出现过
			if i-start+1 > maxLength {
				maxLength = i - start + 1
			}
		}
		lastOccurred[ch] = i
	}
	return maxLength
}

func main() {
	fmt.Println(lengthOfLongestSubstring("saadfeffa"))
	fmt.Println(lengthOfLongestSubstring("bbbbb"))
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("吃葡萄不吐葡萄皮"))  //出现问题
	fmt.Println(lengthOfLongestSubstring1("吃葡萄不吐葡萄皮")) //解决问题

	fmt.Println(lengthOfLongestSubstring("军敌军还有五秒到达战场")) //出现问题

}
