package kclosest

import (
	"reflect"
	"sort"
	"testing"
)

func TestKClosest(t *testing.T) {
    tests := []struct {
        name   string
        points [][]int
        k      int
        want   [][]int
    }{
        {
            name:   "example 1",
            points: [][]int{{1, 3}, {-2, 2}},
            k:      1,
            want:   [][]int{{-2, 2}},
        },
        {
            name:   "example 2",
            points: [][]int{{3, 3}, {5, -1}, {-2, 4}},
            k:      2,
            want:   [][]int{{3, 3}, {-2, 4}},
        },
        {
            name:   "k equals length",
            points: [][]int{{1, 1}, {2, 2}, {3, 3}},
            k:      3,
            want:   [][]int{{1, 1}, {2, 2}, {3, 3}},
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := kClosest(tt.points, tt.k)

            // Sort both slices for comparison
            sortPoints(got)
            sortPoints(tt.want)

            if !reflect.DeepEqual(got, tt.want) {
                t.Errorf("kClosest(%v, %v) = %v, want %v",
                    tt.points, tt.k, got, tt.want)
            }
        })
    }
}

// Helper function to sort points by their coordinates
func sortPoints(points [][]int) {
    sort.Slice(points, func(i, j int) bool {
        if points[i][0] != points[j][0] {
            return points[i][0] < points[j][0]
        }
        return points[i][1] < points[j][1]
    })
}
