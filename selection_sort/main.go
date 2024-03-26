package main

import "fmt"

type Number interface {
	int | float64
}

func SelectionSort[num Number](nums []num) {
	for i := 0; i < len(nums)-1; i++ {
		j := i + 1

		fmt.Println("for executed")
		for j < len(nums) && nums[j] < nums[i] {
			fmt.Println("While executed for", i)
			nums[i], nums[j] = nums[j], nums[i]
			j++
		}
	}
}

func main() {
	nums := []int{12, 3, 7, 9, 14, 6, 11, 2}

	SelectionSort(nums)
	fmt.Printf("nums: %v\n", nums)
}
