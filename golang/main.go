package main

import (
	"leetcode_golang/algorithms"
)

func main() {
	arg := []int{-1, 0, 0, 3, 3, 3, 0, 0, 0}
	arg2 := []int{1, 2, 2}

	algorithms.Merge(arg, 6, arg2, 3)
}
