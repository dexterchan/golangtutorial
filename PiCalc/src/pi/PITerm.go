package pi

type PITerm interface {
	SetTerm(int)
	//getTerm()
	CalculateTerm() float64
	GetFinalCalculation() func(float64) float64
}
