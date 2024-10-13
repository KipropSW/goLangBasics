package main

import "fmt"

func main() {
	LearnSlices()
}

func LearnSlices() {
	var slice []int

	slice = append(slice, 1)
	slice = append(slice, 2)
	fmt.Println(slice)

	sliceWithFixedSize := make([]int, 10)
	fmt.Println(sliceWithFixedSize)

	for _, value := range sliceWithFixedSize {
		fmt.Println(value)
	}
}
