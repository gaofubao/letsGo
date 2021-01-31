package sort

import (
	"fmt"
	"testing"
)

// 插入排序，已排序元素从后往前依次与新元素比较，如果大于新元素则往后移动一位，最后将新元素插入空位置中
func InsertionSort(nums []int) []int {
	for i := 1; i < len(nums); i++ {
		preIndex := i - 1
		current := nums[i]
		for preIndex >= 0 && nums[preIndex] > current {
			nums[preIndex+1] = nums[preIndex]
			preIndex--
		}
		nums[preIndex+1] = current
	}
	return nums
}

func TestInsertion(t *testing.T)  {
	nums := []int{8, 3, 10, 2, 7, 6, 9, 12}
	fmt.Println(InsertionSort(nums))
}
