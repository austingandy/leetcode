package main

func maxOperations(nums []int, k int) int {
	diffs := make(map[int]int)
	pairs := 0
	for _, num := range nums {

		if v, ok := diffs[k-num]; !ok || v < 1 {
			diffs[num] += 1
			continue
		}
		diffs[k-num] -= 1
		pairs += 1
	}
	return pairs
}
