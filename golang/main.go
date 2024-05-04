package main

import (
	"fmt"
	"leetcode_golang/algorithms"
)

func main() {
	arg := []int{-1, 0, 0, 3, 3, 3, 3, 3}
	result := algorithms.MajorityElement(arg)
	fmt.Println(result)
}
