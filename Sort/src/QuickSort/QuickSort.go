package main

import "fmt"

func main() {
	arr := []int{}
	sort(arr[:], 0, 9)
	fmt.Println(arr)
}

func swap(arr []int, index1 int, index2 int) {
	temp := arr[index1]
	arr[index1] = arr[index2]
	arr[index2] = temp
}

func partition(arr []int, start int, end int) int {
	pivot := arr[end]
	wall := start

	for i := start; i <= end-1; i++ {
		if arr[i] < pivot {
			swap(arr, i, wall)
			wall++
		}
	}
	swap(arr, wall, end)
	return wall
}

func sort(arr []int, start int, end int) {
	if start < end {
		pivot := partition(arr, start, end)
		sort(arr, start, pivot-1)
		sort(arr, pivot+1, end)
	}
}
