package search

import "errors"

func BinarySearch(list []int, target int) (int, error) {
	var low int
	high := len(list) - 1

	for low <= high {
		middle := (high + low) / 2
		current := list[middle]

		if current == target {
			return current, nil
		} else if current > target {
			low = middle - 1
		} else {
			low = middle + 1
		}
	}
	return 0, errors.New("Target not found")
}

func LinearSearch(list []int, target int) (int, error) {
	for _, value := range list {
		if value == target {
			return target, nil
		}
	}

	return 0, errors.New("Target not found")
}
