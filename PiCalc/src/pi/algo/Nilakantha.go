package algo

type Nilakantha struct {
	term      int
	isOddTerm bool
}

func (n *Nilakantha) SetTerm(term int) {
	n.term = term
	n.isOddTerm = (term%2 != 0)

}

func (n *Nilakantha) CalculateTerm() float64 {

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

func (n Nilakantha) GetFinalCalculation() func(float64) float64 {
	return func(x float64) float64 {
		return 3.0 + x
	}
}