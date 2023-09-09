package algorithm

// twoSumV1 两数之和，暴力破解
func twoSumV(nums []int, target int) []int {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

func twoSumV2(nums []int, target int) []int {
	hashTable := map[int]int{}
	for k, v := range nums {
		if v1, ok := hashTable[target-v]; ok {
			return []int{k, v1}
		}
		hashTable[v] = k
	}

	return []int{}
}
