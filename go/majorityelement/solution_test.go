package majorityelement

import "testing"

func TestMajorityElement(t *testing.T) {
	// 定義測試案例
	testCases := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			name:     "case 1: [3,2,3]",
			input:    []int{3, 2, 3},
			expected: 3,
		},
		{
			name:     "case 2: [2,2,1,1,1,2,2]",
			input:    []int{2, 2, 1, 1, 1, 2, 2},
			expected: 2,
		},
	}

	// 執行測試
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := majorityElement(tc.input)
			if result != tc.expected {
				t.Errorf("expected %d but got %d", tc.expected, result)
			}
		})
	}
}
