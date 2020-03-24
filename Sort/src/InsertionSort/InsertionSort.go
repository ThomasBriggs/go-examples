package main

import "fmt"

func main() {
	arr := []int{74, 47, 50, 89, 21, 37, 60, 19, 56, 18}
	sort(arr)
	fmt.Println(arr)
}

func sort(arr []int) {
	for i := 1; i < len(arr); i++ {
		j := i
		for j > 0 && arr[j] < arr[j-1] {
			swap(arr, j, j-1)
			j--
		}
	}
}

func swap(arr []int, index1 int, index2 int) {
	temp := arr[index1]
	arr[index1] = arr[index2]
	arr[index2] = temp
}
