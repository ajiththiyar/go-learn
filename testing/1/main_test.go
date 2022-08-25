package main

import (
	"testing"

	"examp.com/sum/acdb"
)

func TestMySum(t *testing.T) {
	x := acdb.Sum(1, 2)

	if x != 3 {
		t.Error("Expected", 3, "got", x)
	}
}

func TestMySumTwo(t *testing.T) {
	type test struct {
		data   []int
		answer int
	}

	var tests []test = []test{
		{[]int{1, 2, 4, 5}, 12},
		{[]int{1, 2, 4}, 7},
		{[]int{2, 4, 5}, 11},
	}

	for _, v := range tests {
		x := acdb.Sum(v.data...)
		if x != v.answer {
			t.Error("Expected", v.answer, "got", x)
		}
	}
}
