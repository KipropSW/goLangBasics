package main

import (
	"fmt"
	"reflect"
)

func main() {
	val1 := 4
	val2 := 3
	booleane, value := Add2(&val1, &val2)
	fmt.Println(booleane, value)
	fmt.Println("Hello World")
	fmt.Println('A')

	var boolean bool = true
	boolean2 := true

	fmt.Println(boolean2)
	fmt.Println(boolean)
	fmt.Println(reflect.TypeOf(boolean2))

	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
		if i == 4 {
			break
		}
	}
}

func hello() {
	fmt.Println("Hello World")
}

func Add(x int, y int) (bool, int) {
	fmt.Println(true, x+y)
	return true, x + y
}

func Add2(x *int, y *int) (bool, int) {
	*x = 9
	fmt.Println(true, *x+*y)
	return true, *x + *y
}
