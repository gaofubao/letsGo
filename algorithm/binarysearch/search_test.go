package binarysearch

import (
	"fmt"
	"testing"
)
// 1 2 3 4 5
func Search1(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + ((high-low)>>1)
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

func TestSearch1(t *testing.T)  {
	nums := []int{1, 2, 3, 4, 5}
	target := 2
	fmt.Println(Search1(nums, target))
}

// 1, 2, 3, 3, 4
// 查找第一个值等于给定值的元素
func Search2(nums []int, target int) int {
	low, high := 0, len(nums) - 1
	for low <= high {
		mid := low + ((high - low) >> 1)
		if nums[mid] > target {
			high = mid - 1
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			if mid == 0 || nums[mid-1] != target {
				return mid
			} else {
				high = mid - 1
			}
		}
	}
	return -1
}

func TestSearch2(t *testing.T)  {
	nums := []int{1, 2, 3, 3, 4}
	target := 3
	fmt.Println(Search2(nums, target))
}


