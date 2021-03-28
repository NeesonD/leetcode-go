package _015_3Sum

import "sort"

func ThreeSum(nums []int) [][]int {

	sort.Ints(nums)

	result := make([][]int, 0)
	dict := make(map[int][]int)
	for i, num := range nums {
		dict[num] = append(dict[num], i)
	}

	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < len(nums); j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			a := nums[i] + nums[j]
			indexList, ok := dict[-a]
			if ok {
				for _, index := range indexList {
					if index > j {
						result = append(result, []int{nums[i], nums[j], nums[index]})
						break
					}
				}
			}
		}
	}
	return result
}
