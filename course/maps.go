package main

import "fmt"

func main() {
	LearnMaps()
}

func LearnMaps() {
	mappy := make(map[int]string)
	mappy[1] = "One"
	mappy[2] = "Two"
	fmt.Println(mappy)

	mappy2 := map[string]int{"One": 1, "Two": 2}
	fmt.Println(mappy2["One"])

}
