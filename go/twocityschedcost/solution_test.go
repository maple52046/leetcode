package twocityschedcost

import (
	"testing"
)

func TestTwoCitySchedCost(t *testing.T) {
	tests := []struct {
		name  string
		costs [][]int
		want  int
	}{
		{
			name: "example 1",
			costs: [][]int{
				{10, 20},
				{30, 200},
				{400, 50},
				{30, 20},
			},
			want: 110,
		},
		{
			name: "example 2",
			costs: [][]int{
				{259, 770},
				{448, 54},
				{926, 667},
				{184, 139},
				{840, 118},
				{577, 469},
			},
			want: 1859,
		},
		{
			name: "equal costs",
			costs: [][]int{
				{100, 100},
				{200, 200},
			},
			want: 300,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoCitySchedCost(tt.costs); got != tt.want {
				t.Errorf("twoCitySchedCost() = %v, want %v", got, tt.want)
			}
		})
	}
}
