package sort

import (
	"math/rand"
	"time"
)

func QuckSortWithFixedPivot(list []int) []int {
	if len(list) < 2 {
		return list
	}
	pivot := list[0]

	var lessThanPivotList []int
	var greaterThanPivotList []int

	for i, item := range list {
		if item < list[0] && i > 0 {
			lessThanPivotList = append(lessThanPivotList, item)
		} else if i > 0 && item > list[0] {
			greaterThanPivotList = append(greaterThanPivotList, item)
		}
	}

	lessThanPivotSorted := QuckSortWithFixedPivot(lessThanPivotList)
	greaterThanPivotSorted := QuckSortWithFixedPivot(greaterThanPivotList)

	var result []int
	result = append(result, lessThanPivotSorted...)
	result = append(result, pivot)
	result = append(result, greaterThanPivotSorted...)

	return result
}

type QuickSort struct {
	seed *rand.Rand
}

func NewQuickSort() *QuickSort {
	seed := rand.New(rand.NewSource(time.Now().UnixNano()))

	return &QuickSort{
		seed,
	}
}

func (qs *QuickSort) QuckSortWithRandomPivot(list []int) []int {
	if len(list) < 2 {
		return list
	}

	randomIndex := qs.seed.Intn(len(list) - 1)

	pivot := list[randomIndex]

	var lessThanPivotList []int
	var greaterThanPivotList []int

	for i, item := range list {
		if item < pivot && i != randomIndex {
			lessThanPivotList = append(lessThanPivotList, item)
		} else if item >= pivot && i != randomIndex {
			greaterThanPivotList = append(greaterThanPivotList, item)
		}
	}

	var result []int
	result = append(result, QuckSortWithFixedPivot(lessThanPivotList)...)
	result = append(result, pivot)
	result = append(result, QuckSortWithFixedPivot(greaterThanPivotList)...)

	return result
}
