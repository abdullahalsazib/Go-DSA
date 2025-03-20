package main

import (
	"fmt"
)

func main() {
	arr := [3]string{"go", "jack", "sparrow"}
	fmt.Print(arr)

	mySlice := make([]int, 10)

	fmt.Println(mySlice)

	mySlice = append(mySlice, 12)
	mySlice = append(mySlice, 100)
	fmt.Println(mySlice)
}
