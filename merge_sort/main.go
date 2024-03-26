package main

import "fmt"

type Number interface {
	int | float64
}

func main() {
	nums := []int{12, 3, 7, 9, 14, 6, 11, 2}

	MergeSort(nums)
	fmt.Printf("nums: %v\n", nums)
}

func MergeSort[num Number](a []num) {
	if len(a) <= 1 {
		return
	}

	q := len(a) / 2

	l := make([]num, q)
	r := make([]num, len(a)-q)
	copy(l, a[:q])
	copy(r, a[q:])

	MergeSort(l)
	MergeSort(r)

	var i, j, k = 0, 0, 0

	for i < len(l) && j < len(r) {
		if l[i] <= r[j] {
			a[k] = l[i]
			i++
		} else {
			a[k] = r[j]
			j++
		}
		k++
	}

	for i < len(l) {
		a[k] = l[i]
		i++
		k++
	}

	for j < len(r) {
		a[k] = r[j]
		j++
		k++
	}
}
