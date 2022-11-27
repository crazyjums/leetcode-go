package main

import (
	"fmt"
	"leetcode-go/simple"

	//_ "leetcode-go/hard"
	_ "leetcode-go/middle"
	_ "leetcode-go/simple"
)

func main() {
	arr := []int{1, 13, 15, 17}
	fmt.Println(simple.TwoSum(arr, 16))
	fmt.Println(simple.TwoSum2(arr, 16))
}
