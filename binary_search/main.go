package main

import "fmt"

type number interface {
	int | float64 | float32
}

func main() {
	arr := []int{1, 2, 3, 4}

	for _, v := range arr {
		index := RecursiveBinarySearch(arr, v, 0, len(arr)-1)

		fmt.Printf("%d: index: %v\n", v, index)
	}

}

func RecursiveBinarySearch[num number](nums []num, target num, low, high int) int {
	if high < low {
		return -1 // element not found
	}
	fmt.Printf("low: %v\n", low)
	fmt.Printf("high: %v\n", high)
	mid := low + (high-low)/2
	fmt.Printf("mid: %v\n", mid)
	if nums[mid] == target {
		return mid // element found
	}

	if nums[mid] > target {
		return RecursiveBinarySearch(nums, target, low, mid-1) // search left half
	}

	return RecursiveBinarySearch(nums, target, mid+1, high) // search right half
}
