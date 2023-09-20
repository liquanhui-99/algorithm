package algorithm

// subSets 求指定切片所有可能的子集，包括空切片
func subSets(nums []int) [][]int {
	res := [][]int{}
	var fn func([]int, int)
	fn = func(arr []int, index int) {
		// terminator
		if index == len(nums) {
			res = append(res, append([]int{}, arr...))
			return
		}

		// process current logic: add index or not
		fn(arr, index+1)
		arr = append(arr, nums[index])
		fn(arr, index+1)
		arr = arr[:len(arr)-1]
	}
	fn([]int{}, 0)
	return res
}
