package dynamic

/*
斐波那契数，通常用 F(n) 表示，形成的序列称为斐波那契数列。该数列由 0 和 1 开始，后面的每一项数字都是前面两项数字的和。也就是：

F(0) = 0,   F(1) = 1
F(N) = F(N - 1) + F(N - 2), 其中 N > 1.
给定 N，计算 F(N)。

输入：2
输出：1
解释：F(2) = F(1) + F(0) = 1 + 0 = 1.
*/

// 递归方式，包含大量的重复计算.时间复杂度O(2^N)
func fib(N int) int {
	if N == 1 || N == 2 {
		return 1
	}
	return fib(N-1) + fib(N-2)
}

// 自顶向下计算
func fibK(N int) int {
	if N < 1 {
		return 0
	}
	result := make([]int, N+1)
	return helper(N, result)
}

// 表格记录之前计算过的数值，无重复计算。空间复杂的O(N) 时间复杂度O(N)
func helper(N int, result []int) int {
	if N == 1 || N == 2 {
		return 1
	}
	if result[N] != 0 {
		return result[N]
	}
	result[N] = helper(N-1, result) + helper(N-2, result)
	return result[N]
}

//自底向上计算,表格记录之前计算过的数值，无重复计算。空间复杂的O(N) 时间复杂度O(N)
func fib2(N int) int {
	if N == 1 || N == 2 {
		return 1
	}
	result := make([]int, N+1)
	result[1] = 1
	result[2] = 1
	for i := 3; i <= N; i++ {
		result[i] = result[i-1] + result[i-2]
	}
	return result[N]
}

// 使用两个变量记录之前已计算数值，空间复杂的O(1)
func fib3(N int) int {
	if N == 1 || N == 2 {
		return 1
	}
	n1 := 1 //n1 N-1
	n2 := 1 //n2 N-2
	var result int
	for i := 3; i <= N; i++ {
		result = n1 + n2
		n2 = n1
		n1 = result
		// fmt.Printf("result[%+v]=%+v  n1=%+v n2=%+v\n", i, result, n1, n2)
	}
	return result
}
