package main

import (
	"fmt"
	"leetcode_golang/algorithms"
)

func main() {
	arg := [][]byte{
		{1, 1, 0, 0},
		{1, 1, 0, 0},
		{0, 0, 1, 0},
		{0, 0, 0, 1},
	}
	result := algorithms.NumberOfIslands(arg)
	fmt.Println(result)
	fmt.Println(arg)
}
