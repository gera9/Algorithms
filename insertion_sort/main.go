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

func RecursiveInsertionSort(numbers []int) {
	if len(numbers) == 0 {
		return
	}

	RecursiveInsertionSort(numbers[:len(numbers)-1])

	i := len(numbers) - 1
	j := i - 1

	for (j > -1 && i > 0) && numbers[j] > numbers[j+1] {
		numbers[j], numbers[j+1] = numbers[j+1], numbers[j]
		j--
	}
}

func main() {
	numbers := []int{5, 2, 4, 6, 1, 3}
	RecursiveInsertionSort(numbers)
	fmt.Println(numbers)
}
