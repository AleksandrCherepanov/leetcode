package main

type MyQueue struct {
	Stack []int
	Queue []int
}

func Constructor() MyQueue {
	return MyQueue{
		Stack: make([]int, 0, 0),
		Queue: make([]int, 0, 0),
	}
}

func (this *MyQueue) Push(x int) {
	this.Stack = append(this.Stack, x)
}

func (this *MyQueue) Pop() int {
	if len(this.Queue) == 0 {
		for len(this.Stack) != 0 {
			last := this.Stack[len(this.Stack)-1]
			this.Stack = this.Stack[:len(this.Stack)-1]
			this.Queue = append(this.Queue, last)
		}
	}

	first := this.Queue[len(this.Queue)-1]
	this.Queue = this.Queue[:len(this.Queue)-1]
	return first
}

func (this *MyQueue) Peek() int {
	if len(this.Queue) == 0 {
		for len(this.Stack) != 0 {
			last := this.Stack[len(this.Stack)-1]
			this.Stack = this.Stack[:len(this.Stack)-1]
			this.Queue = append(this.Queue, last)
		}
	}

	return this.Queue[len(this.Queue)-1]
}

func (this *MyQueue) Empty() bool {
	return len(this.Stack) == 0 && len(this.Queue) == 0
}
