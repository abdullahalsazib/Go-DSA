package main

import "fmt"

func BubbleSortAlgo(arr []int) {
	n := len(arr) // length of array

	for i := 0; i < n-1; i++ {
		swaped := false
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				// swaped element
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swaped = true
			}
		}
		if !swaped {
			break
		}
	}
}

func main() {
	arr := []int{64, 25, 12, 22, 11}
	fmt.Println("array is: ", arr)

	BubbleSortAlgo(arr)

	fmt.Println("sort the array using bubble sort algo: ", arr)

}
