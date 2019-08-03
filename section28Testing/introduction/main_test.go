package main

import "testing"

func TestMySum(t *testing.T) {
	x := mySum(2, 3)
	if x != 5 {
		t.Error("Expected", 5, "Got", x)
	}
}

func TestMySumByTableTest(t *testing.T) {
	type test struct {
		data   []int
		answer int
	}
	tests := []test{
		test{
			data:   []int{20, 80},
			answer: 100,
		},
		test{
			data:   []int{30, 70},
			answer: 100,
		},
	}
	for _, v := range tests {
		x := mySum(v.data...)
		if x != v.answer {
			t.Error("Expected", v.answer, "Got", x)
		}
	}
}
