package groupthepeople

import (
	"reflect"
	"testing"
)

func TestGroupThePeople(t *testing.T) {
    tests := []struct {
        name       string
        groupSizes []int
        want       [][]int
    }{
        {
            name:       "example 1",
            groupSizes: []int{3, 3, 3, 3, 3, 1, 3},
            want: [][]int{
                {5},
                {0, 1, 2},
                {3, 4, 6},
            },
        },
        {
            name:       "example 2",
            groupSizes: []int{2, 1, 3, 3, 3, 2},
            want: [][]int{
                {1},
                {0, 5},
                {2, 3, 4},
            },
        },
        {
            name:       "single group",
            groupSizes: []int{1, 1, 1},
            want: [][]int{
                {0},
                {1},
                {2},
            },
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := groupThePeople(tt.groupSizes)
            // Sort the results for consistent comparison
            if !isEquivalentGroups(got, tt.want) {
                t.Errorf("groupThePeople(%v) = %v, want %v",
                    tt.groupSizes, got, tt.want)
            }
        })
    }
}

// Helper function to check if two group arrangements are equivalent
func isEquivalentGroups(got, want [][]int) bool {
    if len(got) != len(want) {
        return false
    }

    // Create maps of group sizes
    gotMap := make(map[int][]int)
    wantMap := make(map[int][]int)

    for _, group := range got {
        size := len(group)
        gotMap[size] = append(gotMap[size], group...)
    }

    for _, group := range want {
        size := len(group)
        wantMap[size] = append(wantMap[size], group...)
    }

    return reflect.DeepEqual(gotMap, wantMap)
}
