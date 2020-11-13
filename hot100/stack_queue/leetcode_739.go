package stack_queue

/**
请根据每日 气温 列表，重新生成一个列表。对应位置的输出为：要想观测到更高的气温，至少需要等待的天数。如果气温在这之后都不会升高，请在该位置用 0 来代替。

例如，给定一个列表 temperatures = [73, 74, 75, 71, 69, 72, 76, 73]，你的输出应该是 [1, 1, 4, 2, 1, 1, 0, 0]。

提示：气温 列表长度的范围是 [1, 30000]。每个气温的值的均为华氏度，都是在 [30, 100] 范围内的整数。
*/

// 暴力解法，直接使用遍历解决
func dailyTemperatures(T []int) []int {
	var result []int
	for i := 0; i < len(T); i++ {
		v := T[i]
		distance := 0
		// 从当前元素的后面一个开始比较
		for j := i + 1; j < len(T); j++ {
			if T[j] > v {
				distance = j - i
				break
			}
		}
		result = append(result, distance)
	}
	return result
}

// 利用栈来处理,单调递减栈，https://github.com/labuladong/fucking-algorithm/blob/master/%E6%95%B0%E6%8D%AE%E7%BB%93%E6%9E%84%E7%B3%BB%E5%88%97/%E5%8D%95%E8%B0%83%E6%A0%88.md
func dailyTemperatures2(T []int) []int {
	// 初始化结果数组，默认值为0
	res := make([]int, len(T))
	var stack []int
	//数组倒数着计算
	for i := len(T) - 1; i >= 0; i-- {
		// 栈里有值，并且栈顶值小于等于当前值，持续出栈
		for len(stack) > 0 && T[i] >= T[stack[len(stack)-1]] {
			stack = stack[:len(stack)-1]
		}
		// 栈顶值大于当前值，则根据坐标计算间隔
		if len(stack) > 0 {
			res[i] = stack[len(stack)-1] - i
		}
		// 将当前节点下标值压栈
		stack = append(stack, i)
	}
	return res
}
