package sort

import (
	"fmt"
	"testing"
)

// 归并排序
func MergeSort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}
	mid := len(nums) / 2
	right := nums[0:mid]
	left := nums[mid:]
	return Merge(MergeSort(right), MergeSort(left))
}

func Merge(right []int, left []int) []int {
	var result []int
	i, j := 0, 0
	for i < len(right) && j < len(left) {
		if right[i] < left[j] {
			result = append(result, right[i])
			i++
		} else {
			result = append(result, left[j])
			j++
		}
	}
	if i != len(right) {
		result = append(result, right[i])
	}
	if j != len(left) {
		result = append(result, left[j])
	}
	return result
}

func TestMerge(t *testing.T)  {
	nums := []int{8, 3, 10, 2, 7, 6, 9, 12}
	fmt.Println(MergeSort(nums))
}
