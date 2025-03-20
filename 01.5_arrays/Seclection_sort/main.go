package main

import "fmt"

func SeclictionSortAlgo(arr []int, n int) {
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if arr[i] > arr[j] {
				// swaped
				temp := arr[i]
				arr[i] = arr[j]
				arr[j] = temp
			}
		}
	}
}

func main() {
	fmt.Println("hello, world")
	arr := []int{64, 25, 12, 22, 11}
	fmt.Println("array is: ", arr)
	SeclictionSortAlgo(arr, len(arr))
	fmt.Println("sorted array is : ", arr)
}
