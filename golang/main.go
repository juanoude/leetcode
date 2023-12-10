package main

import (
	"fmt"
	"leetcode_golang/algorithms"
)

func main() {
	arg := []int{3, 2, 2, 3}
	result := algorithms.RemoveElement(arg, 3)
	fmt.Println(result)
	fmt.Println(arg)
}
