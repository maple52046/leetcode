package laststoneweight

import "testing"

func TestLastStoneWeight(t *testing.T) {
	tests := []struct {
		name   string
		stones []int
		want   int
	}{
		{
			name:   "example 1",
			stones: []int{2, 7, 4, 1, 8, 1},
			want:   1,
		},
		{
			name:   "example 2",
			stones: []int{1, 3},
			want:   2,
		},
		{
			name:   "single stone",
			stones: []int{5},
			want:   5,
		},
		{
			name:   "same weights",
			stones: []int{2, 2, 2, 2},
			want:   0,
		},
		{
			name:   "empty array",
			stones: []int{},
			want:   0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lastStoneWeight(tt.stones); got != tt.want {
				t.Errorf("lastStoneWeight(%v) = %v, want %v",
					tt.stones, got, tt.want)
			}
		})
	}
}
