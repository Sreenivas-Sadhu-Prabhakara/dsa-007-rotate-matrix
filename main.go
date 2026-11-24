package main

import "fmt"

// Rotate turns an n×n matrix 90 degrees clockwise, in place:
// transpose across the main diagonal, then reverse each row.
func Rotate(m [][]int) {
	n := len(m)
	for r := 0; r < n; r++ {
		for c := r + 1; c < n; c++ {
			m[r][c], m[c][r] = m[c][r], m[r][c]
		}
	}
	for r := 0; r < n; r++ {
		for l, x := 0, n-1; l < x; l, x = l+1, x-1 {
			m[r][l], m[r][x] = m[r][x], m[r][l]
		}
	}
}

func main() {
	m := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	Rotate(m)
	fmt.Println(m)
}
