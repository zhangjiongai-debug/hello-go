package main

import "fmt"

// Square 计算并返回数字的平方
func Square(n float64) float64 {
    return n * n
}

// SquareInt 计算并返回整数的平方
func SquareInt(n int) int {
    return n * n
}

// Hello returns the classic greeting string with square result.
func Hello() string {
    // 添加2的平方结果 (2*2=4)
    return fmt.Sprintf("hi zhangjiong %d", SquareInt(2))
}

func main() {
    fmt.Println(Hello())
}

