package mymath

import "testing"

func TestMathSeq(t *testing.T) {

	type test struct {
		data   []int
		answer int
	}
	tests := []test{
		test{
			data:   []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			answer: 55,
		},
		test{
			data:   []int{2, 4, 6, 8},
			answer: 20,
		},
	}

	for _, s := range tests {
		v := Sum1(s.data...)
		if v != s.answer {
			t.Error("value not match, expect:", s.answer)
		}
	}

}

func BenchmarkMathSeq(t *testing.B) {
	for i := 0; i < t.N; i++ {
		_ = Sum1([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}...)
	}
}
