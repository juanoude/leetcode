package main

import (
	"fmt"
	"leetcode_golang/algorithms"
)

func main() {
	arg := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	result := algorithms.RemoveDuplicates(arg)
	fmt.Println(result)
	fmt.Println(arg)
}
