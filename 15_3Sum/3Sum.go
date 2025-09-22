package sum

import (
	"fmt"
	"sort"
)

func ThreeSum(nums []int) [][]int {
	var ans [][]int
	sort.Ints(nums)

	mp := make(map[int][]int)
	indexes := make(map[string]int)

	for k, v := range nums {
		_, ok := mp[v]
		if ok {
			mp[v] = append(mp[v], k)
		} else {
			mp[v] = []int{k}
		}
	}

	// [-4,-1,-1,0,1,2]
	for k, v := range nums {
		for i := k + 1; i < len(nums); i++ {

			s := v + nums[i]
			diff := 0 - s
			_, ok := mp[diff]
			if ok {
				tmp := []int{diff, v, nums[i]}
				sort.Ints(tmp)
				key := fmt.Sprintf("%d%d%d", tmp[0], tmp[1], tmp[2])
				_, ok1 := indexes[key]
				if !ok1 && mp[diff][0] != k && mp[diff][0] != i && i != k {
					ans = append(ans, []int{diff, v, nums[i]})
					indexes[key] = 1
				}

			}
		}

		if v == nums[len(nums)-1] && len(ans) > 0 {
			break
		}
	}

	return ans
}
