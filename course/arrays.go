package main

import "fmt"

func main() {
	Array()
}

func Array() {
	arr := [3]int{1, 2, 3}
	fmt.Println(arr)

	arr2 := [2]int{}
	arr2[0] = 1
	arr2[1] = 2
	fmt.Println(arr2)

	for _, v := range arr {
		fmt.Println(v)
	}
}
