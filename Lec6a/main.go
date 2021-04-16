package main

import "fmt"

func main() {
	example := []int{1, 2, 3, 4, 5}
	for _, v := range example {
		fmt.Printf("%d ", v)
	}
	fmt.Println()

	for i := 0; i < 5; i++ {
		fmt.Printf("%v ", itea())
	}
	fmt.Println()
}

var start = 0

func itea() int {
	start++
	return start
}
