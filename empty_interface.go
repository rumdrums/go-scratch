package main

import "fmt"

type Barf struct {
	Barfer int
}

func Passer(v Barf) {
	fmt.Println(v.Barfer)
}

func main() {
	var v Barf
	v.Barfer = 10
	Passer(v)
}
