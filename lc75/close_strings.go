package main

func closeStrings(word1 string, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}
	lfreq := make(map[byte]int)
	rfreq := make(map[byte]int)
	for i := 0; i < len(word1); i += 1 {
		lfreq[word1[i]]++
		rfreq[word2[i]]++
	}
	if len(lfreq) != len(rfreq) {
		return false
	}
	for k := range lfreq {
		if _, ok := rfreq[k]; !ok {
			return false
		}
	}
	lcount, rcount := make(map[int]int), make(map[int]int)
	for k := range lfreq {
		lcount[lfreq[k]]++
		rcount[rfreq[k]]++
	}
	for k := range lcount {
		if lcount[k] != rcount[k] {
			return false
		}
	}
	return true
}
