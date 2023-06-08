package main

type MovingAverage struct {
	Queue []int
	Size  int
}

func Constructor(size int) MovingAverage {
	queue := MovingAverage{}
	queue.Queue = make([]int, 0)
	queue.Size = size

	return queue
}

func (this *MovingAverage) Next(val int) float64 {
	this.Queue = append(this.Queue, val)
	for len(this.Queue) > this.Size {
		this.Queue = this.Queue[1:]
	}

	sum := 0
	for i := 0; i < len(this.Queue); i++ {
		sum += this.Queue[i]
	}

	return float64(sum) / float64(len(this.Queue))
}

/**
 * Your MovingAverage object will be instantiated and called as such:
 * obj := Constructor(size);
 * param_1 := obj.Next(val);
 */
