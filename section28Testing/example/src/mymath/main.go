// Package mymath provides ACME inc math solutions.
package mymath

// Sum1 adds an unlimited number of values of type int.
func Sum1(xi ...int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum
}
