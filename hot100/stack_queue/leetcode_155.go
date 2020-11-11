package stack_queue

/*
利用额外的空间存储当前栈中最小的元素
values = [-2,0,1,-3,1]
minValues = [-2,-2,-2,-3,-3]

本题解使用了最简单的slice方式实现栈，还有其他性能更高的实现方式。同时未保证线程安全
参考https://hansedong.github.io/2019/04/02/15/

*/

type MinStack struct {
	values    []int //存储所有元素
	minValues []int //辅助slice，用于存储当前最小元素
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		values:    []int{},
		minValues: []int{},
	}
}

func (this *MinStack) Push(x int) {
	this.values = append(this.values, x)
	minVal := x
	//和辅助slice最后一个元素相比
	if len(this.minValues) > 0 && this.minValues[len(this.minValues)-1] < x {
		minVal = this.minValues[len(this.minValues)-1]

	}
	this.minValues = append(this.minValues, minVal)
}

func (this *MinStack) Pop() {
	if len(this.values) > 0 {
		this.values = this.values[:len(this.values)-1]
	}
	if len(this.minValues) > 0 {
		this.minValues = this.minValues[:len(this.minValues)-1]
	}
}

// 未判断空栈处理情况
func (this *MinStack) Top() int {
	return this.values[len(this.values)-1]
}

// 未判断空栈处理情况
func (this *MinStack) GetMin() int {
	return this.minValues[len(this.minValues)-1]
}
