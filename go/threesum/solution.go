package threesum

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	var ans [][]int

	for i := 0; i < len(nums); i++ {
		target := 0 - nums[i]
		ans = twoSum(nums, i+1, target, ans)
	}

	return ans
}

func twoSum(nums []int, start int, target int, ans [][]int) [][]int {
	data_map := make(map[int]int)

	for i := start; i < len(nums); i++ {
		complement := target - nums[i]
		j, ok := data_map[nums[i]]
		if ok {
			data := []int{nums[start-1], nums[i], nums[j]}
			sort.Ints(data)
			ans = insert(ans, data)
			fmt.Println(data)
		} else {
			data_map[complement] = i
		}
	}

	return ans
}

func insert(ans [][]int, data []int) [][]int {
	match := false
	for i := 0; i < len(ans) && match == false; i++ {
		fmt.Println(i)
		match = true
		for j := 0; j < len(ans[i]); j++ {
			if data[j] != ans[i][j] {
				match = false
				break
			}
		}
	}

	if match != true {
		ans = append(ans, data)
	}

	return ans
}
