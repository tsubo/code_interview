package main

import "fmt"

/*
 * 以下は int の配列 nums と int の target が与えられたとき
 * ２つの値を足すと target になるような nums の要素のインデックスを返す関数です
 *   例１：nums = [2, 7, 11, 15], target = 9
 *        => 2 + 7 = 9 なので、[0, 1] を返す
 *   例２：nums = [3, 2, 4], target = 6
 *        => 2 + 4 = 6 なので、[1, 2] を返す
 *
 * このコードに問題はありますか？（ヒント：パフォーマンス）
 * どのように改善できますか？
 */
func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[j] == target-nums[i] {
				return []int{i, j}
			}
		}
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
