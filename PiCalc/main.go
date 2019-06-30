package main

import "fmt"

type PITerm interface {
	setTerm(int)
	//getTerm()
	calculateTerm() float64
	getFinalCalculation() func(float64) float64
}

type Nilakantha struct {
	term      int
	isOddTerm bool
}

func (n *Nilakantha) setTerm(term int) {
	n.term = term
	n.isOddTerm = (term%2 != 0)

}

/*
func (n Nilakantha) getTerm() int {
	return n.term
}*/
func (n Nilakantha) calculateTerm() float64 {
	var d float64
	var factor float64
	d = (2.0) * (float64(n.term) + 1.0)
	factor = d * (d + 1.0) * (d + 2.0)
	ret := 4.0 / factor
	if n.isOddTerm {
		ret = ret * -1
	}
	return ret
}

func (n Nilakantha) getFinalCalculation() func(float64) float64 {
	return func(x float64) float64 {
		return 3.0 + x
	}
}

func createTerm(numOfStep int) []PITerm {
	terms := make([]PITerm, numOfStep, numOfStep)

	for i := 0; i < numOfStep; i++ {
		n := Nilakantha{}
		piterm := PITerm(&n)
		piterm.setTerm(i)

		//terms = append(terms, piterm)
		terms[i] = piterm
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

func main() {
	terms := createTerm(1000)
	pi := executeTerm(terms)
	fmt.Println(pi)
}
