package pi_test

import (
	"fmt"
	"math"
	"pi"
	"runtime"
	"testing"
)

const numOfStep = 100000

func TestCalculatePiByArray(t *testing.T) {
	pi := pi.CalculatePiByArray(numOfStep)
	if math.Abs(math.Pi-pi) > 0.0000001 {
		t.Errorf("PI calculation inaccurate:%f", pi)
	} else {
		fmt.Printf("Calculate by Array OK: %f\n", pi)
	}
}

func TestCalculatePiByChannel(t *testing.T) {
	pi := pi.CalculatePiByFanInOutChannel(numOfStep, runtime.NumCPU())
	if math.Abs(math.Pi-pi) > 0.00001 {
		t.Errorf("PI calculation inaccurate:%f", pi)
	} else {
		fmt.Printf("Calculate by Channel OK: %f\n", pi)
	}
}

func BenchmarkCalculatePiByArray(t *testing.B) {
	for i := 0; i < t.N; i++ {
		_ = pi.CalculatePiByArray(numOfStep)
	}

}

func BenchmarkCalculatePiByChannel(t *testing.B) {
	for i := 0; i < t.N; i++ {
		_ = pi.CalculatePiByFanInOutChannel(numOfStep, runtime.NumCPU()*10)
	}

}
