package main

import "fmt"

func twoSum(nums []int, target int) []int {
	if nums == nil {
		return []int{}
	}

	size := len(nums)
	for i := 0; i < size; i++ {
		for j := i + 1; j < size; j++ {
			if nums[j]+nums[i] == target {
				return []int{i, j}
			}
		}
	}

	return []int{}
}

func twoSum2(nums []int, target int) []int {
	if nums == nil {
		return []int{}
	}

	midMap := make(map[int]int)
	for i, v := range nums {
		if j, ok := midMap[target-v]; ok {
			return []int{i, j}
		}
		midMap[v] = i
	}

	return []int{}
}

func main() {
	arr := []int{1, 13, 15, 17}
	fmt.Println(twoSum(arr, 16))
	fmt.Println(twoSum2(arr, 16))
}
