package dynamic

/*
给定一个无序的整数数组，找到其中最长上升子序列的长度。
输入: [10,9,2,5,3,7,101,18]
输出: 4
解释: 最长的上升子序列是 [2,3,7,101]，它的长度是 4。

说明:

可能会有多种最长上升子序列的组合，你只需要输出对应的长度即可。
你算法的时间复杂度应该为 O(n2) 。
进阶: 你能将算法的时间复杂度降低到 O(n log n) 吗?
*/
// 思路：dp[i] 表示num[i] 这个数结尾的最长递增子序列的长度。同时dp[i]的初始值应该为1，至少包含它自己
// 参考：https://github.com/labuladong/fucking-algorithm/blob/master/%E5%8A%A8%E6%80%81%E8%A7%84%E5%88%92%E7%B3%BB%E5%88%97/%E5%8A%A8%E6%80%81%E8%A7%84%E5%88%92%E8%AE%BE%E8%AE%A1%EF%BC%9A%E6%9C%80%E9%95%BF%E9%80%92%E5%A2%9E%E5%AD%90%E5%BA%8F%E5%88%97.md
func lengthOfLIS(nums []int) int {
	// dp[i]的初始值应该为1，至少包含它自己
	dp := make([]int, len(nums))
	for i := 0; i < len(dp); i++ {
		dp[i] = 1
	}
	// 计算dp[i],
	for i := 0; i < len(nums); i++ {
		for j := 0; j < i; j++ { // 找出num[i]之前最大的上升子序列
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}
	res := 0
	for i := 0; i < len(dp); i++ {
		res = max(res, dp[i])
	}
	return res
}

func lengthOfLIS2(nums []int) int {
	top := make([]int, len(nums))
	piles := 0
	for i := 0; i < len(nums); i++ {
		poker := nums[i]
		left := 0
		right := piles
		for left < right {
			mid := (left + right) / 2
			if top[mid] > poker {
				right = mid
			} else if top[mid] < poker {
				left = mid + 1
			} else {
				right = mid
			}
		}
		if left == piles {
			piles++
		}
		top[left] = poker
	}
	return piles
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
