package main

func pivotIndex(nums []int) int {
	var lsum, rsum int
	for i := 1; i < len(nums); i += 1 {
		rsum += i
	}
	if rsum == 0 {
		return 0
	}
	for i := 1; i < len(nums); i += 1 {
		lsum += nums[i-1]
		rsum -= nums[i]
		if lsum == rsum {
			return i
		}
	}
	return -1
}
