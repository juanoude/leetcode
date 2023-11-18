package main

import (
	"fmt"
	"leetcode_golang/algorithms"
)

func main() {
	arg4 := algorithms.ListNode{
		Val:  4,
		Next: nil,
	}

	arg3 := algorithms.ListNode{
		Val:  3,
		Next: &arg4,
	}

	arg2 := algorithms.ListNode{
		Val:  2,
		Next: nil,
	}

	arg1 := algorithms.ListNode{
		Val:  1,
		Next: &arg2,
	}

	result := algorithms.MergeTwoLists(&arg1, &arg3)
	fmt.Println(result)
	fmt.Printf("%+v", result)
}
