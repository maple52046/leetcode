package twosum

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	// 定義測試案例
	tests := []struct {
		name     string
		nums     []int
		target   int
		expected []int
	}{
		{
			name:     "case 1",
			nums:     []int{3, 2, 4},
			target:   6,
			expected: []int{1, 2},
		},
		// 可以添加更多測試案例
		{
			name:     "case 2",
			nums:     []int{2, 7, 11, 15},
			target:   9,
			expected: []int{0, 1},
		},
		{
			name:     "case 3",
			nums:     []int{3, 3},
			target:   6,
			expected: []int{0, 1},
		},
	}

	// 執行測試案例
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := twoSum(tt.nums, tt.target)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("twoSum(%v, %d) = %v; want %v",
					tt.nums, tt.target, result, tt.expected)
			}
		})
	}
}
