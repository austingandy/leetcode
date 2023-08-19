package main

import "fmt"

func compress(chars []byte) int {
	writeIndex, count := 0, 0
	prev := chars[0]
	for i := range chars {
		if prev == chars[i] {
			count += 1
			continue
		}
		chars[writeIndex] = prev
		writeIndex += 1
		if count > 1 {
			for _, c := range []byte(fmt.Sprintf("%d", count)) {
				chars[writeIndex] = c
				writeIndex += 1
			}
		}
		count = 1
		prev = chars[i]
	}
	if count > 0 {
		chars[writeIndex] = prev
		writeIndex += 1
		if count > 1 {
			for _, c := range []byte(fmt.Sprintf("%d", count)) {
				chars[writeIndex] = c
				writeIndex += 1
			}
		}
	}
	return writeIndex
}
