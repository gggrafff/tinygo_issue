package main

import (
	"fmt"
)

type IA interface {
}

type A struct {
	Value float64
}

func main() {
	val1 := any(&A{Value: 1.618}).(IA)
	_ = val1
	fmt.Println("Hello, world")
}
