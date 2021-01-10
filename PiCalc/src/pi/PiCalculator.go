package pi

import (
	"pi/algo"
	"sync"
)

func createTermArray(numOfStep int) []PITerm {
	//terms := make([]pi.PITerm, 0, numOfStep)
	terms := make([]PITerm, 0)
	for i := 0; i < numOfStep; i++ {
		n := algo.Nilakantha{}

		piterm := PITerm(&n)
		piterm.SetTerm(i)

		terms = append(terms, piterm)

	}
	return terms
}
func executeTermArray(terms []PITerm) float64 {
	var accum float64 = 0.0
	for _, v := range terms {
		r := v.CalculateTerm()
		accum += r
	}
	finalValueFunc := terms[0].GetFinalCalculation()
	finalValue := finalValueFunc(accum)
	return finalValue
}

//CalculatePiByArray Pi calculate by array
func CalculatePiByArray(seriesLength int) float64 {
	terms := createTermArray(seriesLength)
	pi := executeTermArray(terms)
	return pi
}

//CalculatePiByFanInOutChannel Pi calculation by channel
func CalculatePiByFanInOutChannel(seriesLength int, concurrency int) float64 {
	cPiTermGen := make(chan *PITerm)
	cPiTermResult := make(chan float64)

	generatePiTerm(seriesLength, cPiTermGen)

	go calculateFanOutIn(cPiTermGen, cPiTermResult, concurrency)

	var aggResult float64 = 3
	for v := range cPiTermResult {
		aggResult += v
	}
	return aggResult
}

func generatePiTerm(seriesLength int, cPiTermGen chan *PITerm) {

	go func() {
		for i := 0; i < seriesLength; i++ {

			n := algo.Nilakantha{}
			piterm := PITerm(&n)
			piterm.SetTerm(i)

			cPiTermGen <- &piterm
		}
		close(cPiTermGen)
	}()
}

func calculateFanOutIn(chanIn <-chan *PITerm, chanOut chan float64, concurrency int) {
	var wg sync.WaitGroup
	goroutines := concurrency
	wg.Add(goroutines)

	for i := 0; i < goroutines; i++ {
		go func() {
			for v := range chanIn {
				func(v2 *PITerm) {
					chanOut <- (*v2).CalculateTerm()
				}(v)
			}
			wg.Done()
		}()
	}
	//fmt.Println("GoRoutines\t", runtime.NumGoroutine())
	wg.Wait()
	close(chanOut)
}
