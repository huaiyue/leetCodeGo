package main

import (
	"fmt"
	"leetcodeGo/leetCode/easy"
)

func main()  {
	// 001
	fmt.Println("pppdddpppp")
	nums := []int{1, 3, 4, 6, 7, 10}
	target := 17
	newArr := easy.TwoSum(nums, target)
	fmt.Println(newArr)
}
