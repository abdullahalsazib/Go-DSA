package main

import "fmt"

func partition(arr []int, start int, end int) int {
	pivot := arr[end]
	i := start - 1

	for j := start; j <= end-1; j++ {
		if arr[j] < pivot {
			i++
			temp := arr[i]
			arr[i] = arr[j]
			arr[j] = temp
		}
	}
	i++
	temp := arr[i]
	arr[i] = arr[end]
	arr[end] = temp
	return i
}

func QuickSortAlgo(arr []int, start int, end int) {
	if end <= start {
		return
	} // base cach

	pivot := partition(arr, start, end)
	QuickSortAlgo(arr, start, pivot-1)
	QuickSortAlgo(arr, pivot+1, end)
}
func main() {
	arr := []int{64, 25, 12, 22, 11}

	QuickSortAlgo(arr, 0, len(arr)-1)
	fmt.Println("Array is: ", arr)

}
