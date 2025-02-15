package groupthepeople

import (
	"fmt"
)

func groupThePeople(groupSizes []int) [][]int {
	elements := make(map[int][]int)

	for index, size := range groupSizes {
		_, ok := elements[size]
		if !ok {
			elements[size] = make([]int, 0)
		}
		elements[size] = append(elements[size], index)
	}

	ans := make([][]int, 0)
	for size, indexes := range elements {
		for i := 0; i < len(indexes); i += size {
			el := make([]int, 0)
			for j := i; j < (i + size); j++ {
				el = append(el, indexes[j])
			}
			ans = append(ans, el)
		}
	}

	return ans
}

func main() {
	groupSizes := []int{3, 3, 3, 3, 3, 1, 3}
	ans := groupThePeople(groupSizes)
	fmt.Println(ans)
}
