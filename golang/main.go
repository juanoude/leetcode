package main

import (
	"fmt"
	"leetcode_golang/algorithms"
)

func main() {
	arg := []int{2, 3, 5}
	arg2 := [][]int{{4, 8}, {2, 8}}
	result := algorithms.MinimumSpaceWastedFromPacking(arg, arg2)
	fmt.Println(arg)
	fmt.Println(arg2)
	fmt.Println(result)
}
