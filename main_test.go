package main

import (
	"reflect"
	"testing"
)

func TestRotate(t *testing.T) {
	cases := []struct {
		in, want [][]int
	}{
		{[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, [][]int{{7, 4, 1}, {8, 5, 2}, {9, 6, 3}}},
		{[][]int{{1}}, [][]int{{1}}},
		{[][]int{{1, 2}, {3, 4}}, [][]int{{3, 1}, {4, 2}}},
		{[][]int{}, [][]int{}},
	}
	for _, c := range cases {
		Rotate(c.in)
		if !reflect.DeepEqual(c.in, c.want) {
			t.Errorf("got %v want %v", c.in, c.want)
		}
	}
}
