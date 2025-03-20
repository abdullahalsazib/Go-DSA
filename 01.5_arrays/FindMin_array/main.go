package main

import "fmt"

func MinFinder(arr []int) int {

	if len(arr) == 0 {
		return 0
	}

	min := arr[0]
	for _, num := range arr {
		if num < min {
			min = num
		}
	}
	return min
}

func main() {

	arr := [3]int{1, 2, 3}

	find_element := MinFinder(arr[:])
	fmt.Println(" Min element of the array: ", find_element)

}
