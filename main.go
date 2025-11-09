package main

import "fmt"

// MultiplyBy3 计算并返回数字乘以3的结果
func MultiplyBy3(n float64) float64 {
	return n * 3
}

// MultiplyBy3Int 计算并返回整数乘以3的结果
func MultiplyBy3Int(n int) int {
	return n * 3
}

// Hello returns the classic greeting string with multiply by 3 result.
func Hello() string {
	// 添加2乘以3的结果 (2*3=6)
	return fmt.Sprintf("hi zhangjiong %d", MultiplyBy3Int(2))
}

func main() {
	fmt.Println(Hello())
}
