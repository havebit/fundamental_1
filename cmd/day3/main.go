package main

import (
	"fmt"
	_ "hello/cmd/day3/effect"
)

func init() {
	fmt.Println("initial stage")
}

func init() {
	fmt.Println("before main")
}

func main() {
	fmt.Println("Hello, world")
}
