package sort

import (
	"fmt"
	"testing"
)

// 冒泡排序，相邻元素两两比较
func BubbleSort(nums []int) []int {
	for i := 0; i < len(nums) -1; i++ {
		for j := 0; j < len(nums) - 1 - i; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
	return nums
}

func TestBubble(t *testing.T)  {
	nums := []int{8, 3, 10, 2, 7, 6, 9, 12}
	fmt.Println(BubbleSort(nums))
}

