package leetcode

import "testing"

type Testcases struct {
	Name   string
	Nums   []int
	Target int
}

func TestTwoSum(t *testing.T) {
	testcase := []Testcases{
		Testcases{
			Name:   "t1",
			Nums:   []int{2, 7, 11, 15},
			Target: 9,
		},
		Testcases{
			Name:   "t2",
			Nums:   []int{3, 2, 4},
			Target: 6,
		},
		Testcases{
			Name:   "t3",
			Nums:   []int{3, 3},
			Target: 6,
		},
		Testcases{
			Name:   "t4",
			Nums:   []int{1, 1, 1, 2, 2, 3, 3, 3, 4, 4, 5, 5},
			Target: 5,
		},
		Testcases{
			Name:   "t5",
			Nums:   []int{-1, -2, -3, -4, -5},
			Target: -8,
		},
		Testcases{
			Name:   "t6",
			Nums:   []int{-1, -2, -3, -4, 5},
			Target: -8,
		},
		Testcases{
			Name:   "t7",
			Nums:   []int{-3, 4, 3, 90},
			Target: 0,
		},
	}

	expected := [][]int{
		[]int{1, 0},
		[]int{2, 1},
		[]int{0, 1},
		[]int{8, 0},
		[]int{4, 2},
		[]int{},
		[]int{2, 0},
	}

	compare := func(a []int, b []int) bool {
		if len(a) != len(b) {
			return false
		}

		if len(a) == len(b) && len(a) == 0 {
			return true
		}

		if a[0] != b[0] || a[1] != b[1] {
			return false
		}
		return true
	}
	for key, val := range testcase {
		got := TwoSum(val.Nums, val.Target)
		if !compare(expected[key], got) {
			t.Errorf("%s testcase: %v \ngot: %v, expected: %v", val.Name, val.Nums, got, expected[key])
		}
	}
}
