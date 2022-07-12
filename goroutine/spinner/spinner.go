package main

import (
	"fmt"
	"time"
)

func main() {
	go spinner((100 * time.Millisecond))
	const n = 45
	fmt.Printf("\rFib(%d) = %d\n", fib2(n), fib2(n))
	// for i := 0; i < 10; i++ {
	// 	// fmt.Println(fib2(i + 1))
	// 	// fmt.Println(fib1(i ))
	// 	fmt.Println(fib3(i))

	// }
}

// 斐波那契数列: 实现1： 循环
func fib1(n int) int {
	a, b := 0, 1
	if n == 0 {
		return b
	}
	for i := 0; i < n; i++ {
		b, a = a+b, b
	}
	return b
}

// 斐波那契数列: 实现2: 递归
func fib2(n int) int {

	if n == 0 {
		return 0
	} else if n == 1 || n == 2 {
		return 1
	}
	return fib2(n-1) + fib2(n-2)
}

//斐波那契数列： 实现3: 动态规划
func fib3(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}
	dp := []int{0, 1}
	// dp[0] = 0
	// dp[1] = 1
	for i := 2; i < n; i++ {
		dp = append(dp, dp[i-1]+dp[i-2])
	}
	return dp[n-1]

}

//spinner
func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r) // \r表示归位
			time.Sleep(delay)
		}
	}
}
