package main

import "math"

type Values struct {
	smallestBefore, largestAfter int
}

func increasingTriplet(nums []int) bool {
	smallest, middle := math.MaxInt, math.MaxInt
	for _, num := range nums {
		if num <= smallest {
			smallest = num
		} else if num <= middle {
			middle = num
		} else {
			return true
		}
	}
	return false
}

func dumbIncreasingTriplet(nums []int) bool {
	values := make([]Values, len(nums))
	values[0].smallestBefore = nums[0]
	values[len(nums)-1].largestAfter = nums[len(nums)-1]
	for i := 1; i < len(nums); i += 1 {
		lastSmallest := values[i-1].smallestBefore
		if nums[i-1] < lastSmallest {
			values[i].smallestBefore = nums[i-1]
		} else {
			values[i].smallestBefore = lastSmallest
		}
		oppositeIndex := len(nums) - 1 - i
		prevIndex := oppositeIndex + 1
		lastLargest := values[prevIndex].largestAfter
		if nums[prevIndex] > lastLargest {
			values[oppositeIndex].largestAfter = nums[prevIndex]
		} else {
			values[oppositeIndex].largestAfter = lastLargest
		}
	}
	for i := range nums {
		if values[i].smallestBefore < nums[i] && values[i].largestAfter > nums[i] {
			return true
		}
	}
	return false
}
