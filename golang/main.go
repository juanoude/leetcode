package main

import (
	"fmt"
	"leetcode_golang/algorithms"
)

func main() {
	arg2 := algorithms.TreeNode{
		Val:   10,
		Left:  nil,
		Right: nil,
	}

	arg3 := algorithms.TreeNode{
		Val:   20,
		Left:  nil,
		Right: nil,
	}

	arg := algorithms.TreeNode{
		Val:   15,
		Left:  &arg2,
		Right: &arg3,
	}

	result := algorithms.InvertTree(&arg)
	// fmt.Println(arg)
	fmt.Printf("result: %+v", result.Left.Val)
}
