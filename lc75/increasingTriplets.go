package main

type Values struct {
	smallestUntil, largestAfter int
}

func increasingTriplet(nums []int) bool {
	values := make([]Values, len(nums))
	values[0].smallestUntil = nums[0]
	values[len(nums)-1].largestAfter = nums[len(nums)-1]
	for i := 1; i < len(nums); i += 1 {
		lastSmallest := values[i-1].smallestUntil
		if nums[i] < lastSmallest {
			values[i].smallestUntil = nums[i]
		} else {
			values[i].smallestUntil = lastSmallest
		}
		oppositeIndex := len(nums) - 1 - i
		lastLargest := values[oppositeIndex].largestAfter
		if nums[oppositeIndex] > lastLargest {
			values[oppositeIndex].largestAfter = nums[oppositeIndex]
		} else {
			values[oppositeIndex].largestAfter = lastLargest
		}
	}
	for i := range nums {
		if values[i].smallestUntil < nums[i] && values[i].largestAfter > nums[i] {
			return true
		}
	}
	return false
}
