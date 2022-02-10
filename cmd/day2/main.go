package main

import (
	"fmt"
)

func main() {
	var n int = 2

	switch {
	case n%2 == 0:
		fmt.Println("even")
	default:
		fmt.Println("odd")
	}
}

func emptyInterface() {
	var i interface{}

	fmt.Println(i == nil)

	i = 10
	fmt.Printf("%T %v\n", i, i)

	i = "ten"
	fmt.Printf("%T %v\n", i, i)

	if s, ok := i.(int); ok {
		fmt.Printf("%v is int\n", s)
	}
	if s, ok := i.(string); ok {
		fmt.Printf("%v is string\n", s)
	}
}

func deferDemo() {
	defer func() {
		fmt.Println("in defer func")
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()

	var err error
	fmt.Println(err.Error())
}
