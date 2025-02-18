package mergeintervals

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	// 定義測試案例
	testCases := []struct {
		name     string
		input    [][]int
		expected [][]int
	}{
		{
			name:     "case 1: overlapping intervals",
			input:    [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}},
			expected: [][]int{{1, 6}, {8, 10}, {15, 18}},
		},
		{
			name:     "case 2: adjacent intervals",
			input:    [][]int{{1, 4}, {4, 5}},
			expected: [][]int{{1, 5}},
		},
		{
			name:     "case 3: overlapping intervals",
			input:    [][]int{{1, 4}, {1, 5}},
			expected: [][]int{{1, 5}},
		},
		{
			name:     "case 4: empty intervals",
			input:    [][]int{{1, 4}, {2, 3}},
			expected: [][]int{{1, 4}},
		},
		{
			name:     "case 5: overlapping intervals",
			input:    [][]int{{2, 3}, {4, 5}, {6, 7}, {8, 9}, {1, 10}},
			expected: [][]int{{1, 10}},
		},
	}

	// 執行測試
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := merge(tc.input)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("expected %v but got %v", tc.expected, result)
			}
		})
	}
}
