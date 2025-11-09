package main

import "fmt"

// MultiplyBy5 计算并返回数字乘以5的结果
func MultiplyBy5(n float64) float64 {
	return n * 5
}

// MultiplyBy5Int 计算并返回整数乘以5的结果
func MultiplyBy5Int(n int) int {
	return n * 5
}

// Hello returns the classic greeting string with multiply by 5 result.
func Hello() string {
	// 添加2乘以5的结果 (2*5=10)
	return fmt.Sprintf("hi zhangjiong %d", MultiplyBy5Int(2))
}

func main() {
	fmt.Println(Hello())
}
