package main

type RecentCounter struct {
	queue []int
}

func Constructor() RecentCounter {
	counter := RecentCounter{}
	counter.queue = make([]int, 0)

	return counter
}

func (this *RecentCounter) Ping(t int) int {
	left := t - 3000
	count := 0
	for i := 0; i < len(this.queue); i++ {
		if this.queue[i] < left {
			count++
		} else {
			break
		}
	}

	if len(this.queue) > count {
		this.queue = this.queue[count:]
	} else {
		this.queue = make([]int, 0)
	}

	this.queue = append(this.queue, t)
	return len(this.queue)
}
