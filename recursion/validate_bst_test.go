package recursion

import "testing"

type validateBstTestCase struct {
	name     string
	node     *TreeNode
	expected validateBstResult
}

type validateBstResult bool

func TestValidateBst(t *testing.T) {
	testCases := []validateBstTestCase{
		{
			"Case #1",
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
			true,
		},
		{
			"Case #2",
			&TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val:   1,
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val: 4,
					Left: &TreeNode{
						Val:   3,
						Left:  nil,
						Right: nil,
					},
					Right: &TreeNode{
						Val:   6,
						Left:  nil,
						Right: nil,
					},
				},
			},
			false,
		},
		{
			"Case #3",
			&TreeNode{},
			true,
		},
		{
			"Case #4",
			&TreeNode{
				Val:   5,
				Left:  nil,
				Right: nil,
			},
			true,
		},
		{
			"Case #5",
			&TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val:   4,
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val: 6,
					Left: &TreeNode{
						Val:   3,
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
			false,
		},
		{
			"Case #6",
			&TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val:   0,
						Left:  nil,
						Right: nil,
					},
					Right: &TreeNode{
						Val:   2,
						Left:  nil,
						Right: nil,
					},
				},
				Right: &TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val:   4,
						Left:  nil,
						Right: nil,
					},
					Right: &TreeNode{
						Val:   6,
						Left:  nil,
						Right: nil,
					},
				},
			},
			true,
		},
		{
			"Case #7",
			&TreeNode{
				Val: 32,
				Left: &TreeNode{
					Val: 26,
					Left: &TreeNode{
						Val:  19,
						Left: nil,
						Right: &TreeNode{
							Val:   27,
							Left:  nil,
							Right: nil,
						},
					},
					Right: nil,
				},
				Right: &TreeNode{
					Val:  47,
					Left: nil,
					Right: &TreeNode{
						Val:   56,
						Left:  nil,
						Right: nil,
					},
				},
			},
			false,
		},
		{
			"Case #8",
			&TreeNode{
				Val: 120,
				Left: &TreeNode{
					Val: 70,
					Left: &TreeNode{
						Val: 50,
						Left: &TreeNode{
							Val:   20,
							Left:  nil,
							Right: nil,
						},
						Right: &TreeNode{
							Val:   55,
							Left:  nil,
							Right: nil,
						},
					},
					Right: &TreeNode{
						Val: 100,
						Left: &TreeNode{
							Val:   75,
							Left:  nil,
							Right: nil,
						},
						Right: &TreeNode{
							Val:   110,
							Left:  nil,
							Right: nil,
						},
					},
				},
				Right: &TreeNode{
					Val: 140,
					Left: &TreeNode{
						Val: 130,
						Left: &TreeNode{
							Val:   119,
							Left:  nil,
							Right: nil,
						},
						Right: &TreeNode{
							Val:   135,
							Left:  nil,
							Right: nil,
						},
					},
					Right: &TreeNode{
						Val: 160,
						Left: &TreeNode{
							Val:   150,
							Left:  nil,
							Right: nil,
						},
						Right: &TreeNode{
							Val:   200,
							Left:  nil,
							Right: nil,
						},
					},
				},
			},
			false,
		},
	}

	for _, testCase := range testCases {
		result := validateBstResult(validateBst(testCase.node))

		if result != testCase.expected {
			validateBstPrintResult(t, testCase, result)
		}
	}
}

func validateBstPrintResult(t *testing.T, testCase validateBstTestCase, result validateBstResult) {
	t.Fatalf(`%v: Expected: %v, Actual: %v`, testCase.name, testCase.expected, result)
}
