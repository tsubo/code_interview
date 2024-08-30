package main

import "fmt"

/*
 * 以下は会員ランクに応じてポイントを計算するコードです。
 *
 * このコードに問題はありますか？（ヒント：拡張性）
 * どのように改善できますか？
 */

type Member struct {
	rank string
}

func (m *Member) GetPoint(price int) int {
	p := float64(price)

	if m.rank == "gold" {
		return int(p * 0.05)
	}
	if m.rank == "siler" {
		return int(p * 0.04)
	}
	return int(p * 0.01)
}

func main() {
	m := Member{"gold"}
	fmt.Println(m.GetPoint(100))
}
