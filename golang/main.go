package main

import (
	"fmt"
	"leetcode_golang/algorithms"
)

func main() {
	arg := []int{-1, 1}
	result := algorithms.ThreeSum(arg)
	// fmt.Println(arg)
	fmt.Println(result)
}
