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

	// for _, test := range testcases {
	// 	if actual := lengthOfLongestSubstring(test.str); actual != test.length {
	// 		t.Errorf("Error occurred: lengthOfLongestSubstring(%s): got %d, expected %d", test.str, actual, test.length)
	// 	}
	// }
}

//性能测试
func BenchmarkSubStr(b *testing.B) {
	testStr := "军敌军还有五秒到达战场"
	ans := 10
	b.Logf("len(testStr)=%d", len(testStr))
	b.ResetTimer() //重置计时起点
	for i := 0; i < b.N; i++ {
		if actual := lengthOfLongestSubstring1(testStr); actual != ans {
			if actual := lengthOfLongestSubstring1(testStr); actual != ans {
				b.Errorf("Error occurred: lengthOfLongestSubstring1(%s): got %d, expected %d", testStr, actual, ans)
			}
		}

	}
}
