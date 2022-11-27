package simple

func TwoSum(nums []int, target int) []int {
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

func TwoSum2(nums []int, target int) []int {
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
