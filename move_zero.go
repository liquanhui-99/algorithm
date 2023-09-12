package algorithm

func MoveZeroes(nums []int) []int {
	left, right, n := 0, 0, len(nums)
	for right < n {
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
		right++
	}

	return nums
}

// 题目：移动所有零到数组(切片)的尾部，不能开辟新的内存空间，在原对象的基础上操作，尽量减少操作次数
// 基本解题思路：快慢指针
// 快慢指针在算法中很常见，定义左右指针，当右指针小于数组长度的时候不断的循环
// 如果右指针的值不等于0，就把右指针的值和左指针的值调换一下位置，左右指针都+1
// 如果右指针的值等于0，就只把右指针+1
// 依次遍历，交换元素位置
// 时间复杂度：O(n)
// 空间复杂度：O(1)
