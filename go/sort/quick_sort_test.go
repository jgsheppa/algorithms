package sort_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/jgsheppa/algorithms/go/sort"
)

func TestQuckSortWithFixedPivot(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		input       []int
		want        []int
		description string
	}{{
		input:       []int{},
		want:        []int{},
		description: "list with no length",
	}, {
		input:       []int{1},
		want:        []int{1},
		description: "list with length of 1",
	}, {
		input:       []int{2, 1},
		want:        []int{1, 2},
		description: "list with length of 2",
	}, {
		input:       []int{2, 1, 3},
		want:        []int{1, 2, 3},
		description: "list with length of 3",
	}, {
		input:       []int{4, 2, 1, 3},
		want:        []int{1, 2, 3, 4},
		description: "list with length of 4",
	}, {
		input:       []int{4, 2, 5, 1, 3},
		want:        []int{1, 2, 3, 4, 5},
		description: "list with length of 5",
	}}

	for _, testCase := range testCases {
		t.Run(testCase.description, func(t *testing.T) {
			got := sort.QuckSortWithFixedPivot(testCase.input)

			if !cmp.Equal(got, testCase.want) {
				t.Errorf("got %v but want %v", got, testCase.want)
			}
		})
	}
}

func TestQuckSortWithRandomPivot(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		input       []int
		want        []int
		description string
	}{{
		input:       []int{},
		want:        []int{},
		description: "list with no length",
	}, {
		input:       []int{1},
		want:        []int{1},
		description: "list with length of 1",
	}, {
		input:       []int{2, 1},
		want:        []int{1, 2},
		description: "list with length of 2",
	}, {
		input:       []int{2, 1, 3},
		want:        []int{1, 2, 3},
		description: "list with length of 3",
	}, {
		input:       []int{4, 2, 1, 3},
		want:        []int{1, 2, 3, 4},
		description: "list with length of 4",
	}, {
		input:       []int{5, 2, 5, 1, 3},
		want:        []int{1, 2, 3, 5, 5},
		description: "list with length of 5",
	}}

	qs := sort.NewQuickSort()

	for _, testCase := range testCases {
		t.Run(testCase.description, func(t *testing.T) {
			got := qs.QuckSortWithRandomPivot(testCase.input)

			if !cmp.Equal(got, testCase.want) {
				t.Errorf("got %v but want %v", got, testCase.want)
			}
		})
	}
}

func BenchmarkQuckSortWithRandomPivot(b *testing.B) {
	qs := sort.NewQuickSort()
	for i := 0; i < b.N; i++ {
		qs.QuckSortWithRandomPivot([]int{1, 2, 3, 5, 5, 1, 2, 3, 5, 6, 7, 8, 9, 20, 30, 40, 1, 2, 3, 5, 5, 1, 2, 3, 5, 6, 7, 8, 9, 20, 30, 40, 1, 2, 3, 5, 5, 1, 2, 3, 5, 6, 7, 8, 9, 20, 30, 40})
	}
}

func BenchmarkQuckSortWithFixedPivot(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sort.QuckSortWithFixedPivot([]int{1, 2, 3, 5, 5, 1, 2, 3, 5, 6, 7, 8, 9, 20, 30, 40, 1, 2, 3, 5, 5, 1, 2, 3, 5, 6, 7, 8, 9, 20, 30, 40, 1, 2, 3, 5, 5, 1, 2, 3, 5, 6, 7, 8, 9, 20, 30, 40})
	}
}
