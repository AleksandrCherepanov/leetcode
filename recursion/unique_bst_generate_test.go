package recursion

import "testing"

type uniqueBstGenerateTestCase struct {
	name     string
	n        int
	expected uniqueBstGenerateResult
}

type uniqueBstGenerateResult []*TreeNode

func createTreeNode(val int, left *TreeNode, right *TreeNode) *TreeNode {
	return &TreeNode{
		Val:   val,
		Left:  left,
		Right: right,
	}
}

func TestUniqueBstGenerate(t *testing.T) {
	testCases := []uniqueBstGenerateTestCase{
		{
			"Case #1",
			3,
			[]*TreeNode{
				createTreeNode(1, nil, createTreeNode(2, nil, createTreeNode(3, nil, nil))),
				createTreeNode(1, nil, createTreeNode(3, createTreeNode(2, nil, nil), nil)),
				createTreeNode(2, createTreeNode(1, nil, nil), createTreeNode(3, nil, nil)),
				createTreeNode(3, createTreeNode(1, nil, createTreeNode(2, nil, nil)), nil),
				createTreeNode(3, createTreeNode(2, createTreeNode(1, nil, nil), nil), nil),
			},
		},
		{
			"Case #2",
			1,
			[]*TreeNode{
				createTreeNode(1, nil, nil),
			},
		},
	}

	for _, testCase := range testCases {
		result := uniqueBstGenerateResult(uniqueBstGenerate(testCase.n))

		if len(result) != len(testCase.expected) {
			uniqueBstGeneratePrintResult(t, testCase, result)
		}

		for i, tree := range result {
			res := []int{}
			exp := []int{}

			collectTreeValues(tree, &res)
			collectTreeValues(testCase.expected[i], &exp)
			if len(res) != len(exp) {
				printTreeResult(t, testCase.name, res, exp)
			}

			for j, node := range res {
				if node != exp[j] {
					printTreeResult(t, testCase.name, res, exp)
				}
			}
		}
	}
}

func uniqueBstGeneratePrintResult(t *testing.T, testCase uniqueBstGenerateTestCase, result uniqueBstGenerateResult) {
	t.Fatalf(`%v: Expected: %v, Actual: %v`, testCase.name, testCase.expected, result)
}

func printTreeResult(t *testing.T, name string, res []int, exp []int) {
	t.Fatalf(`%v: Expected: %v, Actual: %v`, name, exp, res)
}
