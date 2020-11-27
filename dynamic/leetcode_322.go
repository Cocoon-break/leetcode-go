package dynamic

import (
	"fmt"
	"math"
)

/**
给定不同面额的硬币 coins 和一个总金额 amount。编写一个函数来计算可以凑成总金额所需的最少的硬币个数。如果没有任何一种硬币组合能组成总金额，返回 -1。

你可以认为每种硬币的数量是无限的。

示例 1：

输入：coins = [1, 2, 5], amount = 11
输出：3
解释：11 = 5 + 5 + 1
*/

// https://github.com/labuladong/fucking-algorithm/blob/master/%E5%8A%A8%E6%80%81%E8%A7%84%E5%88%92%E7%B3%BB%E5%88%97/%E5%8A%A8%E6%80%81%E8%A7%84%E5%88%92%E8%AF%A6%E8%A7%A3%E8%BF%9B%E9%98%B6.md
// 暴力解法递归
func coinChange(coins []int, amount int) int {
	// 边界问题，金额
	if amount == 0 {
		return 0
	}
	if amount < 0 {
		return -1
	}
	// 设置一个最大的数
	res := math.MaxUint8
	for _, coin := range coins {
		subproblem := coinChange(coins, amount-coin)
		if subproblem == -1 { //金额小于0，不做选择
			continue
		}
		res = min(res, 1+subproblem)
	}
	fmt.Printf("amout:%+v num:%+v \n", amount, res)
	if res != math.MaxUint8 {
		return res
	}
	return -1
}

// 备忘录解法
func coinChange2(coins []int, amount int) int {
	// 使用map记录已经计算过的数组 map[金额]数量
	dict := map[int]int{}

	return memoCoinChange(coins, amount, dict)
}
func memoCoinChange(coins []int, amount int, dict map[int]int) int {
	if _, ok := dict[amount]; ok {
		return dict[amount]
	}
	if amount == 0 {
		return 0
	}
	if amount < 0 {
		return -1
	}
	res := math.MaxUint8
	for _, coin := range coins {
		subproblem := memoCoinChange(coins, amount-coin, dict)
		if subproblem == -1 {
			continue
		}
		res = min(res, 1+subproblem)
	}
	if res == math.MaxUint8 {
		res = -1
	}
	dict[amount] = res
	return dict[amount]
}

// dp数组方式
func coinChange3(coins []int, amount int) int {
	// 定义dp[amount]=数量
	dp := make([]int, amount+1)
	for i := 1; i < len(dp); i++ {
		dp[i] = amount + 1
	}
	fmt.Println(dp)
	for i := 0; i < len(dp); i++ {
		for _, coin := range coins {
			if i-coin < 0 {
				continue
			}
			dp[i] = min(dp[i], 1+dp[i-coin])
			fmt.Printf("i=%+v coin=%+v dp[%+v]=%+v \n", i, coin, i, dp[i])
		}
	}
	if dp[amount] == amount+1 {
		return -1
	}
	return dp[amount]
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
