package algorithm

// maxWater 盛最多的水，求最大的体积
func maxWater(height []int) int {
	left, right := 0, len(height)-1
	area := 0
	for left < right {
		newArea := (right - left) * min(height[left], height[right])
		area = max(area, newArea)

		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return area
}

// min 获取最小值
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// max获取最大值
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// 题解：盛最多的水右两种解法
// 解法一：双层循环，比较简单就不解释了，这种解法的时间复杂度是n^2，运行有可能会超时，不推荐
// 解法二：双指针，只需要一层遍历
// 具体解法：左指针和右指针的初始值分别为0和len(height)-1，遍历条件左指针小于右指针
// 每一次遍历都需要求出新的容量，比较当前的容量，如果大于就赋值为最大容量
// 判断左右指针的高度，如果左指针高度小，就左指针++
// 右指针高度小，就右指针--
