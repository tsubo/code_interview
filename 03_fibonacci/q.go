package main

import "fmt"

/*
 * 以下はフィボナッチ数を求める関数です
 *   フィボナッチ数列： 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, ...
 *     F(0) = 0
 *     F(1) = 1
 *     F(n) = F(n-1) + F(n-2)
 *     ※ n >= 2
 *
 * このコードに問題はありますか？（ヒント：パフォーマンス）
 * どのように改善できますか？
 */
func fibonacci(n int) int {
	if n <= 1 {
		return n
	}

	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
	fmt.Println(fibonacci(10))
}
