package main

import (
	"fmt"
)

type Pos struct {
	X int
	Y int
}

type Direction func(pos Pos) Pos

func right(pos Pos) Pos {
	return Pos{pos.X + 1, pos.Y}
}

func left(pos Pos) Pos {
	return Pos{pos.X - 1, pos.Y}
}
func up(pos Pos) Pos {
	return Pos{pos.X, pos.Y - 1}
}

func down(pos Pos) Pos {
	return Pos{pos.X, pos.Y + 1}
}

type Dir int

const (
	rightDir Dir = iota
	downDir
	leftDir
	upDir
)

func getNextDir(dir Dir) Dir {
	switch dir {
	case rightDir:
		return downDir
	case downDir:
		return leftDir
	case leftDir:
		return upDir
	case upDir:
		return rightDir
	}
	return rightDir
}

func getDirection(dir Dir) Direction {
	switch dir {
	case rightDir:
		return right
	case downDir:
		return down
	case leftDir:
		return left
	case upDir:
		return up
	}
	return right
}

func makeGetRuneForDir() func(dir Dir) string {
	n := 1
	return func(dir Dir) string {
		n += 1
		return fmt.Sprintf("%d", n-1)
	}
}

//func getRuneForDir(dir Dir) rune {
//	//switch dir {
//	//case rightDir:
//	//	return '>'
//	//case downDir:
//	//	return 'v'
//	//case leftDir:
//	//	return '<'
//	//case upDir:
//	//	return '^'
//	//}
//	//return ' '
//}

func printSpiral(n int) {
	getRuneForDir := makeGetRuneForDir()
	grid := make([][]string, n)
	for i := range grid {
		grid[i] = make([]string, n)
	}
	currLength := n
	currPos := Pos{0, 0}
	currDir, moveDir := rightDir, right
	for currLength > 0 {
		i := currLength
		for {
			grid[currPos.Y][currPos.X] = getRuneForDir(currDir)
			newPos := moveDir(currPos)
			i -= 1
			if i <= 0 {
				break
			}
			currPos = newPos
		}
		currDir = getNextDir(currDir)
		moveDir = getDirection(currDir)
		currLength -= 1
	}
	for _, row := range grid {
		//s := ""
		fmt.Print("[")
		s := ""
		for _, r := range row {
			s += "'" + r + "'" + ","
		}
		s = s[:len(s)-1]
		fmt.Print(s)
		fmt.Println("]")
		//fmt.Printf("%v", row)
		//println("")
	}
}
