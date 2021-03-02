package main

import (
	"fmt"
)

func computeArea(A int, B int, C int, D int, E int, F int, G int, H int) int {
	x0, y0, x1, y1 := max(A, E), max(B, F), min(C, G), min(D, H)
	return area(A, B, C, D) + area(E, F, G, H) - area(x0, y0, x1, y1)
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func area(X0, Y0, X1, Y1 int) int {
	h, w := X1-X0, Y1-Y0

	if h <= 0 || w <= 0 {
		return 0
	}

	return h * w
}

func main() {
	res := computeArea(-3, 0, 3, 4, 0, -1, 9, 2)
	fmt.Println(res)
}
