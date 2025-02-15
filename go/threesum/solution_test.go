package threesum

import (
	"reflect"
	"sort"
	"testing"
)

func TestThreeSum(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected [][]int
	}{
		{
			name: "basic test",
			nums: []int{-1, 0, 1, 2, -1, -4},
			expected: [][]int{
				{-1, -1, 2},
				{-1, 0, 1},
			},
		},
		{
			name:     "empty array",
			nums:     []int{},
			expected: [][]int{},
		},
		{
			name: "only zeros",
			nums: []int{0, 0, 0},
			expected: [][]int{
				{0, 0, 0},
			},
		},
		{
			name:     "no solution",
			nums:     []int{1, 2, 4, 5},
			expected: [][]int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := threeSum(tt.nums)
			// Sort both slices for comparison
			sortSliceOfSlices(got)
			sortSliceOfSlices(tt.expected)

			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("threeSum(%v) = %v, want %v",
					tt.nums, got, tt.expected)
			}
		})
	}
}

// Helper function to sort slice of slices
func sortSliceOfSlices(slc [][]int) {
	// Sort each inner slice
	for _, s := range slc {
		sort.Ints(s)
	}
	// Sort outer slice
	sort.Slice(slc, func(i, j int) bool {
		for k := 0; k < len(slc[i]) && k < len(slc[j]); k++ {
			if slc[i][k] != slc[j][k] {
				return slc[i][k] < slc[j][k]
			}
		}
		return len(slc[i]) < len(slc[j])
	})
}
