package main

import "fmt"

func twoSum(nums []int, target int) []int {
	data_map := make(map[int]int)
	var i, j, complement int
	var ok bool

	for i = 0; i < len(nums); i++ {
		complement = target - nums[i]
		j, ok = data_map[nums[i]]
		if (ok) {
			break
		} else {
			data_map[complement] = i
			fmt.Println(data_map)
		}
	}

	return []int {j, i}
}

func main() {
	var data = []int {3, 2, 4}
	var target int = 6
	fmt.Println(twoSum(data, target))
}