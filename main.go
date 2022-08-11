package main

import "github.com/AleksandrCherepanov/leetcode/structure"

func main() {
	root := structure.NewTreeNode(
		1,
		structure.NewTreeNode(
			2,
			structure.NewTreeNode(
				3,
				nil,
				nil,
			),
			structure.NewTreeNode(
				4,
				structure.NewTreeNode(5, nil, nil),
				structure.NewTreeNode(6, nil, nil),
			),
		),
		structure.NewTreeNode(
			7,
			nil,
			structure.NewTreeNode(
				8,
				structure.NewTreeNode(
					9,
					nil,
					nil,
				),
				nil,
			),
		),
	)

	root.PrintPreOrder()
	root.PrintInOrder()
}
