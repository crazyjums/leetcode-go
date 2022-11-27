package main

import (
	"fmt"
	"leetcode-go/simple"

	//_ "leetcode-go/hard"
	_ "leetcode-go/middle"
	_ "leetcode-go/simple"
)

func main() {
	arr := []int{5, 1, 13, 15, 17, -7}
	//fmt.Println(simple.TwoSum(arr, 16))
	//fmt.Println(simple.TwoSum2(arr, 16))

	//arr2 := simple.BubbleSort(arr, true)
	//fmt.Println(arr2)
	//
	//arr2 = simple.BubbleSort(arr, false)
	//fmt.Println(arr2)
	//
	//simple.BubbleSortPtr(&arr, len(arr), true)
	//fmt.Println(arr2)

	//simple.BubbleSortPtr(&arr, len(arr), false)
	//fmt.Println(arr2)

	//arr2 := simple.SelectionSort(arr, true)
	//fmt.Println(arr2)

	simple.SelectionSortPtr(&arr, len(arr), false)
	fmt.Println(arr)

	simple.SelectionSortPtr(&arr, len(arr), true)
	fmt.Println(arr)
}
