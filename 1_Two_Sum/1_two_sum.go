package leetcode

import "sort"

func TwoSum(nums []int, target int) []int {
	var ans []int

	mp := make(map[int][]int)
	for key, val := range nums {
		_, ok := mp[val]
		if ok {
			data := mp[val]
			data = append(data, key)
			mp[val] = data
		} else {
			mp[val] = []int{key}
		}
	}

	sort.Slice(nums, func(i, j int) bool {
		return i > j
	})
	for _, val := range nums {
		diff := target - val
		_, ok := mp[diff]

		if ok {
			if diff == val {
				if len(mp[val]) > 1 {
					data := mp[val]
					ans = append(ans, data[0])
					ans = append(ans, data[1])
					break
				}
			} else {
				data := mp[val]
				ans = append(ans, data[0])
				data = mp[diff]
				ans = append(ans, data[0])
				break
			}
		}
	}
	return ans
}
