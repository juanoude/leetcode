package main

import (
	"fmt"
	"leetcode_golang/algorithms"
)

func main() {
	arg := []int{1, 3, 5, 6, 8, 9}
	result := algorithms.ProductExceptSelf(arg)
	fmt.Println(result)
}
