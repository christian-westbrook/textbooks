package main

import "fmt"

func main() {
	var b byte = 255
	var smallI int32 = 2147483647
	var bigI uint64 = 18446744073709551615

	fmt.Println("Before")
	fmt.Println(b)
	fmt.Println(smallI)
	fmt.Println(bigI)
	fmt.Println()

	b += 1
	smallI += 1
	bigI += 1

	fmt.Println("After")
	fmt.Println(b)
	fmt.Println(smallI)
	fmt.Println(bigI)
}
