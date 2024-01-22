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

	arg4 := algorithms.TreeNode{
		Val:   20,
		Left:  nil,
		Right: nil,
	}

	result := algorithms.IsSameTree(&arg, &arg4)
	// fmt.Println(arg)
	fmt.Println(result)
}
