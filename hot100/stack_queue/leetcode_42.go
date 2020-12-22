package stack_queue

import "fmt"

// 接雨水问题，第一种思路，暴力法，直接遍历左边和右边柱子的大于当前柱子的最小高度
// min(maxLeft-current,maxRight-current)
func trap(height []int) int {
	count := 0
	for i := 1; i < len(height); i++ {
		maxLeft := 0
		for l := 0; l < i; l++ {
			if height[l] > height[i] && height[l] > maxLeft {
				maxLeft = height[l]
			}
		}
		if maxLeft == 0 {
			continue
		}
		maxRight := 0
		for r := i + 1; r < len(height); r++ {
			if height[r] > height[i] && height[r] > maxRight {
				maxRight = height[r]
			}
		}
		if maxRight == 0 {
			continue
		}
		x := min(maxLeft-height[i], maxRight-height[i])
		fmt.Printf("i=%+v trap:%+v \n", i, x)
		count = count + x

	}
	return count
}

// 第二种思路，优化暴力法，通过数组减少重复计算
func trap2(height []int) int {
	if len(height) <= 0 {
		return 0
	}
	count := 0
	maxLeftSlice := make([]int, len(height))
	maxLeftSlice[0] = height[0]
	maxRightSlice := make([]int, len(height))
	maxRightSlice[len(height)-1] = height[len(height)-1]
	// 一次性计算没个节点，左右最大值
	for i := 1; i < len(height); i++ {
		maxLeft := maxLeftSlice[i-1]
		if maxLeft >= height[i] {
			maxLeftSlice[i] = maxLeft
		} else {
			maxLeftSlice[i] = height[i]
		}
		maxRight := maxRightSlice[len(maxRightSlice)-1-i+1]
		if maxRight >= height[len(height)-1-i] {
			maxRightSlice[len(maxRightSlice)-1-i] = maxRight
		} else {
			maxRightSlice[len(maxRightSlice)-1-i] = height[len(maxRightSlice)-1-i]
		}
		fmt.Printf("maxLeftSlice[%+v]=%+v \n", i, maxLeftSlice[i])
		fmt.Printf("maxRightSlice[%+v]=%+v \n", len(maxRightSlice)-i, maxRightSlice[len(maxRightSlice)-i])
	}
	for i := 1; i < len(height); i++ {
		maxLeft := maxLeftSlice[i]
		maxRight := maxRightSlice[i]
		x := min(maxLeft-height[i], maxRight-height[i])
		count = count + x
	}
	return count
}

// todo: 双指针法
func trap3(height []int) int {
	return 0
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
