package stack_queue

import "fmt"

/*
给定一个用字符数组表示的 CPU 需要执行的任务列表。其中包含使用大写的 A - Z 字母表示的26 种不同种类的任务。任务可以以任意顺序执行，并且每个任务都可以在 1 个单位时间内执行完。CPU 在任何一个单位时间内都可以执行一个任务，或者在待命状态。

然而，两个相同种类的任务之间必须有长度为 n 的冷却时间，因此至少有连续 n 个单位时间内 CPU 在执行不同的任务，或者在待命状态。

你需要计算完成所有任务所需要的最短时间。
*/

// 核心公式 (最多的任务种类-1)*(长度为 n 的冷却时间+1)+最后一行多出来的任务个数
func leastInterval(tasks []byte, n int) int {
	cnt := map[byte]int{}
	// 计算各种任务执行次数
	for _, t := range tasks {
		cnt[t]++
	}

	maxExec, maxExecCnt := 0, 0
	for key, c := range cnt {
		fmt.Printf("key: %+v ,cnt:%+v \n", key, c)
		if c > maxExec {
			// 获得最多任务种类，最后一行的任务数为固定为1
			maxExec, maxExecCnt = c, 1
		} else if c == maxExec {
			// 和最大相同的任务数量一致的时候，最后一行要加上一个任务
			maxExecCnt++
		}
	}

	if time := (maxExec-1)*(n+1) + maxExecCnt; time > len(tasks) {
		return time
	}
	return len(tasks)
}
