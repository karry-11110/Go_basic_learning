package main

import "fmt"

func main() {
	fmt.Println(getLeastNumbers([]int{0, 0, 1, 2, 4, 2, 2, 3, 1, 4}, 8))
}
func getLeastNumbers(arr []int, k int) []int {
	n := len(arr)
	if n == 0 || k <= 0 || n < k {
		return nil
	}
	var partition func(arr []int, begin, end int) int
	partition = func(arr []int, begin, end int) int {
		left, right := begin+1, end
		for left < right {
			if arr[left] > arr[begin] {
				arr[left], arr[right] = arr[right], arr[left]
				right--
			} else {
				left++
			}
		}
		if arr[left] >= arr[begin] {
			left--
		}
		arr[begin], arr[left] = arr[left], arr[begin]
		return left
	}
	index := partition(arr, 0, n-1)
	fmt.Println(index)
	for index != k-1 {
		if index > k-1 {
			index = partition(arr, s, index-1)
		} else {
			index = partition(arr, index+1, n-1)
		}
	}
	return arr[:k]
}
