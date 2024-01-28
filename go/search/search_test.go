package search_test

import (
	"errors"
	"testing"

	"github.com/jgsheppa/algorithms/go/search"
)

func TestBinarySearch_Success(t *testing.T) {
	t.Parallel()
	list := []int{1, 3, 5, 7, 9}
	target := 9

	got, err := search.BinarySearch(list, target)
	if err != nil {
		t.Errorf("could not complete binary search: %e", err)
	}

	want := 9

	if want != got {
		t.Errorf("want %d not equal to got %d", want, got)
	}
}

func TestBinarySearch_Error(t *testing.T) {
	t.Parallel()
	list := []int{1, 3, 5, 7, 9}
	target := 10

	_, got := search.BinarySearch(list, target)
	if got == nil {
		t.Error("no error returned from binary search")
	}

	want := errors.New("Target not found").Error()

	if want != got.Error() {
		t.Errorf("want %s not equal to got %s", want, got)
	}
}

func BenchmarkBinarySearch(b *testing.B) {
	list := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
	target := 13

	for i := 0; i < b.N; i++ {
		_, err := search.BinarySearch(list, target)
		if err != nil {
			b.Errorf("error executing benchmark: %e", err)
		}
	}
}

func BenchmarkLinearSearch(b *testing.B) {
	list := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
	target := 13

	for i := 0; i < b.N; i++ {
		search.LinearSearch(list, target)
	}
}
