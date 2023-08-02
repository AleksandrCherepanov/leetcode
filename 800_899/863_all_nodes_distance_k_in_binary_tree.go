package main

var graph map[int][]int

func distanceK(root *TreeNode, target *TreeNode, k int) []int {
	graph = make(map[int][]int, 0)
	buildGraph(root, nil)

	queue := []int{target.Val}
	seen := make(map[int]bool, 0)
	seen[target.Val] = true
	distance := 0

	for len(queue) > 0 && distance < k {
		l := len(queue)

		for i := 0; i < l; i++ {
			el := queue[0]
			queue = queue[1:]

			for _, n := range graph[el] {
				if _, ok := seen[n]; !ok {
					queue = append(queue, n)
					seen[n] = true
				}
			}
		}
		distance++
	}

	return queue
}

func buildGraph(root *TreeNode, parent *TreeNode) {
	if root == nil {
		return
	}

	if parent != nil {
		graph[parent.Val] = append(graph[parent.Val], root.Val)
		graph[root.Val] = append(graph[root.Val], parent.Val)
	}

	buildGraph(root.Left, root)
	buildGraph(root.Right, root)
}
