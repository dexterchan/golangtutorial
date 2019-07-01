package pi

type PITerm interface {
	setTerm(int)
	//getTerm()
	calculateTerm() float64
	getFinalCalculation() func(float64) float64
}
