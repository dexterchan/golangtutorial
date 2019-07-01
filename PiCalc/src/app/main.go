package main

import (
	"fmt"
	"pi"
)

/*
func createTerm(numOfStep int) []PITerm {
	terms := make([]PITerm, 0, numOfStep)

	for i := 0; i < numOfStep; i++ {
		n := Nilakantha{}
		piterm := PITerm(&n)
		piterm.setTerm(i)

		terms = append(terms, piterm)
		//terms[i] = piterm
	}
	return terms
}
func executeTerm(terms []PITerm) float64 {
	var accum float64 = 0.0
	for _, v := range terms {
		r := v.calculateTerm()
		accum += r
	}
	finalValueFunc := terms[0].getFinalCalculation()
	finalValue := finalValueFunc(accum)
	return finalValue
}
*/
func createTerm (numOfStep int) []pi.PITerm{
	terms := make([]pi.PITerm, 0, numOfStep)
	return terms
}
func main() {
	
	terms := createTerm(1000)
	//pi := executeTerm(terms)
	fmt.Println(terms)
}
