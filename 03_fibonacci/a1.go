package main

import "fmt"

/*
 * 問題点：計算量がO(2^n)となるため、パフォーマンスが悪い
 * 改善方法：再帰をループに変えて計算量をO(n)にする
 */

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}

	fib := make([]int, n+1)
	fib[0], fib[1] = 0, 1

	for i := 2; i <= n; i++ {
		fib[i] = fib[i-1] + fib[i-2]
	}

	return fib[n]
}

func main() {
	fmt.Println(fibonacci(10))
}
