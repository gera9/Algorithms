package main

import (
	"fmt"
)

type Number interface {
	int | float64
}

func InsertionSort[num Number](nums []num) {
	for i := 1; i < len(nums); i++ {
		j := i - 1

		for j >= 0 && nums[j] > nums[j+1] {
			nums[j], nums[j+1] = nums[j+1], nums[j]
			j--
		}
	}
}

func InsertionSortBook[num Number](nums []num) {
	for i := 1; i < len(nums); i++ {
		key := nums[i]
		j := i - 1

		for j >= 0 && nums[j] > key {
			nums[j+1] = nums[j]
			j--
		}
		nums[j+1] = key
	}
}

func InsertionSortReversed[num Number](nums []num) {
	for i := len(nums) - 2; i >= 0; i-- {
		key := nums[i]
		j := i + 1

		for j < len(nums) && nums[j] < key {
			nums[j-1] = nums[j]
			j++
		}
		nums[j-1] = key
	}
}

func RecursiveInsertionSort[num Number](nums []num) {
	if len(nums) == 0 {
		return
	}

	RecursiveInsertionSort(nums[:len(nums)-1])
	fmt.Printf("nums: %v\n", nums)

	i := len(nums) - 1
	j := i - 1

	for (j > -1 && i > 0) && nums[j] > nums[j+1] {
		fmt.Printf("%v ", j)
		nums[j], nums[j+1] = nums[j+1], nums[j]
		j--
	}
	fmt.Println()
}

func BinarySearch[num Number](nums []num, target, low, high int) int {
	if low >= high {
		return high
	}

	mid := low + (high-low)/2

	if num(target) == nums[mid] {
		return mid
	} else if num(target) < nums[mid] {
		return BinarySearch(nums, target, low, mid-1)
	} else {
		return BinarySearch(nums, target, mid+1, high)
	}
}

func InsertionSortBinarySearch[num Number](nums []num) []num {
	for i := 1; i < len(nums); i++ {
		key := nums[i]

		// Find the index where the key should be inserted
		pos := BinarySearch(nums, int(key), 0, i)

		for j := i; j <= pos; j-- {
			nums[j] = nums[j-1]
		}

		nums[pos] = key
	}

	return nums
}

func main() {
	nums := []int{3, 2, 1}
	fmt.Printf("%v\n", InsertionSortBinarySearch(nums))
	fmt.Printf("nums: %v\n", nums)
}
