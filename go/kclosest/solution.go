package kclosest

func kClosest(points [][]int, k int) [][]int {

	n := len(points)

	// 建立 min heap
	//
	// 在一個完全二元樹中：
	// 1. 葉子節點的索引從 n/2 開始到 n-1
	// 2. 最後一個非葉子節點的索引是 n/2-1
	// 3. 所有葉子節點都可以被視為已經是合法的 heap（因為它們沒有子節點）
	// 4. 所以我們只需要從最後一個非葉子節點開始，向上逐層進行 heapify
	//
	// 例如對於陣列 [A,B,C,D,E,F,G]，長度為 7：
	//          A(0)
	//        /      \
	//     B(1)      C(2)
	//    /   \     /   \
	// D(3)   E(4) F(5) G(6)
	//
	// - n/2-1 = 2，從索引 2 開始
	// - 索引 3,4,5,6 都是葉子節點，不需要 heapify
	// - 只需要對索引 2,1,0 進行 heapify
	for i := n/2 - 1; i >= 0; i-- {
		heapify(points, i, n)
	}

	// 將前 k 個元素移動到 points 後面
	// 並重新排序 (排序時忽略已經移動到後面的元素)
	for i := n - 1; i >= n-k; i-- {
		points[0], points[i] = points[i], points[0]
		heapify(points, 0, i)
	}

	// 回傳最後 k 個 point
	return points[n-k:]
}

// heapify 維護 min heap 的性質
// 從給定的 index 節點開始，透過比較並交換的方式
// 確保父節點的值小於其子節點的值
//
// 運作方式：
// 1. 比較當前節點與其左右子節點的值
// 2. 如果子節點比父節點小，則與最小的子節點交換
// 3. 交換後，繼續對交換後的子節點進行相同操作
// 4. 直到當前節點小於等於其子節點，或到達 cap 界限為止
//
// 參數：
// - points: 儲存節點的陣列
// - index: 當前要處理的節點索引
// - cap: 限制 heap 的範圍，用於 heap sort 過程中逐步縮小排序範圍
func heapify(points [][]int, index, cap int) {
	for {
		left := index*2 + 1
		right := index*2 + 2
		smallest := index

		// 先比較左邊
		if left < cap && distance(points[left]) < distance(points[smallest]) {
			smallest = left
		}

		// 再比較右邊(確保選擇最小的)
		if right < cap && distance(points[right]) < distance(points[smallest]) {
			smallest = right
		}

		if smallest != index {
			// 已找到更小的子節點，交換位置，然後繼續向下檢查
			points[index], points[smallest] = points[smallest], points[index]
			index = smallest
			continue
		}
		return
	}
}

func distance(point []int) int {
	x := point[0]
	y := point[1]
	return x*x + y*y
}
