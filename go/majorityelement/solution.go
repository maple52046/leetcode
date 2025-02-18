package majorityelement

func majorityElement(nums []int) int {

	var majority = -1
	var max = -1

	counts := map[int]int{}
	for i := 0; i < len(nums); i++ {

		num := nums[i]
		count, ok := counts[num]
		if !ok {
			count = 0
		}

		count += 1
		counts[num] = count

		if count > max {
			majority = num
			max = count
		}
	}

	return majority
}

func majorityElementV2(nums []int) int {

	// build heap
	n := len(nums)
	for i := n/2 - 1; i >= 0; i-- {
		heapify(nums, i, n)
	}

	var (
		majority int = -1
		max      int = -1
		element  int = -1
		count    int = -1
	)

	for i := n - 1; i >= 0; i-- {

		if element != nums[0] {
			// reset current count and element
			element = nums[0]
			count = 0
		}

		count += 1

		if count > max {
			majority = nums[0]
			max = count
		}

		if i > 0 {
			nums[0], nums[i] = nums[i], nums[0]
			heapify(nums, 0, i-1)
		}
	}

	return majority
}

func heapify(nums []int, parent int, cap int) {
	for {
		smallest := parent
		left := parent*2 + 1
		right := parent*2 + 2

		if left < cap && nums[left] < nums[parent] {
			smallest = left
		}

		if right < cap && nums[right] < nums[smallest] {
			smallest = right
		}

		if smallest == parent {
			break
		}

		nums[parent], nums[smallest] = nums[smallest], nums[parent]
		parent = smallest
	}
}
