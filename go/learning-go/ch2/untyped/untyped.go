package main

import "fmt"

const untyped = 10

func main() {
	var i int = untyped
	var f float64 = untyped

	fmt.Println(i)
	fmt.Println(f)
}
