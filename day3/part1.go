package main

import (
	"strconv"
)

// Notation:
//    path is a list of segments
//    segment is individual direction (ex U4)
//    vector is numeric conversion of segment

func seg2Vec(seg string) [2]int {
	val, _ := strconv.Atoi(seg[1:])
	switch seg[0] {
	case 'U':
		return [2]int{0, val}
	case 'D':
		return [2]int{0, -val}
	case 'R':
		return [2]int{val, 0}
	case 'L':
		return [2]int{-val, 0}
	default:
		return [2]int{0, 0}
	}
}

func path2Vec(path []string) (vecpath [][2]int) {
	for _, seg := range path {
		vecpath = append(vecpath, seg2Vec(seg))
	}
	return
}

func part1() {
}
