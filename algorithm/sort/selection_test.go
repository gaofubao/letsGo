package sort

import (
	"fmt"
	"testing"
)

// 选择排序，从非排序区中找出最小值放到排序区后面
func SelectionSort(nums []int) []int {
	for i := 0; i < len(nums) - 1; i++ {
		minIndex := i
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[minIndex] {
				minIndex = j
			}
		}
		nums[i], nums[minIndex] = nums[minIndex], nums[i]
	}
	return nums
}

func TestSelection(t *testing.T)  {
	nums := []int{8, 3, 10, 2, 7, 6, 9, 12}
	fmt.Println(SelectionSort(nums))
}

