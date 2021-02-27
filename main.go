package main

import (
	"fmt"
)

func findOrder(numCourses int, prerequisites [][]int) []int {
	in := make([]int, numCourses)
	frees := make([][]int, numCourses)
	next := make([]int, 0, numCourses)

	for _, courses := range prerequisites {
		in[courses[0]]++
		frees[courses[1]] = append(frees[courses[1]], courses[0])
	}

	for c, i := range in {
		if i == 0 {
			next = append(next, c)
		}
	}

	for i := 0; i != len(next); i++ {
		c := next[i]

		for _, v := range frees[c] {
			in[v]--
			if in[v] == 0 {
				next = append(next, v)
			}
		}
	}

	if len(next) == numCourses {
		return next
	}

	return []int{}
}

func main() {
	res := findOrder(4, [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}})
	fmt.Println(res)
}
