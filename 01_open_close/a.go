package main

import "fmt"

/*
 * 問題点：会員ランクが追加された場合、既存のコードに手を加える必要がある
 * 改善方法：新しいコードを追加するだけで、既存のコードを変更することなく拡張できるようにする（オープン・クローズドの原則）
 */

type Member interface {
	GetPoint(price int) int
}

type GoldMember struct{}

func (m *GoldMember) GetPoint(price int) int {
	return int(float64(price) * 0.05)
}

type SilverMember struct{}

func (m *SilverMember) GetPoint(price int) int {
	return int(float64(price) * 0.04)
}

type NormalMember struct{}

func (m *NormalMember) GetPoint(price int) int {
	return int(float64(price) * 0.01)
}

func main() {
	var m Member = &GoldMember{}
	fmt.Println(m.GetPoint(100))
}
