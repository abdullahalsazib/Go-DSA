package main

import "fmt"

func FindMax(arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	max := arr[0]
	for _, num := range arr {
		if num > max {
			max = num
		}
	}
	return max
}

func main() {
	arr := [4]int{2, 31, 2, 44}

	maxElement := FindMax(arr[:])
	fmt.Println("Max element of array is: ", maxElement)
}
