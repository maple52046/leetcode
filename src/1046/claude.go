package main

// Solution provided by Claude 3.5 Sonnet
//
// 此解法使用 max heap 資料結構來解決問題
// 1. 首先將輸入的 stones 陣列建立成 max heap
// 2. 每次從 heap 中取出最大的兩個石頭進行碰撞
// 3. 如果碰撞後還有剩餘，則將剩餘的石頭放回 heap
// 4. 重複此過程直到 heap 中只剩一顆或沒有石頭

func lastStoneWeight(stones []int) int {
	// 將輸入陣列建立成 max heap
	heapify(stones)

	n := len(stones)
	for n > 1 {
		// 取出最大的石頭 x
		x := stones[0]
		stones[0] = stones[n-1] // 將最後一個元素移到 root
		n--                     // 縮小 heap 大小
		siftDown(stones, 0, n)  // 重新調整 heap 結構

		// 取出第二大的石頭 y
		y := stones[0]
		stones[0] = stones[n-1]
		n--
		siftDown(stones, 0, n)

		// 如果碰撞後有剩餘 (x > y)，將差值放回 heap
		if diff := x - y; diff > 0 {
			stones[n] = diff
			n++
			siftUp(stones, n-1) // 將新加入的元素調整到正確位置
		}
	}

	// 如果 heap 是空的返回 0，否則返回最後一顆石頭的重量
	if n == 0 {
		return 0
	}
	return stones[0]
}

// heapify 將一個陣列轉換成 max heap
// 從最後一個非葉節點開始，依序往上進行 siftDown
func heapify(arr []int) {
	for i := len(arr)/2 - 1; i >= 0; i-- {
		siftDown(arr, i, len(arr))
	}
}

// siftDown 將節點往下移動到合適的位置
// arr: 原陣列
// i: 當前節點的索引
// n: heap 的有效大小
func siftDown(arr []int, i, n int) {
	for {
		largest := i
		left := 2*i + 1  // 左子節點索引
		right := 2*i + 2 // 右子節點索引

		// 找出當前節點、左子節點和右子節點中的最大值
		if left < n && arr[left] > arr[largest] {
			largest = left
		}
		if right < n && arr[right] > arr[largest] {
			largest = right
		}

		// 如果當前節點就是最大的，結束調整
		if largest == i {
			break
		}

		// 交換當前節點與最大值節點
		arr[i], arr[largest] = arr[largest], arr[i]
		i = largest
	}
}

// siftUp 將節點往上移動到合適的位置
// arr: 原陣列
// i: 當前節點的索引
func siftUp(arr []int, i int) {
	for {
		parent := (i - 1) / 2 // 父節點索引
		// 如果已經到達 root 或當前節點不大於父節點，結束調整
		if parent < 0 || arr[parent] >= arr[i] {
			break
		}
		// 交換當前節點與父節點
		arr[parent], arr[i] = arr[i], arr[parent]
		i = parent
	}
}
