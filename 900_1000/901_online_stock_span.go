package main

type StockSpanner struct {
	Stack []int
	Dict  map[int]int
}

func Constructor() StockSpanner {
	return StockSpanner{
		Stack: make([]int, 0, 0),
		Dict:  make(map[int]int, 0),
	}
}

func (this *StockSpanner) Next(price int) int {
	answer := 0
	for len(this.Stack) != 0 && this.Stack[len(this.Stack)-1] < price {
		last := this.Stack[len(this.Stack)-1]

		this.Stack = this.Stack[:len(this.Stack)-1]
		answer += this.Dict[last]
	}

	this.Stack = append(this.Stack, price)
	this.Dict[price] = answer
	return answer + 1
}
