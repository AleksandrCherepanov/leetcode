package main

func finalPrices(prices []int) []int {
	stack := make([]int, 0, len(prices))
	ans := make([]int, len(prices))

	for i := 0; i < len(prices); i++ {
		for len(stack) != 0 && prices[stack[len(stack)-1]] >= prices[i] {
			ans[stack[len(stack)-1]] = prices[stack[len(stack)-1]] - prices[i]
			stack = stack[:len(stack)-1]
		}

		stack = append(stack, i)
		ans[i] = prices[i]
	}

	return ans
}
