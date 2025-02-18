package mergeintervals

func merge(intervals [][]int) [][]int {

	// build min heap
	n := len(intervals)
	for i := n - 1; i >= 0; i-- {
		heapify(intervals, i, n)
	}

	cap := 1        // 有效的節點數 (控制返回的數量)
	edge := n - cap // 邊際節點

	// 利用 heap sort 的特性
	// 不斷取出最小元素，然後檢查是否與 edge 重疊：
	// 1. 完全沒有重疊：
	//    - 有效節點數+1
	//    - 將元素設定為新的邊際節點 (可能會覆蓋掉被 overlap 的節點)
	// 2. 有重疊：
	//    - 終點同時也在 edge 範圍內，不用更新 edge
	//    - 終點在 edge 之外，則更新 edge
	for i := n - 1; i >= 0; i-- {

		intervals[0], intervals[i] = intervals[i], intervals[0]
		if i > 0 {
			heapify(intervals, 0, i)
		}

		current := intervals[i]

		// 當前節點的起點在 edge 之外，是一個新的有效節點
		if intervals[edge][1] < current[0] {
			cap += 1       // 有效節點數+1
			edge = n - cap // 設定新的 edge
			intervals[edge] = current
			continue
		}

		// 檢查是否要更新 edge 終點
		if intervals[edge][1] < current[1] {
			intervals[edge][1] = current[1] // 更新 edge 節點的終點 ([1,3], [2,6] => [1,6])
		}
	}

	// reverse
	for i := 0; i < n/2; i++ {
		j := n - i - 1
		intervals[i], intervals[j] = intervals[j], intervals[i]
	}

	return intervals[:cap]
}

// min heap
func heapify(intervals [][]int, index, cap int) {
	for {
		smallest := index
		left := index*2 + 1
		right := index*2 + 2

		if right < cap && shouldMoveUp(intervals, right, smallest) {
			smallest = right
		}

		if left < cap && shouldMoveUp(intervals, left, smallest) {
			smallest = left
		}

		if smallest == index {
			break
		}

		intervals[index], intervals[smallest] = intervals[smallest], intervals[index]
		index = smallest
	}
}

// 判斷 i 是否要往上移動
func shouldMoveUp(intervals [][]int, i, j int) bool {
	// 如果起點一樣，則終點越大，涵蓋範圍越廣，應該要排前面
	if intervals[i][0] == intervals[j][0] {
		return intervals[i][1] > intervals[i][1]
	}

	// 起點不一樣，則起點小的放前面
	return intervals[i][0] < intervals[j][0]
}
