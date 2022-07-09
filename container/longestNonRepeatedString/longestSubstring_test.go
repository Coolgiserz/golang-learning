package main

import "testing"

func TestLongestSubstring(t *testing.T) {
	testcases := []struct {
		str    string
		length int
	}{
		//边界测试
		{"", 0},
		{"a", 1},
		{"aaaa", 1},

		//常规功能测试
		{"aaaba", 2},
		{"aaaasasdfba", 5},
		{"一次性医用外科口罩", 9},
		{"abcdefg", 7},
		{"0123456", 7},

		// //错误用例
		// {"aaaasasdfba", 51},
		// {"一次性医用外科口罩", 10},

	}

	for _, test := range testcases {
		if actual := lengthOfLongestSubstring1(test.str); actual != test.length {
			t.Errorf("Error occurred: lengthOfLongestSubstring1(%s): got %d, expected %d", test.str, actual, test.length)
		}
	}
}
