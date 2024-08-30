package main

import "fmt"

/*
 * 問題点：ループが入れ子になっていて、計算量がO(2^n)となるため、パフォーマンスが悪い
 * 改善方法：ハッシュを使いループを１重に変えて計算量をO(n)にする
 */

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, num := range nums {
		x := target - num
		if j, ok := m[x]; ok {
			return []int{i, j}
		}
		m[num] = i
	}
	return nil
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 18
	result := twoSum(nums, target)
	if result == nil {
		fmt.Println("No solution found")
	} else {
		fmt.Println("Indeces: ", result[0], result[1])
	}
}
