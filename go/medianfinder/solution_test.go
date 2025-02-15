package medianfinder

import (
	"testing"
)

func TestMedianFinder(t *testing.T) {
	tests := []struct {
		name       string
		operations []struct {
			op   string
			val  int
			want float64
		}
	}{
		{
			name: "test case 1",
			operations: []struct {
				op   string
				val  int
				want float64
			}{
				{"addNum", -1, 0},
				{"addNum", -2, 0},
				{"findMedian", 0, -1.5},
				{"addNum", -3, 0},
				{"addNum", -4, 0},
				{"addNum", -5, 0},
				{"findMedian", 0, -3},
			},
		},
		{
			name: "test case 2",
			operations: []struct {
				op   string
				val  int
				want float64
			}{
				{"addNum", 1, 0},
				{"findMedian", 0, 1},
				{"addNum", 2, 0},
				{"findMedian", 0, 1.5},
				{"addNum", 3, 0},
				{"findMedian", 0, 2},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mf := Constructor()

			for _, op := range tt.operations {
				switch op.op {
				case "addNum":
					mf.AddNum(op.val)
				case "findMedian":
					got := mf.FindMedian()
					if got != op.want {
						t.Errorf("FindMedian() = %v, want %v", got, op.want)
					}
				}
			}
		})
	}
}
