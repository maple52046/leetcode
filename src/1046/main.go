package main

import (
	"fmt"
)

func lastStoneWeight(stones []int) int {

	// 排序成 max heap
	for i := 0; i < len(stones); i++ {
		if i == 0 {
			continue
		}

		sort(stones, i)
		// fmt.Println("sorted stones:", stones[:i+1])
	}

	for len(stones) >= 2 {

		// 取出最大
		x := take(stones)
		// fmt.Println("take", x, "from", stones)
		stones = stones[:len(stones)-1] // 縮小 slice

		// 取出次大
		y := take(stones)
		// fmt.Println("take", y, "from", stones)

		lastId := len(stones) - 1

		switch z := x - y; z {
		case 0:
			stones = stones[:len(stones)-1] // 縮小 slice
		default:
			stones[lastId] = z
			// fmt.Println("put", z, "to", "stones[", lastId, "]")
			sort(stones, lastId)
		}

		// fmt.Println("stones:", stones)
	}

	if len(stones) == 1 {
		return stones[0]
	}

	return 0
}

func sort(stones []int, index int) {

	for {
		parentId := (index - 1) / 2
		if parentId < 0 || parentId == index {
			return
		}

		parent := stones[parentId]
		leaf := stones[index]
		if parent >= leaf {
			return
		}

		stones[parentId] = leaf
		stones[index] = parent
		index = parentId
		// fmt.Println("sort:", stones)
	}
}

func take(stones []int) int {
	// 取出最大值
	out := stones[0]

	// 將最後一個節點放到 root
	leaf := stones[len(stones)-1]
	stones[0] = leaf
	// stones[len(stones)-1] = out

	// 重新排序 (排除最後一個)
	length := len(stones) - 1 // 排除最後一個
	i := 0
	for i < length {

		leftId := i*2 + 1
		rightId := i*2 + 2

		var childId int

		switch {
		// 左右節點都在
		case leftId < length && rightId < length:
			// 找出哪個比較大，然後用最大的跟當前節點比較
			if stones[leftId] > stones[rightId] {
				childId = leftId
			} else {
				childId = rightId
			}

		// 只有左節點在
		case leftId < length:
			childId = leftId

		// 只有右節點在
		case leftId < length:
			childId = rightId

		// 左右節點都不在
		default:
			childId = i
		}

		// 左右都不在，直接結束
		if childId == i {
			break
		}

		child := stones[childId]
		current := stones[i]

		// 子節點沒有比當前節點大，無需更換
		if child <= current {
			break
		}

		stones[childId] = current
		stones[i] = child

		i = childId
	}

	return out
}

func main() {

	stones := []int{2, 7, 4, 1, 8, 1}
	fmt.Println(lastStoneWeight(stones))

	stones = []int{1, 3}
	fmt.Println(lastStoneWeight(stones))

}
