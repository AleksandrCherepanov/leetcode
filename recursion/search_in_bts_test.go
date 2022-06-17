package recursion

import "testing"

type searchBSTTestCase struct {
	name     string
	node     *TreeNode
	val      int
	expected *TreeNode
}

type searchBSTResult *TreeNode

func TestSearchBST(t *testing.T) {
	testCases := []searchBSTTestCase{
		{
			"Case #1",
			&TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val:   1,
						Left:  nil,
						Right: nil,
					},

					Right: &TreeNode{
						Val:   3,
						Left:  nil,
						Right: nil,
					},
				},
				Right: &TreeNode{
					Val:   7,
					Left:  nil,
					Right: nil,
				},
			},
			2,
			&TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val:   1,
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val:   3,
					Left:  nil,
					Right: nil,
				},
			},
		},
		{
			"Case #2",
			&TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val:   1,
						Left:  nil,
						Right: nil,
					},

					Right: &TreeNode{
						Val:   3,
						Left:  nil,
						Right: nil,
					},
				},
				Right: &TreeNode{
					Val:   7,
					Left:  nil,
					Right: nil,
				},
			},
			5,
			nil,
		},
	}

	for _, testCase := range testCases {
		result := searchBSTResult(searchBST(testCase.node, testCase.val))

		expected := make([]int, 0, 0)
		res := make([]int, 0, 0)

		collectTreeValues(result, &res)
		collectTreeValues(testCase.expected, &expected)

		if len(res) != len(expected) {
			searchBSTPrintResult(t, testCase, expected, res)
		}

		for i, n := range res {
			if n != expected[i] {
				searchBSTPrintResult(t, testCase, expected, res)
			}
		}
	}
}

func collectTreeValues(node *TreeNode, res *[]int) {
	if node == nil {
		return
	}

	*res = append(*res, node.Val)

	collectTreeValues(node.Left, res)
	collectTreeValues(node.Right, res)
}

func searchBSTPrintResult(t *testing.T, testCase searchBSTTestCase, expected, result []int) {
	t.Fatalf(`%v: Expected: %v, Actual: %v`, testCase.name, expected, result)
}
