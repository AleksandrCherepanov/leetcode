package recursion

import "testing"

type maximumDepthBinaryTreeTestCase struct {
	name     string
	node     *TreeNode
	expected maximumDepthBinaryTreeResult
}

type maximumDepthBinaryTreeResult int

func TestMaximumDepthBinaryTree(t *testing.T) {
	testCases := []maximumDepthBinaryTreeTestCase{
		{
			"Case #1",
			&TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val:   9,
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val: 20,
					Left: &TreeNode{
						Val:   15,
						Left:  nil,
						Right: nil,
					},
					Right: &TreeNode{
						Val:   7,
						Left:  nil,
						Right: nil,
					},
				},
			},
			3,
		},
		{
			"Case #2",
			&TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:   2,
					Left:  nil,
					Right: nil,
				},
				Right: nil,
			},
			2,
		},
	}

	for _, testCase := range testCases {
		result := maximumDepthBinaryTreeResult(maximumDepthBinaryTree(testCase.node))

		if result != testCase.expected {
			maximumDepthBinaryTreePrintResult(t, testCase, result)
		}
	}
}

func maximumDepthBinaryTreePrintResult(t *testing.T, testCase maximumDepthBinaryTreeTestCase, result maximumDepthBinaryTreeResult) {
	t.Fatalf(`%v: Expected: %v, Actual: %v`, testCase.name, testCase.expected, result)
}
