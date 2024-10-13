package main

import "fmt"

func main() {
	happy := make(map[int]string)
	happy[1] = "a"
	happy[2] = "b"
	happy[3] = "c"
	happy[4] = "d"
	happy[5] = "e"
	happy[6] = "f"
	happy[7] = "g"
	happy[8] = "h"
	happy[9] = "i"
	happy[10] = "j"
	Assignment(happy)
}

func Assignment(myapp map[int]string) {
	for i, _ := range myapp {
		//delete(myapp, i)
		fmt.Println(i)

	}
	fmt.Println(myapp)
}
