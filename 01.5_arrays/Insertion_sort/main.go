package main

import "fmt"

func InsertionSortAlgo(arr []int) {
	n := len(arr)

	for i := 0; i < n; i++ {
		key := arr[i]
		j := i - 1

		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

func main() {
	arr := []int{64, 25, 12, 22, 11}
	fmt.Println("arr before sorted: ", arr)
	InsertionSortAlgo(arr)
	fmt.Println("arr after sorted: ", arr)

}
