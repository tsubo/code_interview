package main

import "fmt"

/*
 * 問題点：計算量がO(2^n)となるため、パフォーマンスが悪い
 * 改善方法：メモ化を使用して計算量をO(n)にする
 */

var memo = map[int]int{}

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}

	if val, ok := memo[n]; ok {
		return val
	}

	memo[n] = fibonacci(n-1) + fibonacci(n-2)
	return memo[n]
}

func main() {
	fmt.Println(fibonacci(10))
}
